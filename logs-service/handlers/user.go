package handlers

import (
	"context"
	"errors"
	"logs/models"
	pb "logs/pb/generated"

	"gorm.io/gorm"
)

type LogHandler struct {
	pb.UnimplementedLogServiceServer
	db         *gorm.DB
}

func NewLogHandler(db *gorm.DB) *LogHandler {
	return &LogHandler{
		db:         db,
	}
}

func (l *LogHandler) CreateLog(ctx context.Context, req *pb.AddLogRequest) (*pb.AddLogResponse, error) {
	//validate requests
	if req.UserId == 0 {
		return nil, errors.New("user id is required")
	}

	if req.ActionType == "" {
		return nil, errors.New("action type is required")
	}

	if req.Details == "" {
		return nil, errors.New("details is required")
	}

	// create new log
	log := models.ActivityLog{
		UserID: 	 uint(req.UserId),
		ActionType:  req.ActionType,
		ActionDetails: req.Details,
	}

	err := l.db.Create(&log).Error
	if err != nil {
		return nil, err
	}

	return &pb.AddLogResponse{
		Status: "Log Added Successfully",
		LogEntry: &pb.LogEntry{
			Id: string(log.ID),
			UserId: uint32(log.UserID),
			ActionType: log.ActionType,
			Details: log.ActionDetails,
		},
	}, nil
}