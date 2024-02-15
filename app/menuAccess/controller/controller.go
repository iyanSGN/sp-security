package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	GetAllMenu(c echo.Context) error
	GetMenuByID(c echo.Context) error
}