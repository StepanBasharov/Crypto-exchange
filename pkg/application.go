package pkg

import (
	"cex/pkg/models"
	"cex/pkg/repository"
	"cex/pkg/routes"
	"fmt"
	"github.com/labstack/echo/v4"
)

func StartServer(port string) {
	app := echo.New()
	db := repository.InitCursor()
	db.AutoMigrate(models.User{}, models.Crypto{})
	repository.CreateCrypto(db)
	routes.RegisterServiceRouters(app)
	routes.RegisterCryptoRouters(app)
	fmt.Println("Run ")
	app.Start(port)
	fmt.Println("RUn....")
}
