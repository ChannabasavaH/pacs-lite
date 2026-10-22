package database

import (
	"patient-service/internal/config"
	"gorm.io/gorm"
	"fmt"
	"gorm.io/driver/postgres"
	"log"
)

func ConnectPostgres(cfg config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to database: %v", err.Error())
	}

	log.Println("Connected to database")

	return db
}


