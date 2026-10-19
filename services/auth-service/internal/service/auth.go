package service

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"

)

var ErrInvalidCredentials = errors.New("Invalid Email or Password")

type AuthService struct{
	Repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService{
	return &AuthService{
		Repo: repo,
	}
}

func (s *AuthService) SignUpUser(account *models.Account) (*models.Account, error){
	if err := s.Repo.SignUp(account); err != nil {
		return nil, err
	}
	return account, nil
}

func (s *AuthService) LoginUser(email string, password string) (*models.Account, error){
	account, err := s.Repo.Login(email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInvalidCredentials
	}

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return account, nil
}