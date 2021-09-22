package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnect *gorm.DB

func DB() *gorm.DB {
	dbConnect, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("DB error")
	}
	return dbConnect
}
