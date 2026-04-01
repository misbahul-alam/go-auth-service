package database

import (
	"log"

	"github.com/misbahul-alam/go-auth-service/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open(config.DB_DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB = db
}
