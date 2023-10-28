package controllers

import (
	"cex/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	response := models.Response{Message: "OK"}
	return c.JSON(http.StatusOK, response)
}
