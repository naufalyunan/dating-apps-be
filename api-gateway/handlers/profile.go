package handlers

import (
	pb "api-gateway/pb/generated"
	"api-gateway/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) HandleCreateProfile(c echo.Context) error {
	var req pb.CreateProfileRequest
	err := c.Bind(&req)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "invalid request body", err.Error())
	}

	ctx := utils.CreateContext(c)
	res, err := h.ProfileClient.CreateProfile(
		ctx,
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Handlers) HandleGetProfile(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}
	req := pb.GetProfileRequest{
		UserId: uint32(id),
	}

	ctx := utils.CreateContext(c)
	res, err := h.ProfileClient.GetProfile(
		ctx,
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Handlers) HandleUpdateProfile(c echo.Context) error {
	// get id param
	idParam := c.Param("id")
	bfId, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid profile id")
	}

	var req pb.UpdateProfileRequest
	err = c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	req.UserId = uint32(bfId)

	ctx := utils.CreateContext(c)
	res, err := h.ProfileClient.UpdateProfile(
		ctx,
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Handlers) HandleDeleteProfile(c echo.Context) error {
	// get id param
	idParam := c.Param("id")
	bfId, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid profile id")
	}

	var req pb.DeleteProfileRequest
	req.UserId = uint32(bfId)

	ctx := utils.CreateContext(c)
	res, err := h.ProfileClient.DeleteProfile(
		ctx,
		&req,
	)
	if err != nil {
		return utils.NewAppError(http.StatusBadRequest, "service error", err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}
