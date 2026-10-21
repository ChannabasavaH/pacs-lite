package validations

import (
	"auth-service/internal/models"
	"github.com/go-playground/validator/v10"
	"auth-service/internal/types"
)

type Profile struct {
	FirstName string `json:"firstName" binding:"required,min=2,max=20"`
	LastName string `json:"lastName" binding:"required,min=2,max=20"`
	Gender models.Gender `json:"gender" binding:"required,oneof=MALE"`
	DOB types.DateOnly `json:"dob" binding:"required"`
}

func ValidateProfile(profile *Profile) error {
	validate := validator.New()
	return validate.Struct(profile)
}