package controller

import (
	"fmt"
	"net/http"
	menuaccess "smartpatrol/app/menuAccess"
	"smartpatrol/app/menuAccess/repository"
	"smartpatrol/app/menuAccess/service"
	"smartpatrol/pkg/util/response"
	"strconv"

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

func (co *controllerImpl) GetAllMenu(c echo.Context) error {
	result, err := co.Service.GetAllMenu(c)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Menu Access", result)
}

func (co *controllerImpl) GetMenuByID(c echo.Context) error {
	menUID := c.Param("id")

	id, err := strconv.Atoi(menUID)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	result, err := co.Service.GetMenuByID(c, id)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get Menu Access By id", result)
}


func CreateMenuAccess(c echo.Context) error {
	var per menuaccess.MenuAccessReq
	if err := c.Bind(&per);
	err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err := repository.CreateMenuAccess(per)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Create Menu Access", "")
}

func UpdateMenuAccess(c echo.Context) error {
	menuID := c.Param("id")
	id, err := strconv.Atoi(menuID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid id"))
	}

	updatedMenuAccess := menuaccess.MenuAccessReq{}
	if err := c.Bind(&updatedMenuAccess);
	err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = repository.UpdateMenuAccess(id, updatedMenuAccess)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Menu Access Successfully Updated",
		"status_code": http.StatusOK,
	})
}

func DeleteMenuAccess(c echo.Context) error {
	menuID := c.Param("id")
	id, err := strconv.Atoi(menuID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid user id"))
	}

	err = repository.DeleteDepartment(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "Menu Access deleted successfully",
		"status_code" : http.StatusOK,
	})
}


