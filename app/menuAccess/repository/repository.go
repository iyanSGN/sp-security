package repository

import (
	"smartpatrol/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllMenu(c echo.Context, DB *gorm.DB) ([]models.SecurityPermission, error)
	GetMenuByID(c echo.Context, DB *gorm.DB, ID int) (models.SecurityPermission, error)
}