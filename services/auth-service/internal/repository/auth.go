package repository

import (
	"auth-service/internal/models"
	"errors"

	"gorm.io/gorm"
)

var ErrAccountExists = errors.New("email or username already exists")

type AuthRepository struct {
	DB *gorm.DB
}

func NewSignUpRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (r *AuthRepository) SignUp(account *models.Account) error {
	var count int64
	err := r.DB.Model(&models.Account{}).
		Where(
			"emial = ? OR username = ?",
			account.Email,
			account.Username,
		).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return ErrAccountExists
	}

	return r.DB.Create(account).Error
}


