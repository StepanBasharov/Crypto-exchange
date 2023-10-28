package models

type Crypto struct {
	ID       int    `json:"ID" gorm:"primary_key;autoIncrement"`
	CryptoID int    `json:"cryptoID"`
	Symbol   string `json:"symbol"`
}

type Wallet struct {
	ID           int      `json:"ID" gorm:"primary_key;autoIncrement"`
	UserID       int      `json:"user_id"`
	Address      string   `json:"address"`
	CryptoWallet []Crypto `json:"crypto_wallet" gorm:"foreignKey:CryptoID"`
	Balance      string   `json:"balance"`
}
