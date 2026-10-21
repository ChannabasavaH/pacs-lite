package handlers

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"auth-service/internal/validations"
	"net/http"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	Service *service.ProfileService
}

func NewProfileHandler(service *service.ProfileService) *ProfileHandler{
	return	&ProfileHandler{
		Service: service,
	}
}

func (h *ProfileHandler) CreateProfile(c *gin.Context){
	var request validations.Profile

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validations.ValidateProfile(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uuid := uuid.New()
	accountIdValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthrozied",
		})
		return
	}

	accountID, ok := accountIdValue.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	var profile = &models.Profile{
		ID: uuid.String(),
		AccountID: accountID,
		FirstName: request.FirstName,
		LastName: request.LastName,
		Gender: request.Gender,
		DOB: request.DOB.Time,
	}

	profile, err := h.Service.CreateProfile(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"profile": profile,
	})
}