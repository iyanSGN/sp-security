package main

import (
	"fmt"
	"net/http"
	"os"
	util "smartpatrol/pkg"
	"smartpatrol/pkg/database"
	"smartpatrol/pkg/util/environment"
	"smartpatrol/routes"

	myMiddleware "smartpatrol/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	environment.Init(".env.company")
	environment.Init(".env.security")
	database.Init("postgresql")
	database.Migrate()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}),

		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: fmt.Sprintf("\n%s | ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${method} | ${uri} ",
				environment.Get("APP_NAME"),
			),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
	)
	e.HTTPErrorHandler = myMiddleware.NewErrorHandler
	e.Validator = &util.CustomValidation{Validator: validator.New()}

	// Route
	routes.MobileInit(e.Group("api/v1"))
	routes.WebInit(e.Group("api/v1"))

	e.Logger.Fatal(e.Start(":" + environment.Get("APP_PORT")))
}
