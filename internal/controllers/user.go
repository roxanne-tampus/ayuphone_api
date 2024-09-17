package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac ApiController) GetProfile(c *gin.Context) {
	// Extract user ID from the context set by AuthMiddleware
	userIdString, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, _ := (userIdString).(int64)

	// Fetch user from the database
	user, err := ac.DbService.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	})
}

func (ac ApiController) CheckRole(c *gin.Context, role string) bool {
	userIdString, exists := c.Get("user_id")
	if !exists {
		return false
	}
	userID, _ := (userIdString).(int64)

	// Fetch user from the database
	user, err := ac.DbService.GetUserByID(c, userID)
	if err != nil {
		return false
	}

	if role != user.Role {
		return false
	}
	return true
}
