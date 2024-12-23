package handlers

import (
	"context"
	"errors"
	"fmt"
	"profiles-service/entities"
	"profiles-service/models"
	pb "profiles-service/pb/generated"
	"profiles-service/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	db          *gorm.DB
	userService services.UserService
	logService  services.LogService
}

func NewProfileHandler(db *gorm.DB, userService services.UserService, logService services.LogService) *ProfileHandler {
	return &ProfileHandler{
		db:          db,
		userService: userService,
		logService:  logService,
	}
}

func (p *ProfileHandler) GetProfilesSuggestion(ctx context.Context, req *pb.GetProfilesSuggestionRequest) (*pb.GetProfilesSuggestionResponse, error) {
	//validate the field
	if req.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	// get all profiles except the user

	profiles := []models.Profile{}
	err := p.db.Not("user_id = ?", req.UserId).Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	profilesResponse := []*pb.Profile{}
	for _, profile := range profiles {
		profilesResponse = append(profilesResponse, &pb.Profile{
			Id:     uint32(profile.ID),
			UserId: uint32(profile.UserID),
			Age:    int32(profile.Age),
			Bio:    profile.Bio,
			Photos: profile.Photos,
		})
	}

	return &pb.GetProfilesSuggestionResponse{
		Profiles: profilesResponse,
	}, nil
}

func (p *ProfileHandler) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	// validate token and get user
	user, err := p.userService.ValidateAndGetUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid token '%s'", err.Error())
	}

	//validate the field
	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	profile := models.Profile{}
	err = p.db.Where("id = ?", req.Id).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}

	_, err = p.logService.AddLog(entities.ActivityLog{
		UserID:        user.ID,
		ActionType:    "Get Profile",
		ActionDetails: fmt.Sprintf("User %s Getting Profile with Profile ID %d", user.Username, profile.ID),
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetProfileResponse{
		Profile: &pb.Profile{
			Id:     uint32(profile.ID),
			UserId: uint32(profile.UserID),
			Age:    int32(profile.Age),
			Bio:    profile.Bio,
			Photos: profile.Photos,
		},
	}, nil
}

func (p *ProfileHandler) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	// validate token and get user
	user, err := p.userService.ValidateAndGetUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid token '%s'", err.Error())
	}

	//validate the field
	if req.Age == 0 {
		return nil, errors.New("age is required")
	}

	if req.Bio == "" {
		return nil, errors.New("bio is required")
	}

	// create new profile
	profile := models.Profile{
		UserID: uint(user.ID),
		Age:    int(req.Age),
		Bio:    req.Bio,
		Photos: req.Photos,
	}

	err = p.db.Create(&profile).Error
	if err != nil {
		return nil, err
	}

	_, err = p.logService.AddLog(entities.ActivityLog{
		UserID:        user.ID,
		ActionType:    "Create Profile",
		ActionDetails: fmt.Sprintf("User %s Creating Profile with ID %d", user.Username, profile.ID),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateProfileResponse{
		Status: "Successfully created profile",
		Profile: &pb.Profile{
			Id:     uint32(profile.ID),
			UserId: uint32(profile.UserID),
			Age:    int32(profile.Age),
			Bio:    profile.Bio,
			Photos: profile.Photos,
		},
	}, nil
}

func (p *ProfileHandler) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	// validate token and get user
	user, err := p.userService.ValidateAndGetUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid token '%s'", err.Error())
	}

	//validate the field
	if req.Age == 0 {
		return nil, errors.New("age is required")
	}

	if req.Bio == "" {
		return nil, errors.New("bio is required")
	}

	//check if the profile exists
	profile := models.Profile{}
	err = p.db.Where("id = ?", req.Id).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}

	// create updated new profile

	toUpdate := models.Profile{
		UserID: uint(user.ID),
		Age:    int(req.Age),
		Bio:    req.Bio,
		Photos: req.Photos,
	}

	err = p.db.Model(&models.Profile{}).Where("id = ?", req.Id).Updates(&toUpdate).Error
	if err != nil {
		return nil, err
	}

	_, err = p.logService.AddLog(entities.ActivityLog{
		UserID:        user.ID,
		ActionType:    "Update Profile",
		ActionDetails: fmt.Sprintf("User %s Updating Profile with ID %d", user.Username, profile.ID),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProfileResponse{
		Status: "Successfully updated profile",
		Profile: &pb.Profile{
			Id:     uint32(profile.ID),
			UserId: uint32(profile.UserID),
			Age:    int32(profile.Age),
			Bio:    profile.Bio,
			Photos: profile.Photos,
		},
	}, nil
}

func (p *ProfileHandler) DeleteProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {
	// validate token and get user
	user, err := p.userService.ValidateAndGetUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid token '%s'", err.Error())
	}

	//validate the field
	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	err = p.db.Where("id = ?", req.Id).Delete(&models.Profile{}).Error
	if err != nil {
		return nil, err
	}

	_, err = p.logService.AddLog(entities.ActivityLog{
		UserID:        user.ID,
		ActionType:    "Delete Profile",
		ActionDetails: fmt.Sprintf("User %s Deleting Profile with ID %d", user.Username, req.Id),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProfileResponse{
		Status: "Successfully deleted profile",
	}, nil
}
