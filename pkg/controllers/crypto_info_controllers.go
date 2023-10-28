package controllers

import (
	"cex/pkg/models"
	"cex/pkg/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllCrypto(c echo.Context) error {
	var crypto []models.Crypto
	db := repository.InitCursor()
	result := db.Find(&crypto)
	if result.Error != nil {
		panic(result.Error)
	}
	return c.JSON(http.StatusOK, crypto)

}

func GetCryptoByID(c echo.Context) error {
	var crypto models.Crypto
	id := c.Param("id")
	db := repository.InitCursor()
	db.First(&crypto, id)
	if crypto.ID != 0 {
		return c.JSON(http.StatusOK, crypto)
	} else {
		return c.JSON(http.StatusNotFound, models.Response{Message: "Not Found"})
	}
}
