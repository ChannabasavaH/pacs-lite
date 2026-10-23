package main

import (
	"fmt"
	"log"
	"patient-service/internal/config"
	"patient-service/internal/database"
	"patient-service/internal/handlers"
	"patient-service/internal/repository"
	"patient-service/internal/routes"
	"patient-service/internal/service"
	"patient-service/internal/utils"
	"patient-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main(){
	dbCfg := config.DBLoadConfig()
	appCfg := config.APPLoadConfig()
	jwtConfig := config.JWTLoadConfig()
	db := database.ConnectPostgres(dbCfg)
	patientRepo := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepo)
	patientHandlers := handlers.NewPatientHandler(patientService)

	publicKey, err := utils.LoadPublicKey(jwtConfig.JWT_Public_Key_Path)
	if err != nil {
		log.Fatal("Failed to laod jwt public key: ", err)
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	routes.SetUpPatientRoutes(router, patientHandlers, middleware.AuthMiddleware(publicKey))

	port := appCfg.APPPort

	if port == "" {
		port = "8000"
	}

	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
	}
}