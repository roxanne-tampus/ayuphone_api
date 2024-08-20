package routes

import (
	"ayuphone_api/internal/controllers"
	middleware "ayuphone_api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, apiController controllers.ApiController) {
	auth := router.Group("auth")
	{
		auth.POST("/register", apiController.Register)
		auth.POST("/login", apiController.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware)
	protected.GET("/me", apiController.GetProfile)

	protected.POST("/transactions", apiController.CreateTransaction)
	protected.GET("/transactions/:id", apiController.GetTransaction)
	protected.PUT("/transactions/:id", apiController.UpdateTransaction)
	protected.DELETE("/transactions/:id", apiController.DeleteTransaction)

	// Technician assignment routes
	protected.POST("/transactions/:id/assign", apiController.AssignTechnician)
	protected.POST("/transactions/:id/unassign/:techId", apiController.UnassignTechnician)
	protected.GET("/transactions/:id/technicians", apiController.GetTechnicianAssignments)
}
