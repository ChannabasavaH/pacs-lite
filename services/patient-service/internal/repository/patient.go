package repository

import (
	"patient-service/internal/models"
	"gorm.io/gorm"
)

type PatientRepository struct{
	DB *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository{
	return &PatientRepository{
		DB: db,
	}
}

func (r *PatientRepository) CreatePatient(patient *models.Patient) error {
	return r.DB.Create(patient).Error
}

