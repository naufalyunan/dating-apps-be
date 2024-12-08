package handlers

import (
	"context"
	"errors"
	"logs-service/models"
	pb "logs-service/pb/generated"
	"time"

	"gorm.io/gorm"
)

type LogHandler struct {
	pb.UnimplementedLogServiceServer
	db *gorm.DB
}

func NewLogHandler(db *gorm.DB) *LogHandler {
	return &LogHandler{
		db: db,
	}
}

func (l *LogHandler) AddLog(ctx context.Context, req *pb.AddLogRequest) (*pb.AddLogResponse, error) {
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
		UserID:        uint(req.UserId),
		ActionType:    req.ActionType,
		ActionDetails: req.Details,
	}

	err := l.db.Create(&log).Error
	if err != nil {
		return nil, err
	}

	return &pb.AddLogResponse{
		Status: "Log Added Successfully",
		LogEntry: &pb.LogEntry{
			Id:         uint32(log.ID),
			UserId:     uint32(log.UserID),
			ActionType: log.ActionType,
			Details:    log.ActionDetails,
			Timestamp:  time.Now().Format(time.RFC3339),
		},
	}, nil
}

func (l *LogHandler) GetLogs(ctx context.Context, req *pb.GetLogsRequest) (*pb.GetLogsResponse, error) {
	// get all logs
	var logs []models.ActivityLog
	err := l.db.Find(&logs).Error
	if err != nil {
		return nil, err
	}

	// convert logs to pb
	var pbLogs []*pb.LogEntry
	for _, log := range logs {
		pbLogs = append(pbLogs, &pb.LogEntry{
			Id:         uint32(log.ID),
			UserId:     uint32(log.UserID),
			ActionType: log.ActionType,
			Details:    log.ActionDetails,
			Timestamp:  time.Now().Format(time.RFC3339),
		})
	}

	return &pb.GetLogsResponse{
		Logs: pbLogs,
	}, nil
}

func (l *LogHandler) StreamLogs(req *pb.StreamLogsRequest, stream pb.LogService_StreamLogsServer) error {
	// Keep track of the last processed log ID
	var lastLogID uint

	for {
		// Fetch new logs since the last processed ID
		var logs []models.ActivityLog
		err := l.db.Where("user_id = ? AND id > ?", req.UserId, lastLogID).Find(&logs).Error
		if err != nil {
			return err
		}

		// Stream new logs to the client
		for _, log := range logs {
			err := stream.Send(&pb.StreamLogsResponse{
				LogEntry: &pb.LogEntry{
					Id:         uint32(log.ID),
					UserId:     uint32(log.UserID),
					ActionType: log.ActionType,
					Details:    log.ActionDetails,
					Timestamp:  time.Now().Format(time.RFC3339),
				},
			})
			if err != nil {
				return err
			}

			// Update the last processed ID
			lastLogID = log.ID
		}

		// Wait for a short interval before polling again
		time.Sleep(1 * time.Second)
	}
}
