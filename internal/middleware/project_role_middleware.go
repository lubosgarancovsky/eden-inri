package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/services"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
)

func ProjectRoleMiddleware(projectUserService *services.ProjectUserService, allowed ...models.ProjectRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID := helpers.ExtractID(c, "projectId")
		userID := helpers.GetUserContext(c).ID

		role, err := projectUserService.GetUserRole(c.Request.Context(), projectID, userID)
		if err != nil {
			c.Error(err)
			c.Abort()
		}

		if err = projectUserService.RequireRole(role, allowed...); err != nil {
			c.Error(err)
			c.Abort()
		}

		c.Next()
	}
}
