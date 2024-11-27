// interface/api/middleware/recovery.go
package middleware

import (
	"net/http"
	"runtime/debug"
	"github.com/punchanabu/portfolio-tracker/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Get stack trace
				stack := string(debug.Stack())

				// Log the error with stack trace
				log.GetLogger().Error("panic recovered",
					zap.Any("error", err),
					zap.String("stack", stack),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
				)

				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "An internal error occurred",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
