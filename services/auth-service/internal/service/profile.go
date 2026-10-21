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

func (r *ProfileService) CreateProfile(profile *models.Profile) (*models.Profile, error) {
	if err := r.Repository.CreateProfile(profile); err != nil {
		return nil, err
	}

	return profile, nil
}