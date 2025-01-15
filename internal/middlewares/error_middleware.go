package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
}
