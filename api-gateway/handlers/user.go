package handlers

import (
	pb "api-gateway/pb/generated"
	"api-gateway/utils"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) HandleRegister(c echo.Context) error {
	var req pb.CreateUserRequest
	err := c.Bind(&req)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "invalid request body", err.Error())
	}

	res, err := h.UserClient.Register(
		context.TODO(),
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)

}

func (h *Handlers) HandleLogin(c echo.Context) error {
	var req pb.LoginUserRequest
	err := c.Bind(&req)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "invalid request body", err.Error())
	}

	res, err := h.UserClient.Login(
		context.TODO(),
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)

}
