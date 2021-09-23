package creditcard

import (
	"log"

	db "github.com/abc7468/rest.go/DB"
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	Id   int
	Name string
}

func MakeCreditCard(id int, name string) error {
	db, err := db.DB()
	if err != nil {
		log.Print(err)
		return err
	}
	db.AutoMigrate(&CreditCard{})
	db.Create(&CreditCard{Id: id, Name: name})
	return nil
}
