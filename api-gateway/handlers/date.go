package handlers

import (
	pb "api-gateway/pb/generated"
	"api-gateway/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) HandleGetSuggestions(c echo.Context) error {
	var req pb.GetSuggestionsRequest
	err := c.Bind(&req)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "invalid request body", err.Error())
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
	var req pb.RecordSwipeRequest
	err := c.Bind(&req)
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
