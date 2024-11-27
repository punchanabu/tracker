// interface/api/middleware/logger.go
package middleware

import (
	"time"

	"github.com/punchanabu/portfolio-tracker/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get our configured logger
		logger := log.GetLogger()

		// Record starting time
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate request duration
		duration := time.Since(start)

		// Log request details with structured fields
		logger.Info("request completed",
			// Request information
			zap.String("path", path),
			zap.String("query", query),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),

			// Response information
			zap.Int("status", c.Writer.Status()),
			zap.Int("size", c.Writer.Size()),

			// Timing
			zap.Duration("duration", duration),

			// Error details if any
			zap.String("error", c.Errors.String()),
		)
	}
}
