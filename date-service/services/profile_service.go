package services

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"date-service/entities"
	pb "date-service/pb/generated"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ProfileService interface {
	GetProfiles(id int, limit int) ([]*entities.Profile, error)
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
	// conn, err := grpc.Dial(addr, grpc.WithInsecure())
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

func (p *profileService) GetProfiles(id int, limit int) ([]*entities.Profile, error) {
	res, err := p.profileClient.GetProfilesSuggestion(context.TODO(), &pb.GetProfilesSuggestionRequest{
		UserId: uint32(id),
	})
	if err != nil {
		return nil, err
	}

	profiles := make([]*entities.Profile, 0)
	for _, profile := range res.Profiles {
		profiles = append(profiles, &entities.Profile{
			ID:     int(profile.Id),
			UserID: uint(profile.UserId),
			Age:    int(profile.Age),
			Bio:    profile.Bio,
			Photos: profile.Photos,
		})
	}

	return profiles, nil
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
