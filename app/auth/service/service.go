package service

import (
	"smartpatrol/app/auth"

	"github.com/labstack/echo/v4"
)

type Service interface {
	Login(c echo.Context, r auth.LoginRequest) (auth.LoginResponse, error)
}