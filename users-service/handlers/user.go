package handlers

import (
	"context"
	"errors"
	"os"
	"users-service/models"
	pb "users-service/pb/generated"
	"users-service/utils"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
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
