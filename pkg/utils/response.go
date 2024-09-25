package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONResponse formats the response with a success flag, message, and data.
func JSONResponse(c *gin.Context, success bool, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse formats the response for error scenarios.
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": message,
	})
}
