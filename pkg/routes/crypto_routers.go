package routes

import (
	"cex/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterCryptoRouters(router *echo.Echo) {
	router.GET("/crypto", controllers.GetAllCrypto)
	router.GET("/crypto/:id", controllers.GetCryptoByID)
}
