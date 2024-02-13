package handlers

import (
	"smartpatrol/app/company/controller"
	"smartpatrol/app/company/repository"
	"smartpatrol/app/company/service"
	"smartpatrol/pkg/database"

	"github.com/labstack/echo/v4"
)

type handlerCompany struct {
	Controller controller.Controller
}

func WebCompanyHandler() *handlerCompany {
	s := service.NewService(database.DBManager(), repository.NewRepository())

	return &handlerCompany{
		Controller: controller.NewController(s),
	}
}

func (h *handlerCompany) WebRoute(g *echo.Group) {
	g.GET("", h.Controller.GetAll)
	g.GET("/:id", h.Controller.GetById)
	g.POST("", controller.CreateCompany)
	g.PUT("/:id", controller.UpdateCompany)
	g.DELETE("/:id", controller.DeleteCompany)
	g.POST("/isactive/:id", controller.IsActiveCompany)
}