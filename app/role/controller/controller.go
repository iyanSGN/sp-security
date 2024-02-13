package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	GetAllUser(c echo.Context) error
}