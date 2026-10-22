package main

import (
	"fmt"
	"patient-service/internal/config"
	"patient-service/internal/database"
	"patient-service/internal/handlers"
	"patient-service/internal/repository"
	"patient-service/internal/routes"
	"patient-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main(){
	dbCfg := config.DBLoadConfig()
	appCfg := config.APPLoadConfig()
	db := database.ConnectPostgres(dbCfg)
	patientRepo := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepo)
	patientHandlers := handlers.NewPatientHandler(patientService)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	jwtAuth := existingJWTMiddleware()

	routes.SetUpPatientRoutes(router, patientHandlers, jwtAuth)

	port := appCfg.APPPort

	if port == "" {
		port = "8000"
	}

	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
	}
}