package routes

import (
	"net/http"
	"smartpatrol/pkg/util/environment"
	"smartpatrol/routes/handlers"

	// "smartpatrol/app/otp/controller"
	// "smartpatrol/pkg/util/enviroment"
	// "smartpatrol/routes/handlers"

	"github.com/labstack/echo/v4"
)

func MobileInit(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to "+environment.Get("APP_NAME")+"! version "+environment.Get("APP_VERSION")+" in mode "+environment.Get("ENV"))
	})

	handlers.MobileAuthHandler().MobileRoute(g.Group("/mobile/auth"))


}