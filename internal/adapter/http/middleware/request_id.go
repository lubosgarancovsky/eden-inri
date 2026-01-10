package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	// RequestIDHeader is the header key for request ID
	RequestIDHeader = "X-Request-ID"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(RequestIDHeader)
		if requestID != "" {
			requestID = uuid.New().String()
		}

		c.Writer.Header().Set(RequestIDHeader, requestID)
		c.Set(RequestIDHeader, requestID)

		c.Next()
	}
}
