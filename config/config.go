package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SERVER_ADDRESS string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitConfiguration() Config {
	return Config{
		SERVER_ADDRESS: os.Getenv("SERVER_ADDRESS"),
		DB_USERNAME:    os.Getenv("DB_USERNAME"),
		DB_PASSWORD:    os.Getenv("DB_PASSWORD"),
		DB_PORT:        os.Getenv("DB_PORT"),
		DB_HOST:        os.Getenv("DB_HOST"),
		DB_NAME:        os.Getenv("DB_NAME"),
	}
}
