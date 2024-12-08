package services

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"profiles-service/entities"
	pb "profiles-service/pb/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type LogService interface {
	AddLog(req entities.ActivityLog) (*entities.ActivityLog, error)
}

func NewLogClient() pb.LogServiceClient {
	addr := os.Getenv("LOG_SERVICE_ADDR")

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

	client := pb.NewLogServiceClient(conn)

	return client
}

func NewLogService() LogService {
	return &logService{
		logClient: NewLogClient(),
	}
}

type logService struct {
	logClient pb.LogServiceClient
}

func (l *logService) AddLog(req entities.ActivityLog) (*entities.ActivityLog, error) {
	res, err := l.logClient.AddLog(context.TODO(), &pb.AddLogRequest{
		UserId:     uint32(req.UserID),
		ActionType: req.ActionType,
		Details:    req.ActionDetails,
	})
	if err != nil {
		return nil, err
	}

	return &entities.ActivityLog{
		ID:            uint(res.LogEntry.Id),
		UserID:        uint(res.LogEntry.UserId),
		ActionType:    res.LogEntry.ActionType,
		ActionDetails: res.LogEntry.Details,
	}, nil
}
