package handlers

import (
	"context"
	"date-service/entities"
	"date-service/models"
	pb "date-service/pb/generated"
	"date-service/services"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type SwipeHandler struct {
	pb.UnimplementedSwipeServiceServer
	db             *gorm.DB
	profileService services.ProfileService
	userService    services.UserService
	logService     services.LogService
}

func NewSwipeHandler(db *gorm.DB, profileService services.ProfileService, userService services.UserService, logService services.LogService) *SwipeHandler {
	return &SwipeHandler{
		db:             db,
		profileService: profileService,
		userService:    userService,
		logService:     logService,
	}
}

func (s *SwipeHandler) RecordSwipe(ctx context.Context, req *pb.RecordSwipeRequest) (*pb.RecordSwipeResponse, error) {
	// validate token and get user
	_, err := s.userService.ValidateAndGetUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid token '%s'", err.Error())
	}

	//validate requests
	if req.SwiperUserId == 0 {
		return nil, errors.New("swiper_user_id is required")
	}

	if req.SwipedProfileUserId == 0 {
		return nil, errors.New("swiped_user_id is required")
	}

	if req.Action == "" {
		return nil, errors.New("action is required")
	}

	//check if swiper user id and swiped user id is the same
	if req.SwiperUserId == req.SwipedProfileUserId {
		return nil, errors.New("You cannot swipe yourself")
	}

	//check if user already swiped
	var swipe models.Swipe
	err = s.db.Where("swiper_user_id = ? AND swiped_profile_user_id = ?", req.SwiperUserId, req.SwipedProfileUserId).First(&swipe).Error
	if err == nil {
		return nil, errors.New("You have already swiped this profile")
	}

	//check if it is already swiping 10 times in history for 24 hours, if it is then return Error
	timeLimit := time.Now().Add(-24 * time.Hour)

	var swipes []models.Swipe
	err = s.db.Where("swiper_user_id = ? AND created_at > ?", req.SwiperUserId, timeLimit).Find(&swipes).Error
	if err != nil {
		return nil, err
	}

	if len(swipes) >= 10 {
		return nil, errors.New("You have reached the limit of swiping 10 times in 24 hours")
	}

	// create new swipe
	swipe = models.Swipe{
		SwiperUserID:        uint(req.SwiperUserId),
		SwipedProfileUserID: uint(req.SwipedProfileUserId),
		Action:              req.Action,
	}

	err = s.db.Create(&swipe).Error
	if err != nil {
		return nil, err
	}

	_, err = s.logService.AddLog(entities.ActivityLog{
		UserID:        uint(req.SwiperUserId),
		ActionType:    "Swipe",
		ActionDetails: fmt.Sprintf("User %d swipe user %d with action %s", swipe.SwiperUserID, swipe.SwipedProfileUserID, swipe.Action),
	})
	if err != nil {
		return nil, err
	}

	// Check for a match if the swipe action is "like"
	if req.Action == "like" {
		var reverseSwipe models.Swipe
		err = s.db.Where("swiper_user_id = ? AND swiped_profile_user_id = ? AND action = ?", req.SwipedProfileUserId, req.SwiperUserId, "like").First(&reverseSwipe).Error
		if err == nil {
			// Mutual "like" found, create a match
			match := models.Match{
				User1ID: uint(req.SwiperUserId),
				User2ID: uint(req.SwipedProfileUserId),
			}
			err = s.db.Create(&match).Error
			if err != nil {
				return nil, err
			}

			_, err = s.logService.AddLog(entities.ActivityLog{
				UserID:        swipe.SwiperUserID,
				ActionType:    "Found Match",
				ActionDetails: fmt.Sprintf("Found Match beteween user %d and user %d", req.SwiperUserId, req.SwipedProfileUserId),
			})
			if err != nil {
				return nil, err
			}

			// Optionally, notify the users of the match (e.g., via a message queue or push notification)
			fmt.Printf("Match found between user %d and user %d\n", req.SwiperUserId, req.SwipedProfileUserId)
		}
	}

	return &pb.RecordSwipeResponse{
		Status: fmt.Sprintf("Successfully %s profile with user id %d", req.Action, req.SwipedProfileUserId),
		Swipe:  &pb.SwipeAction{Id: uint32(swipe.ID), SwiperUserId: uint32(swipe.SwiperUserID), SwipedProfileUserId: uint32(swipe.SwipedProfileUserID), Action: swipe.Action},
	}, nil
}

func contains(slice []uint, value uint) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func (s *SwipeHandler) GetSuggestions(ctx context.Context, req *pb.GetSuggestionsRequest) (*pb.GetSuggestionsResponse, error) {
	// validate token and get user
	_, err := s.userService.ValidateAndGetUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid token '%s'", err.Error())
	}

	//validate requests
	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	if req.Limit == 0 {
		req.Limit = 20
	}

	//get suggested profiles
	profiles, err := s.profileService.GetProfiles(int(req.UserId), int(req.Limit))
	if err != nil {
		return nil, err
	}

	//make sure the profiles are not the ones that the user has swiped
	//get all the swipes of the user
	var swipes []models.Swipe
	err = s.db.Where("swiper_user_id = ?", req.UserId).Find(&swipes).Error
	if err != nil {
		return nil, err
	}

	//get the ids of the profiles that the user has swiped
	swipedProfileIds := make([]uint, 0)
	for _, swipe := range swipes {
		swipedProfileIds = append(swipedProfileIds, swipe.SwipedProfileUserID)
	}

	//filter out the profiles that the user has swiped
	filteredProfiles := make([]*entities.Profile, 0)
	for _, profile := range profiles {
		if !contains(swipedProfileIds, uint(profile.ID)) {
			filteredProfiles = append(filteredProfiles, profile)
		}
	}

	//insert limit after filter
	if len(filteredProfiles) > int(req.Limit) {
		filteredProfiles = filteredProfiles[:req.Limit]
	}

	//convert the profiles to proto profiles
	converted := make([]*pb.ProfileShow, 0)
	for _, profile := range filteredProfiles {
		converted = append(converted, &pb.ProfileShow{
			Id:     uint32(profile.ID),
			UserId: uint32(profile.UserID),
			Age:    int32(profile.Age),
			Bio:    profile.Bio,
			Photos: profile.Photos,
		})
	}

	return &pb.GetSuggestionsResponse{
		Profiles: converted,
	}, nil
}

func (s *SwipeHandler) GetSwipeHistory(ctx context.Context, req *pb.GetSwipeHistoryRequest) (*pb.GetSwipeHistoryResponse, error) {
	//validate requests
	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	//get all the swipes of the user
	var swipes []models.Swipe
	err := s.db.Offset(int(req.Offset)).Limit(int(req.Limit)).Where("swiper_user_id = ?", req.UserId).Find(&swipes).Error
	if err != nil {
		return nil, err
	}

	//convert the swipes to proto swipes
	converted := make([]*pb.SwipeAction, 0)
	for _, swipe := range swipes {
		converted = append(converted, &pb.SwipeAction{
			Id:                  uint32(swipe.ID),
			SwiperUserId:        uint32(swipe.SwiperUserID),
			SwipedProfileUserId: uint32(swipe.SwipedProfileUserID),
			Action:              swipe.Action,
		})
	}

	return &pb.GetSwipeHistoryResponse{
		Swipes: converted,
	}, nil
}
