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

	// Devices
	protected.GET("/devices", apiController.GetAllDevice)
	protected.GET("/devices_issues", apiController.GetAllDeviceIssues)

	// Philippine Addresses

	protected.GET("/regions", apiController.GetRegions)
	protected.GET("/regions/:region_id/provinces", apiController.GetProvinces)
	protected.GET("/provinces/:province_id/municipalities", apiController.GetMunicipalities)
	protected.GET("/municipalities/:municipality_id/barangays", apiController.GetBarangays)

	// SuperAdmin
	protected.POST("/organizations", apiController.CreateOrganization)
	protected.GET("/organizations", apiController.GetOrganizations)
	protected.POST("/organizations/:organization_id/existing-user/:user_id", apiController.InviteToOrganizationUser)
	protected.POST("/organizations/:organization_id/members", apiController.CreateOrganizationUser)
	protected.GET("/organizations/:organization_id/members", apiController.GetOrganizationUsers)

	// Users
	protected.GET("/users", apiController.GetUsers)

	// Admin

	// Transactions
	protected.POST("/transactions", apiController.CreateTransaction)
	protected.GET("/transactions", apiController.GetTransactions)
	protected.GET("/transactions/:id", apiController.GetTransaction)
	protected.PUT("/transactions/:id", apiController.UpdateTransaction)
	protected.DELETE("/transactions/:id", apiController.DeleteTransaction)

	// Technician assignment routes
	protected.POST("/transactions/:id/assign", apiController.AssignTechnician)
	protected.POST("/transactions/:id/unassign/:techId", apiController.UnassignTechnician)
	protected.GET("/transactions/:id/technicians", apiController.GetTechnicianAssignments)

}
