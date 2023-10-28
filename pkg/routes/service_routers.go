package routes

import (
	"cex/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterServiceRouters(router *echo.Echo) {
	router.GET("/healthcheck", controllers.HealthCheck)
}
