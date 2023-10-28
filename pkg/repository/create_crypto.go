package repository

import (
	"cex/pkg/models"
	"gorm.io/gorm"
)

func CreateCrypto(db *gorm.DB) {
	var crypto []models.Crypto
	crypto = append(
		crypto,
		models.Crypto{CryptoID: 0, Symbol: "BTC"},
		models.Crypto{CryptoID: 1, Symbol: "ETH"},
		models.Crypto{CryptoID: 2, Symbol: "XMR"},
	)

	for _, item := range crypto {
		db.Create(&item)
	}
}
