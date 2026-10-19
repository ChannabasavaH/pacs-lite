package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Port string
	DBName string
	DBHost string
	DBUser string
	DBPassword string
}

type APPConfig struct {
	Port string
}

func APPLoadConfig() APPConfig {
	err := godotenv.Load()

	if err != nil {
		log.Println("Failed to load app env")
	}

	return APPConfig{
		Port: os.Getenv("APP_PORT"),
	}
}

func DBLoadConfig() DBConfig {
	err := godotenv.Load()

	if err != nil {
		log.Println("Failed to load env")
	}

	return DBConfig{
		Port: os.Getenv("PORT"),
		DBName: os.Getenv("DB_NAME"),
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}