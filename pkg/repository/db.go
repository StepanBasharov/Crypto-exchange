package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitCursor() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Database not working")
	}
	return db

}
