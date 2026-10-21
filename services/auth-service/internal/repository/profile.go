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

func (r *ProfileRepository) GetProfile(accountID string) (*models.Profile, error) {
	var profile *models.Profile

	err := r.DB.Preload("Account", func(db *gorm.DB) *gorm.DB{
		return db.Select("id", "username", "email", "account_type")
	}).
		Where("account_id = ?", accountID).
		First(&profile).Error

	if err != nil {
		return nil, err
	}

	return profile, nil
}



