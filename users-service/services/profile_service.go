package services

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"users-service/entities"
	pb "users-service/pb/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ProfileService interface {
	CreateProfile(req entities.Profile) (*entities.Profile, error)
	UpdateProfile(req entities.Profile) (*entities.Profile, error)
	GetProfile(id int) (*entities.Profile, error)
}

func NewProfileClient() pb.ProfileServiceClient {
	addr := os.Getenv("PROFILE_SERVICE_ADDR")

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
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewProfileServiceClient(conn)

	return client
}

func NewProfileService() ProfileService {
	return &profileService{
		profileClient: NewProfileClient(),
	}
}

type profileService struct {
	profileClient pb.ProfileServiceClient
}

func (p *profileService) CreateProfile(req entities.Profile) (*entities.Profile, error) {
	res, err := p.profileClient.CreateProfile(context.TODO(), &pb.CreateProfileRequest{
		UserId: uint32(req.UserID),
		Age:    int32(req.Age),
		Bio:    req.Bio,
		Photos: req.Photos,
	})
	if err != nil {
		return nil, err
	}

	return &entities.Profile{
		ID:     int(res.Profile.Id),
		UserID: uint(res.Profile.UserId),
		Age:    int(res.Profile.Age),
		Bio:    res.Profile.Bio,
		Photos: res.Profile.Photos,
	}, nil
}

func (p *profileService) UpdateProfile(req entities.Profile) (*entities.Profile, error) {
	res, err := p.profileClient.UpdateProfile(context.TODO(), &pb.UpdateProfileRequest{
		UserId: uint32(req.UserID),
		Age:    int32(req.Age),
		Bio:    req.Bio,
		Photos: req.Photos,
	})
	if err != nil {
		return nil, err
	}

	return &entities.Profile{
		ID:     int(res.Profile.Id),
		UserID: uint(res.Profile.UserId),
		Age:    int(res.Profile.Age),
		Bio:    res.Profile.Bio,
		Photos: res.Profile.Photos,
	}, nil
}

func (p *profileService) GetProfile(id int) (*entities.Profile, error) {
	res, err := p.profileClient.GetProfile(context.TODO(), &pb.GetProfileRequest{
		UserId: uint32(id),
	})
	if err != nil {
		return nil, err
	}

	return &entities.Profile{
		ID:     int(res.Profile.Id),
		UserID: uint(res.Profile.UserId),
		Age:    int(res.Profile.Age),
		Bio:    res.Profile.Bio,
		Photos: res.Profile.Photos,
	}, nil
}
