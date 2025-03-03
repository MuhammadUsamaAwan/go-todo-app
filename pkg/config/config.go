package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbURL string
}

func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}

	return Config{
		DbURL: dbURL,
	}
}
