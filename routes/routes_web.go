package routes

import (
	"net/http"
	"smartpatrol/pkg/util/environment"
	"smartpatrol/routes/handlers"

	// "serelo-backend-golang/app/otp/controller"
	// "serelo-backend-golang/pkg/util/enviroment"
	// "serelo-backend-golang/routes/handlers"

	"github.com/labstack/echo/v4"
)

func WebInit(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to "+environment.Get("APP_NAME")+"! version "+environment.Get("APP_VERSION")+" in mode "+environment.Get("ENV"))
	})

	// //OTP
	// g.POST("/generateotp", controller.GenerateOTP)
	// g.POST("/resendotp", controller.ResendOTP)
	// g.POST("/verifyotp", controller.VerifyOtp)



	// // Routes
	handlers.WebAuthHandler().WebRoute(g.Group("/web/auth"))
	handlers.WebCompanyHandler().WebRoute(g.Group("/web/company"))
	handlers.WebUsersHandler().WebRoute(g.Group("/web/user"))
	handlers.WebRoleHandler().WebRoute(g.Group("/web/role"))
	// handlers.WebUserHandler().WebRoute(g.Group("/web/user"))
	// handlers.WebHandlerApproval().WebRoute(g.Group("/web/approval"))
	// handlers.WebUplaodHandler().WebRoute(g.Group("/web/upload"))
	// handlers.WebBuildingHandler().WebRoute(g.Group("/web/building"))
	// handlers.WebHandlerFloor().WebRoute(g.Group("/web/floor"))
	// handlers.WebHandlerDepartment().WebRoute(g.Group("/web/department"))
	// handlers.WebHandlerJobTitle().WebRoute(g.Group("/web/jobtitle"))
	// handlers.WebHandlerCamera().WebRoute(g.Group("/web/camera"))
	// handlers.WebHandlersMeeting().WebRoute(*g.Group("/web/meeting"))
	// handlers.WebHandlerDesk().WebRoute(g.Group("/web/desk"))
	// handlers.WebHandlerPackage().WebRoute(g.Group("/web/package"))
	// handlers.WebHandlerVisitRes().WebRoute(g.Group("/web/visitres"))
	// handlers.WebHandlerDeskReservation().WebRoute(g.Group("/web/deskres"))
	// handlers.WebHandlerRoomReservation().WebRoute(g.Group("/web/roomres"))

}