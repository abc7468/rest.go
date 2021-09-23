package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnect *gorm.DB

func DB() (*gorm.DB, error) {
	dbConnect, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return dbConnect, nil
}
