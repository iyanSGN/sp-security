package handlers

import (
	"smartpatrol/app/auth/controller"
	"smartpatrol/app/auth/repository"
	"smartpatrol/app/auth/service"
	"smartpatrol/pkg/database"

	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	Controller controller.Controller
}

func WebAuthHandler() *handlerAuth {
	s := service.NewService(database.DBManager(), repository.NewRepository())

	return &handlerAuth{
		Controller: controller.NewController(s),
	}
}

func (h *handlerAuth) WebRoute(g *echo.Group) {
	g.POST("/login", h.Controller.Login)
	g.POST("/user", controller.CreateAdmin)
}

func MobileAuthHandler() *handlerAuth {
	s := service.NewService(database.DBManager(), repository.NewRepository())

	return &handlerAuth{
		Controller: controller.NewController(s),
	}
}

func (h *handlerAuth) MobileRoute(g *echo.Group) {
	g.POST("/login", h.Controller.Login)
	g.POST("/user", controller.CreateAdmin)
}