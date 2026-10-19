package service

import (
	"auth-service/internal/repository"
	"auth-service/internal/models"
)

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