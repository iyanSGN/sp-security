package repository

import (
	"smartpatrol/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllUser(c echo.Context,db *gorm.DB) ([]models.SecurityUser, error)
	GetUserByID(c echo.Context,db *gorm.DB, id int) (models.SecurityUser, error)
}