package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/go-kit"
)

type AuthContext struct {
	UserID        uuid.UUID `header:"X-User-ID" binding:"required,uuid4"`
	UserRole      string    `header:"X-User-Role" binding:"required,oneof=admin user"`
	Authorization *string   `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := &AuthContext{}

		err := c.ShouldBindHeader(headers)
		if err != nil {
			handle.Error(c, go_kit.ErrUnauthorized.WithMessage("Authorization token is expired, missing or malformed"))
			c.Abort()
			return
		}

		c.Set("user", headers)

		c.Next()
	}
}
