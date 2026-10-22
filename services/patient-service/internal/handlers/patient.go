package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"patient-service/internal/models"
	"patient-service/internal/service"
	"patient-service/internal/validations"
	"net/http"
)

type PatientHandler struct {
	Service *service.PatientService
}

func NewPatientHandler(service *service.PatientService) *PatientHandler {
	return &PatientHandler{
		Service: service,
	}
}

func (h *PatientHandler) CreatePatient(c *gin.Context){
	var request validations.Patient

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validations.ValidatePatient(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accIDValue, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	accountID, ok := accIDValue.(string)

	if !ok || accountID == " " {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	id := uuid.New().String()

	patient := &models.Patient{
		ID: id,
		MRN: request.MRN,
		FirstName: request.FirstName,
		LastName: request.LastName,
		Gender: request.Gender,
		DOB: request.DOB.Time,
		Phone: request.Phone,
		Email: request.Email,
		Address: request.Address,
		CreatedBy: accountID,
		IsActive: request.IsActive,
	}

	patient, err := h.Service.CreatePatient(patient)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"patient": patient,
	})
}