package validations

import (
	"github.com/go-playground/validator/v10"
	"patient-service/internal/models"
	"patient-service/internal/types"
)

type Patient struct {
	MRN string `json:"mrn" binding:"required"`
	FirstName string `json:"firstName" binding:"required;min=2;max=20"`
	LastName string `json:"lastName" binding:"required;min=2;max=20"`
	Gender models.Gender `json:"gender" binding:"required"`
	DOB types.DateOnly `json:"dob" binding:"required"`

	Phone string `json:"phone" binding:"required;len=13"`
	Email string `json:"email" binding:"required;email"`
	Address string `json:"address" binding:"required;min=2;max=256"`

	IsActive bool `json:"is_active"`
}

func validatePatient(patient *Patient) error {
	validate := validator.New()
	return validate.Struct(patient)
} 