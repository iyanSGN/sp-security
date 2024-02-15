package service

import (
	menuaccess "smartpatrol/app/menuAccess"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAllMenu(c echo.Context) ([]menuaccess.MenuAccessRes, error)
	GetMenuByID(c echo.Context, id int) (menuaccess.MenuAccessRes, error)
}