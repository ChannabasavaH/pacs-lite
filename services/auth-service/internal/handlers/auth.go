package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"errors"

	"auth-service/internal/validations"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlers struct {
	Service *service.AuthService
}

func NewAuthHandlers(service *service.AuthService) *AuthHandlers {
	return &AuthHandlers{
		Service: service,
	}
}

func (h *AuthHandlers) SignUp(c *gin.Context) {

	var request validations.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid body request",
		})
		return
	}

	if err:= validations.ValidateUser(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return 
	}

	account := &models.Account{
		ID: id.String(),
		Username: request.Username,
		Email: request.Email,
		Password: string(hashedPassword),
		AccountType: request.AccountType,
	}

	user, err := h.Service.SignUpUser(account)

	if err != nil {
		if errors.Is(err, repository.ErrAccountExists){
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user": user,
	})

}
