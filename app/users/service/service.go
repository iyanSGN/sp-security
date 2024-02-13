package service

import (
	"smartpatrol/app/users"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAllUser(c echo.Context) ([]users.UserResponse, error)
	GetUserByID(c echo.Context, id int) (users.UserResponse, error)
}