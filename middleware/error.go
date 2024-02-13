package middleware

import (
	"net/http"

	res "smartpatrol/pkg/util/response"

	"github.com/labstack/echo/v4"
)

func NewErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusFound, err.Error())
	}

	switch report.Code {
	case http.StatusNotFound:
		err = res.BuildCustomError(http.StatusNotFound, "Route Not Found")
	case http.StatusInternalServerError:
		err = res.BuildError(res.ErrServerError, err)
	default:
		err = res.BuildError(res.ErrServerError, err)
	}

	res.ErrorResponse(c, err)
}