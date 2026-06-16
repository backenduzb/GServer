package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL string
	PORT string
	JWT_SECRET string
}

var Envs Config

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env not found")
	}

	Envs = Config{
		DB_URL: os.Getenv("DATABASE_URL"),
		PORT: os.Getenv("PORT"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
}
