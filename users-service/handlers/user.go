package handlers

import (
	"context"
	"errors"
	"fmt"
	"os"
	"users-service/entities"
	"users-service/models"
	pb "users-service/pb/generated"
	"users-service/services"
	"users-service/utils"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	db         *gorm.DB
	logService services.LogService
}

func NewUserHandler(db *gorm.DB, logService services.LogService) *UserHandler {
	return &UserHandler{
		db:         db,
		logService: logService,
	}
}

func (u *UserHandler) Register(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	//validate requests
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	if req.Username == "" {
		return nil, errors.New("username is required")
	}

	//hash password
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// create new user
	user := models.User{
		Email:        req.Email,
		PasswordHash: hash,
		Username:     req.Username,
		IsPremium:    false,
		IsVerified:   false,
	}

	err = u.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	_, err = u.logService.AddLog(entities.ActivityLog{
		UserID:        user.ID,
		ActionType:    "Register",
		ActionDetails: fmt.Sprintf("User %s registered", user.Username),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Status: "User Created Successfully",
		User: &pb.User{
			Id:         uint32(user.ID),
			Email:      user.Email,
			Username:   user.Username,
			IsPremium:  user.IsPremium,
			IsVerified: user.IsVerified,
		},
	}, nil
}

func (u *UserHandler) Login(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	//validate requests
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	if req.Password == "" {
		return nil, errors.New("password is required")
	}

	// check if user exists
	var user models.User
	err := u.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}

	key := os.Getenv("JWT_SECRET")

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      user.Email,
		"username":   user.Username,
		"id":         user.ID,
		"is_premium": user.IsPremium,
	})

	s, err := t.SignedString([]byte(key))
	if err != nil {
		return nil, errors.New("failed to sign token")
	}

	_, err = u.logService.AddLog(entities.ActivityLog{
		UserID:        user.ID,
		ActionType:    "Login",
		ActionDetails: fmt.Sprintf("User %s Login succesfully", user.Username),
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginUserResponse{
		Status: "Login Successful",
		Token:  s,
		User: &pb.User{
			Id:         uint32(user.ID),
			Email:      user.Email,
			Username:   user.Username,
			IsPremium:  user.IsPremium,
			IsVerified: user.IsVerified,
		},
	}, nil
}

func (u *UserHandler) IsValidToken(ctx context.Context, req *pb.IsValidTokenRequest) (*pb.IsValidTokenResponse, error) {
	//validate requests
	if req.Token == "" {
		return nil, errors.New("token is required")
	}
	fmt.Printf("Token diterima di user service handler :%s\n", req.Token)
	key := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token '%v',", token)
	}

	var user models.User
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	err = u.db.Where("email = ?", claims["email"]).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	res := &pb.IsValidTokenResponse{
		Valid: true,
		User: &pb.User{
			Id:         uint32(user.ID),
			Email:      user.Email,
			Username:   user.Username,
			IsPremium:  user.IsPremium,
			IsVerified: user.IsVerified,
		},
	}
	return res, nil
}

func (u *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	//validate requests
	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	var user models.User
	err := u.db.Where("id = ?", req.Id).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:         uint32(user.ID),
			Email:      user.Email,
			Username:   user.Username,
			IsPremium:  user.IsPremium,
			IsVerified: user.IsVerified,
		},
	}, nil
}

func (u *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	//validate requests
	if req.Id == 0 {
		return nil, errors.New("id is required")
	}

	var user models.User
	err := u.db.Where("id = ?", req.Id).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Username != "" {
		user.Username = req.Username
	}

	if req.IsPremium {
		user.IsPremium = req.IsPremium
	}

	if req.IsVerified {
		user.IsVerified = req.IsVerified
	}

	//check if email is already used by another user and id is not the same
	var user2 models.User
	err = u.db.Where("email = ? AND id != ?", req.Email, req.Id).First(&user2).Error
	if err == nil {
		return nil, errors.New("email already used by another user")
	}

	err = u.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		Status: "User Updated Successfully",
		User: &pb.User{
			Id:         uint32(user.ID),
			Email:      user.Email,
			Username:   user.Username,
			IsPremium:  user.IsPremium,
			IsVerified: user.IsVerified,
		},
	}, nil
}
