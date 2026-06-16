package database

import (
	"app/internal/models"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connections error", err)
	}

	db.AutoMigrate(&models.User{})
	DB = db
}