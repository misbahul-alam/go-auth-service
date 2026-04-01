package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT       string
	DB_DSN     string
	JWT_SECRET string
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	DB_DSN = os.Getenv("DB_DSN")
	JWT_SECRET = os.Getenv("JWT_SECRET")
}
