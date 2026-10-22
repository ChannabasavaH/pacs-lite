package routes

import (
	"patient-service/internal/handlers"
	"github.com/gin-gonic/gin"
	"patient-service/internal/middleware"
)

func SetUpPatientRoutes(router *gin.Engine, patientHandlers *handlers.PatientHandler, jwtAuth gin.HandlerFunc){
	api := router.Group("/api/patients")
	protected := api.Group("")
	protected.Use(jwtAuth)
	protected.Use(middleware.RequirePatientStaff())

	{
		protected.POST("", patientHandlers.CreatePatient)
	}
}

