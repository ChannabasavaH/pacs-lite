package routes

import (
	"auth-service/internal/handlers"

	"github.com/gin-gonic/gin"
	"auth-service/internal/middleware"
)

func SetUpAuthRoutes (router *gin.Engine, authHandler *handlers.AuthHandlers){
	api := router.Group("api/auth")
	{
		api.POST("/signup", authHandler.SignUp)
		api.POST("/login", authHandler.LoginUser)
	}

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())

	{
		protected.POST("/logout", authHandler.Logout)
	}
}