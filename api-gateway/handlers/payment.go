package handlers

import (
	pb "api-gateway/pb/generated"
	"api-gateway/utils"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) HandlePaymentCallback(c echo.Context) error {
	// pb definitions have json annotations, can use it directly
	var req pb.CompletePaymentReq
	err := c.Bind(&req)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "invalid request body", err.Error())
	}

	// verify webhook token
	verifToken := c.Request().Header.Get("x-callback-token")
	if verifToken == "" {
		return utils.NewAppError(http.StatusUnauthorized, "invalid webhook token", "")
	}
	// set token from header
	req.CallbackToken = verifToken

	res, err := h.PaymentClient.CompletePayment(
		context.TODO(),
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handlers) HandleCreateUserSubscription(c echo.Context) error {
	// pb definitions have json annotations, can use it directly
	var req pb.CreateUserSubcriptionReq
	err := c.Bind(&req)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "invalid request body", err.Error())
	}

	ctx := utils.CreateContext(c)
	// forward
	req.UserId = 1
	res, err := h.PaymentClient.CreateUserSubcription(
		ctx,
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}
