package main

import (
	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/handlers"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/internal/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){

	cfg := config.DBLoadConfig()

	appcfg := config.APPLoadConfig()

	db := database.ConnectPostgres(cfg)

	repo := repository.NewSignUpRepository(db)

	service := service.NewAuthService(repo)

	handlers := handlers.NewAuthHandlers(service)

	router := gin.Default()

	routes.SetUpRoutes(router, handlers)

	fmt.Println("Welcome to pacs lite")

	port := appcfg.Port
	if port == "" {
		port = "8000"
	}

	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
	}
}

