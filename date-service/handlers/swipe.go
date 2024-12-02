package handlers

import (
	"context"
	"date-service/entities"
	"date-service/models"
	pb "date-service/pb/generated"
	"date-service/services"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SwipeHandler struct {
	pb.UnimplementedSwipeServiceServer
	db             *gorm.DB
	profileService services.ProfileService
}

func NewSwipeHandler(db *gorm.DB, profileService services.ProfileService) *SwipeHandler {
	return &SwipeHandler{
		db:             db,
		profileService: profileService,
	}
}

func (s *SwipeHandler) RecordSwipe(ctx context.Context, req *pb.RecordSwipeRequest) (*pb.RecordSwipeResponse, error) {
	//validate requests
	if req.SwiperId == 0 {
		return nil, errors.New("swiper_id is required")
	}

	if req.SwipedProfileId == 0 {
		return nil, errors.New("swiped_id is required")
	}

	if req.Action == "" {
		return nil, errors.New("action is required")
	}

	// create new swipe
	swipe := models.Swipe{
		SwiperID:        uint(req.SwiperId),
		SwipedProfileID: uint(req.SwipedProfileId),
		Action:          req.Action,
	}

	err := s.db.Create(&swipe).Error
	if err != nil {
		return nil, err
	}

	return &pb.RecordSwipeResponse{
		Status: fmt.Sprintf("Successfully %s profile with id %d", req.Action, req.SwipedProfileId),
		Swipe:  &pb.SwipeAction{Id: uint32(swipe.ID), SwiperId: uint32(swipe.SwiperID), SwipedProfileId: uint32(swipe.SwipedProfileID), Action: swipe.Action},
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
	//validate requests
	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	if req.Limit == 0 {
		req.Limit = 20
	}

	//get suggested profiles
	profiles, err := s.profileService.GetProfiles(int(req.UserId))
	if err != nil {
		return nil, err
	}

	//make sure the profiles are not the ones that the user has swiped
	//get all the swipes of the user
	var swipes []models.Swipe
	err = s.db.Where("swiper_id = ?", req.UserId).Find(&swipes).Error
	if err != nil {
		return nil, err
	}

	//get the ids of the profiles that the user has swiped
	swipedProfileIds := make([]uint, 0)
	for _, swipe := range swipes {
		swipedProfileIds = append(swipedProfileIds, swipe.SwipedProfileID)
	}

	//filter out the profiles that the user has swiped
	filteredProfiles := make([]*entities.Profile, 0)
	for _, profile := range profiles {
		if !contains(swipedProfileIds, uint(profile.ID)) {
			filteredProfiles = append(filteredProfiles, profile)
		}
	}

	//limit the number of profiles to the limit
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
	err := s.db.Where("swiper_id = ?", req.UserId).Find(&swipes).Error
	if err != nil {
		return nil, err
	}

	//convert the swipes to proto swipes
	converted := make([]*pb.SwipeAction, 0)
	for _, swipe := range swipes {
		converted = append(converted, &pb.SwipeAction{
			Id:              uint32(swipe.ID),
			SwiperId:        uint32(swipe.SwiperID),
			SwipedProfileId: uint32(swipe.SwipedProfileID),
			Action:          swipe.Action,
		})
	}

	return &pb.GetSwipeHistoryResponse{
		Swipes: converted,
	}, nil
}
