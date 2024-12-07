package handlers

import (
	pb "api-gateway/pb/generated"
	"api-gateway/utils"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) HandleGetSuggestions(c echo.Context) error {
	token := utils.ExtractAuthToken(c)
	user, err := h.UserClient.IsValidToken(context.TODO(), &pb.IsValidTokenRequest{Token: token})
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}
	req := pb.GetSuggestionsRequest{
		UserId: user.User.Id,
	}

	limit := c.QueryParam("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return utils.NewAppError(http.StatusBadRequest, "invalid limit", err.Error())
		}
		req.Limit = uint32(l)
	}

	ctx := utils.CreateContext(c)

	res, err := h.DateClient.GetSuggestions(
		ctx,
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Handlers) HandleRecordSwipe(c echo.Context) error {
	token := utils.ExtractAuthToken(c)
	user, err := h.UserClient.IsValidToken(context.TODO(), &pb.IsValidTokenRequest{Token: token})
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}
	req := pb.RecordSwipeRequest{
		SwiperUserId: user.User.Id,
	}
	err = c.Bind(&req)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "invalid request body", err.Error())
	}

	ctx := utils.CreateContext(c)
	res, err := h.DateClient.RecordSwipe(
		ctx,
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Handlers) HandleSwipeHistory(c echo.Context) error {
	token := utils.ExtractAuthToken(c)
	user, err := h.UserClient.IsValidToken(context.TODO(), &pb.IsValidTokenRequest{Token: token})
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}
	req := pb.GetSwipeHistoryRequest{
		UserId: user.User.Id,
	}

	limit := c.QueryParam("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return utils.NewAppError(http.StatusBadRequest, "invalid limit", err.Error())
		}
		req.Limit = uint32(l)
	} else {
		req.Limit = 0
	}

	offset := c.QueryParam("offset")
	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return utils.NewAppError(http.StatusBadRequest, "invalid offset", err.Error())
		}
		req.Offset = uint32(o)
	} else {
		req.Offset = 0
	}

	err = c.Bind(&req)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "invalid request body", err.Error())
	}

	ctx := utils.CreateContext(c)
	res, err := h.DateClient.GetSwipeHistory(
		ctx,
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}
