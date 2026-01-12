package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/go-kit"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check for errors in context
		if len(c.Errors) > 0 {
			debug.PrintStack()
			err := c.Errors.Last().Err
			handle.Error(c, err)
			return
		}

		// Handle errors that might be set in context but not in c.Errors
		if c.Writer.Status() >= 400 {
			debug.PrintStack()

			switch c.Writer.Status() {
			case http.StatusBadRequest:
				handle.Error(c, go_kit.ErrBadRequest)
			case http.StatusNotFound:
				handle.Error(c, go_kit.ErrNotFound)
			case http.StatusUnauthorized:
				handle.Error(c, go_kit.ErrUnauthorized)
			case http.StatusForbidden:
				handle.Error(c, go_kit.ErrForbidden)
			case http.StatusTooManyRequests:
				handle.Error(c, go_kit.ErrTooManyRequests)
			default:
				handle.Error(c, go_kit.ErrInternalServer)
			}
		}

	}
}
