package handlers

import (
	"smartpatrol/app/role/controller"
	"smartpatrol/app/role/repository"
	"smartpatrol/app/role/service"
	"smartpatrol/pkg/database"

	"github.com/labstack/echo/v4"
)

type handlerRole struct {
	Controller controller.Controller
}

func WebRoleHandler() *handlerRole {
	s := service.NewService(database.DBManager(), repository.NewRepository())

	return &handlerRole{
		Controller: controller.NewController(s),
	}
}

func (h *handlerRole) WebRoute(g *echo.Group) {
	g.GET("", h.Controller.GetAllUser)

}

