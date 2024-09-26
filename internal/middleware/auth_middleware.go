package middleware

import (
	"ayuphone_api/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		utils.ErrorResponse(c, http.StatusUnauthorized, " Invalid token")
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, " Invalid token")
		c.Abort()
		return
	}

	c.Set("user_id", claims.ID)
	c.Next()
}
