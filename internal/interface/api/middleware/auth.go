package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort() 
			return
		}
		
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// userID := "extracted-token-user-id"
		c.Set("userID", token)

		c.Next()
	}
}