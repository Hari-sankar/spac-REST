package middleware

import (
	"spac-REST/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create context with request info
		ctx := logger.WithRequest(c.Request.Context(), c.Request)

		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// Use your existing logger with context
		log.With(ctx,
			"method", method,
			"path", path,
			"status", statusCode,
			"latency", latency,
		).Info("Request processed")
	}
}
