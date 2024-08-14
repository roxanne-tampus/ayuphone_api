package controllers

import (
	"net/http"

	"ayuphone_api/internal/services"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	// Extract user ID from the context set by AuthMiddleware
	userIdString, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, _ := (userIdString).(int64)

	// Fetch user from the database
	user, err := services.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"role":      user.Role,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
	})
}
