package controller

import (
	"net/http"
	"smartpatrol/app/role/service"
	"smartpatrol/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type controllerImpl struct {
	Service service.Service
}

func NewController(Service service.Service) Controller {
	return &controllerImpl{
		Service: Service,
	}
}

func (co *controllerImpl) GetAllUser(c echo.Context) error {
	result, err := co.Service.GetAllUser(c)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Department", result)
}