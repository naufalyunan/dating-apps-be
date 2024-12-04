package services

import (
	"context"
	"fmt"
	"log"
	"os"
	pb "payment-service/pb/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type User struct {
	ID         int
	Username   string
	Email      string
	IsPremium  bool
	IsVerified bool
}

func NewUserClient() pb.UserServiceClient {
	addr := os.Getenv("USER_SERVICE_ADDR")
	log.Printf("user service url: %s", addr)
	// Set up a connection to the server.
	// opts := []grpc.DialOption{}
	// systemRoots, err := x509.SystemCertPool()
	// if err != nil {
	// 	log.Fatalf("filed to get certs: %v", err)
	// }
	// cred := credentials.NewTLS(&tls.Config{
	// 	RootCAs: systemRoots,
	// })
	// opts = append(opts, grpc.WithTransportCredentials(cred))
	// conn, err := grpc.NewClient(addr, opts...)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewUserServiceClient(conn)

	return client
}

type UserService interface {
	IsValidToken(token string) (*User, error)
	ValidateAndGetUser(c context.Context) (*User, error)
	GetUserByID(id int) (*User, error)
	UpdateUser(user *User) (*User, error)
}

func NewUserService() UserService {
	return &userService{
		userClient: NewUserClient(),
	}
}

type userService struct {
	userClient pb.UserServiceClient
}

func (u *userService) IsValidToken(token string) (*User, error) {
	res, err := u.userClient.IsValidToken(context.TODO(), &pb.IsValidTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:         int(res.User.Id),
		Username:   res.User.Username,
		Email:      res.User.Email,
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

func (u *userService) ValidateAndGetUser(c context.Context) (*User, error) {
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

	return &User{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		IsPremium:  user.IsPremium,
		IsVerified: user.IsVerified,
	}, nil
}

func (u *userService) GetUserByID(id int) (*User, error) {

	res, err := u.userClient.GetUser(context.TODO(), &pb.GetUserRequest{
		Id: uint32(id),
	})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:         id,
		Username:   res.User.Username,
		Email:      res.User.Email,
		IsPremium:  res.User.IsPremium,
		IsVerified: res.User.IsVerified,
	}, nil
}

func (u *userService) UpdateUser(user *User) (*User, error) {
	res, err := u.userClient.UpdateUser(context.TODO(), &pb.UpdateUserRequest{
		Username:   user.Username,
		Email:      user.Email,
		IsPremium:  user.IsPremium,
		IsVerified: user.IsVerified,
		Id:         uint32(user.ID),
	})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:         int(res.User.Id),
		Username:   res.User.Username,
		Email:      res.User.Email,
		IsPremium:  res.User.IsPremium,
		IsVerified: res.User.IsVerified,
	}, nil
}
