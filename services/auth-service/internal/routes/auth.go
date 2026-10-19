package routes

import (
	"auth-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes (router *gin.Engine, authHandler *handlers.AuthHandlers){
	api := router.Group("api/auth")
	{
		api.POST("/signup", authHandler.SignUp)
	}
}