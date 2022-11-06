package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("postgres", "host=localhost port=5436 user=postgres dbname=musicDB password=qwerty sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Track{})
	DB = db
}
