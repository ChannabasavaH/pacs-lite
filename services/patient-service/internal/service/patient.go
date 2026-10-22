package service

import (
	"patient-service/internal/models"
	"patient-service/internal/repository"
)

type PatientService struct {
	Repository *repository.PatientRepository
}

func NewPatientService(Repo *repository.PatientRepository) *PatientService {
	return &PatientService{
		Repository: Repo,
	}
}

func (s *PatientService) CreatePatient(patient *models.Patient) (*models.Patient, error) {
	err := s.Repository.CreatePatient(patient)

	if err != nil {
		return nil, err
	}

	return patient, nil
}