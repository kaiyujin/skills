package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Health(c echo.Context) error {
	result := map[string]interface{}{
		"status": "OK",
	}
	return c.JSON(http.StatusOK, result)
}
