package migrations

import (
	"auth-service/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Account{},
		&models.Profile{},
	)
}