package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBUser string
	DBName string
	DBPassword string
	DBPort string
	DBHost string
}

type APPConfig struct {
	APPPort string
}

func DBLoadConfig() DBConfig {

	err := godotenv.Load()

	if err != nil {
		log.Print("Failed to load database config")
	}

	return DBConfig{
		DBUser: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort: os.Getenv("DB_PORT"),
		DBHost: os.Getenv("DB_HOST"),
	}
}

func APPLoadConfig() APPConfig {
	err := godotenv.Load()

	if err != nil {
		log.Println("Failed to load app env")
	}

	return APPConfig{
		APPPort: os.Getenv("APP_PORT"),
	}
}