package routes

import (
	"ayuphone_api/internal/controllers"
	middlewares "ayuphone_api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	auth := router.Group("auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware)
	protected.GET("/me", controllers.GetProfile)
}
