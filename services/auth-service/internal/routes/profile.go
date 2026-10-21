package routes

import (
	"github.com/gin-gonic/gin"
	"auth-service/internal/handlers"
	"auth-service/internal/middleware"
)

func SetUpProfileRoutes(router *gin.Engine, profileHandlers *handlers.ProfileHandler){
	api := router.Group("/api")
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.POST("/create-profile", profileHandlers.CreateProfile)
	}
}