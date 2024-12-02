package handlers

import (
	"context"
	"date-service/models"
	pb "date-service/pb/generated"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SwipeHandler struct {
	pb.UnimplementedSwipeServiceServer
	db *gorm.DB
}

func NewSwipeHandler(db *gorm.DB) *SwipeHandler {
	return &SwipeHandler{
		db: db,
	}
}

func (s *SwipeHandler) RecordSwipe(ctx context.Context, req *pb.RecordSwipeRequest) (*pb.RecordSwipeResponse, error) {
	//validate requests
	if req.SwiperId == 0 {
		return nil, errors.New("swiper_id is required")
	}

	if req.SwipedProfileId == 0 {
		return nil, errors.New("swiped_id is required")
	}

	if req.Action == "" {
		return nil, errors.New("action is required")
	}

	// create new swipe
	swipe := models.Swipe{
		SwiperID:        uint(req.SwiperId),
		SwipedProfileID: uint(req.SwipedProfileId),
		Action:          req.Action,
	}

	err := s.db.Create(&swipe).Error
	if err != nil {
		return nil, err
	}

	return &pb.RecordSwipeResponse{
		Status: fmt.Sprintf("Successfully %s profile with id %d", req.Action, req.SwipedProfileId),
		Swipe:  &pb.SwipeAction{Id: uint32(swipe.ID), SwiperId: uint32(swipe.SwiperID), SwipedProfileId: uint32(swipe.SwipedProfileID), Action: swipe.Action},
	}, nil
}
