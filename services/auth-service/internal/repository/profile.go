package repository

import (
	"auth-service/internal/models"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{
		DB: db,
	}
}

func (r *ProfileRepository) CreateProfile(profile *models.Profile) error {
	return r.DB.Create(profile).Error
}



