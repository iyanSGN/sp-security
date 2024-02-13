package controller

import (
	"fmt"
	"net/http"
	"smartpatrol/app/users"
	"smartpatrol/app/users/repository"
	"smartpatrol/app/users/service"
	"smartpatrol/pkg/helpers"
	"smartpatrol/pkg/util/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controllerImpl struct {
	Service service.Service
}

func NewController(Service service.Service) *controllerImpl {
	return &controllerImpl{
		Service: Service,
	}
}

func (co *controllerImpl) GetAllUser(c echo.Context) error {
	result, err := co.Service.GetAllUser(c)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Users Active", result)
}

func (co *controllerImpl) GetUserByID(c echo.Context) error {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrBadRequest, err))
	}

	result, err := co.Service.GetUserByID(c, id)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get User By ID", result)
}

func CreateUser(c echo.Context) error {
	data := users.UserRequest{}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := helpers.HashedPassword(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err := repository.CreateUser(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success created user", "")

}

func UpdateUser(c echo.Context) error {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("invalid user id"))
	}

	userUpdate := users.UserRequest{}
	if err := c.Bind(&userUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if userUpdate.Password != "" {
		if err := helpers.HashedPassword(&userUpdate); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	err = repository.UpdateUser(id, userUpdate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Successfully updated user", "")
}


func DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid user id"))
	}

	err = repository.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Successfully deleted user", "")
}