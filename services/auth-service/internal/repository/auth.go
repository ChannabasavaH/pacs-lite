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
			"email = ? OR username = ?",
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

func (r *AuthRepository) Login(email string) (*models.Account, error) {
	var existingAccount models.Account
	err := r.DB.Model(&models.Account{}).
		Where("email = ?",
	email).First(&existingAccount).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email/username or password")
		}

		return nil, err
	}
	return &existingAccount, nil
}

