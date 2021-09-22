package creditcard

import (
	db "github.com/abc7468/rest.go/DB"
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	Id   int
	name string
}

func makeCreditCard() {
	db := db.DB()
	db.AutoMigrate(&CreditCard{})
	db.Create(&CreditCard{Id: 2, name: "royJang"})
}
