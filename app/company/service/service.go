package service

import (
	"smartpatrol/app/company"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]company.CompanyResponse, error)
	GetById(c echo.Context, id uint) (company.CompanyResponse, error)
}