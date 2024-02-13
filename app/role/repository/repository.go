package repository

import (
	"smartpatrol/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllUser(c echo.Context, DB *gorm.DB) ([]models.SecurityRole, error)
}