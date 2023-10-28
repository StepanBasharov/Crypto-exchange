package models

import "time"

type User struct {
	ID            string    `json:"ID" gorm:"primary_key;autoIncrement"`
	UserName      string    `json:"username"`
	Email         string    `json:"email"`
	UserPassword  string    `json:"password"`
	UserPhoto     string    `json:"photo"`
	UserBirthDate time.Time `json:"birthdate"`
	Wallet        []Wallet  `json:"wallet" gorm:"foreignKey:UserID"`
}
