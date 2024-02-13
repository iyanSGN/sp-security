package repository

import (
	"smartpatrol/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	Login(c echo.Context, DB *gorm.DB, email string) ( models.SecurityAccount, error)
}