package db

import (
	"log"
	"url-shortner/db/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("shortner.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database due to", err)
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Urls{})
}
