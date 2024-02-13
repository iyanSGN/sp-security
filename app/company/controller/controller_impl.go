package controller

import (
	"fmt"
	"net/http"
	"smartpatrol/app/company"
	"smartpatrol/app/company/repository"
	"smartpatrol/app/company/service"
	"smartpatrol/models"
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

func (co *controllerImpl) GetAll(c echo.Context) error {
	result, err := co.Service.GetAll(c)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get All Company", result)
}

func (co *controllerImpl) GetById(c echo.Context) error {
	compId := c.Param("id")

	id, err := strconv.ParseUint(compId, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrBadRequest, err))
	}

	result, err := co.Service.GetById(c, uint(id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get Companny By ID", result)
}

func CreateCompany(c echo.Context) error {
	var req company.CompanyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "bad request",
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
		})
	}

	_, err := repository.CreateCompany(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":     "Failed to create user",
			"error":       err.Error(),
			"status_code": http.StatusInternalServerError,
		})
	}

	return response.SuccessResponse(c, http.StatusOK, "Successfully created company", "")
}

func UpdateCompany(c echo.Context) error {
	compID := c.Param("id")
	id, err := strconv.Atoi(compID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":     "Invalid company ID",
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
		})
	}

	UpdatedCompany := company.CompanyRequest{}
	if err := c.Bind(&UpdatedCompany); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = repository.UpdateCompany(id, UpdatedCompany)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":     "Failed to update company",
			"error":       err.Error(),
			"status_code": http.StatusInternalServerError,
		})
	}

	return response.SuccessResponse(c, http.StatusOK, "Successfully Updating Company", "")
}


func DeleteCompany(c echo.Context) error {
	compID := c.Param("id")
	id, err := strconv.Atoi(compID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid company id"))
	}

	err = repository.DeleteCompany(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Deleting Company", "")
}

func IsActiveCompany(c echo.Context) error {
	companyID := c.Param("id")
	id, err := strconv.ParseUint(companyID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid building id"))
	}
	var company models.CompanyCompany

	err = repository.IsActiveCompany(uint(id), &company)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	message := "company isnot active"
	if company.IsActive == 1 {
		message = "company is active"
	}

	fmt.Println("================================")
	fmt.Println(company.IsActive)

	return response.SuccessResponse(c, http.StatusOK, "Success", message)
}
