package handlers

import (
	"smartpatrol/app/users/controller"
	"smartpatrol/app/users/repository"
	"smartpatrol/app/users/service"
	"smartpatrol/pkg/database"

	"github.com/labstack/echo/v4"
)

type handlerUsers struct {
	Controller controller.Controller
}

func WebUsersHandler() *handlerUsers {
	s := service.NewService(database.DBManager(), repository.NewRepository())

	return &handlerUsers{
		Controller: controller.NewController(s),
	}
}

func (h *handlerUsers) WebRoute(g *echo.Group) {
	g.GET("", h.Controller.GetAllUser)
	g.GET("/:id", h.Controller.GetUserByID)
	g.POST("", controller.CreateUser)
	g.PUT("/:id", controller.UpdateUser)
	g.DELETE("/:id", controller.DeleteUser)

}

