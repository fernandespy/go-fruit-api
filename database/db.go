package database

import (
	"go-fruit-api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("fruits.db?_pragma=busy_timeout=5000"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db

	err = DB.AutoMigrate(&models.Fruit{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
