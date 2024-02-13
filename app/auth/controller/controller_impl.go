package controller

import (
	"net/http"
	"smartpatrol/app/auth"
	"smartpatrol/app/auth/repository"
	"smartpatrol/app/auth/service"
	// "smartpatrol/models"
	"smartpatrol/pkg/helpers"
	"smartpatrol/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type controllerImpl struct {
	Service service.Service
}

func NewController(Service service.Service) Controller {
	return &controllerImpl {
		Service: Service,
	}
}

func (co *controllerImpl) Login(c echo.Context) error {
	var request auth.LoginRequest

	if err := c.Bind(&request); err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrBadRequest, err))
	}

	result, err := co.Service.Login(c, request)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Berhasil login", result)
}

func CreateAdmin(c echo.Context) error {
	var newUser auth.LoginRequest
	if err := c.Bind(&newUser);

	err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message" : "bad request",
			"error" : err.Error(),
			"status_code" : http.StatusBadRequest,
		})
	}

	if err := helpers.HashPassword(&newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	_, err := repository.CreateUser(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":     "Failed to create user",
			"error":       err.Error(),
			"status_code": http.StatusInternalServerError,
		})
	}

	return response.SuccessResponse(c, http.StatusCreated, "Success", "")
	}