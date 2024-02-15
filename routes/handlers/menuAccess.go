package handlers

import (
	"smartpatrol/app/menuAccess/controller"
	"smartpatrol/app/menuAccess/repository"
	"smartpatrol/app/menuAccess/service"
	"smartpatrol/pkg/database"

	"github.com/labstack/echo/v4"
)

type handlerMenuAccess struct {
	Controller controller.Controller
}

func WebMenuAccessHandler() *handlerMenuAccess {
	s := service.NewService(database.DBManager(), repository.NewRepository())

	return &handlerMenuAccess{
		Controller: controller.NewController(s),
	}
}

func (h *handlerMenuAccess) WebRoute(g *echo.Group) {
	g.GET("", h.Controller.GetAllMenu)
	g.GET("/:id", h.Controller.GetMenuByID)
	g.POST("", controller.CreateMenuAccess)
	g.PUT("/:id", controller.UpdateMenuAccess)
	g.DELETE("/:id", controller.DeleteMenuAccess)

}