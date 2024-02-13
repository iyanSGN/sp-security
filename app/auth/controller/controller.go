package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	Login(c echo.Context) error
}