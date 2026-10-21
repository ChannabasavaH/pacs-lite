package service

import (
	"auth-service/internal/repository"
	"auth-service/internal/models"
)

type ProfileService struct {
	Repository *repository.ProfileRepository
}

func NewProfileService(Repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{
		Repository: Repo,
	}
}

func (s *ProfileService) CreateProfile(profile *models.Profile) (*models.Profile, error) {
	if err := s.Repository.CreateProfile(profile); err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *ProfileService) GetProfile(accountID string) (*models.Profile, error){
	profile, err := s.Repository.GetProfile(accountID)

	if err != nil {
		return nil, err
	}

	return profile, nil
}