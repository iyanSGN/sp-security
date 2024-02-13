package service

import (
	"smartpatrol/app/role"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAllUser(c echo.Context) ([]role.RoleResponse, error)
}