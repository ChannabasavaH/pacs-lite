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

type JWTConfig struct {
	JWT_Secret string
	JWT_Private_Key_Path string
	JWT_Public_Key_Path string
}

func JWTLoadConfig() JWTConfig {
	err := godotenv.Load()

	if err != nil {
		log.Println("Failed to load jwt secret")
	}

	return JWTConfig{
		JWT_Secret: os.Getenv("JWT_SECRET"),
		JWT_Private_Key_Path: os.Getenv("JWT_PRIVATE_KEY_PATH"),
		JWT_Public_Key_Path: os.Getenv("JWT_PUBLIC_KEY_PATH"),
	}
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