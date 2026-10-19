package validations

import (
	"auth-service/internal/models"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Username    string             `json:"username" binding:"required,min=3, max=50"`
	Email       string             `json:"email" binding:"required,email"`
	Password    string             `json:"password" binding:"required,min=8"`
	AccountType models.AccountType `json:"account_type" binding:"required,oneof=RADIOLOGIST"`
}

func ValidateUser(user *User) error {
	validate := validator.New()
	return validate.Struct(user)
}

