package services

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"date-service/entities"
	pb "date-service/pb/generated"
	"errors"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type UserService interface {
	IsValidToken(token string) (*entities.User, error)
	ValidateAndGetUser(c context.Context) (*entities.User, error)
}

func NewUserClient() pb.UserServiceClient {
	addr := os.Getenv("USER_SERVICE_ADDR")

	opts := []grpc.DialOption{}
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred))
	conn, err := grpc.NewClient(addr, opts...)
	// conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserServiceClient(conn)

	return client
}

func NewUserService() UserService {
	return &userService{
		userClient: NewUserClient(),
	}
}

type userService struct {
	userClient pb.UserServiceClient
}

func (u *userService) IsValidToken(token string) (*entities.User, error) {
	//validate requests
	if token == "" {
		return nil, errors.New("token is required")
	}

	res, err := u.userClient.IsValidToken(context.TODO(), &pb.IsValidTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:         uint(res.User.Id),
		Email:      res.User.Email,
		Username:   res.User.Username,
		IsPremium:  res.User.IsPremium,
		IsVerified: res.User.IsVerified,
	}, nil
}

func extractAuthToken(ctx context.Context) (string, error) {
	// Retrieve metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Print("no metadata found in context")
		return "", fmt.Errorf("no metadata found in context")
	}
	// Check if 'auth_token' exists in the metadata
	tokens := md["auth_token"]
	if len(tokens) == 0 {
		log.Print("auth_token not found in metadata")
		return "", fmt.Errorf("auth_token not found in metadata")
	}

	log.Printf("auth_token found '%s'", tokens[0])
	// Return the first token (in case there are multiple)
	return tokens[0], nil
}

func (u *userService) ValidateAndGetUser(c context.Context) (*entities.User, error) {
	// extract token
	token, err := extractAuthToken(c)
	if err != nil {
		return nil, err
	}

	// call user Service
	user, err := u.IsValidToken(token)
	if err != nil {
		return nil, err
	}

	return user, nil
}
