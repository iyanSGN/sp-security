package repository

import (
	"smartpatrol/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, DB *gorm.DB) ([]models.CompanyCompany, error)
	GetById(c echo.Context, DB *gorm.DB, ID uint) (models.CompanyCompany, error)
}