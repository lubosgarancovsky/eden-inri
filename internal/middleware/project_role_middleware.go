package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
)

func ProjectRoleMiddleware(projectUserService *service.ProjectUserService, allowed ...model.ProjectRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		projectID := helpers.ExtractID(c, "projectId")
		userID := helpers.GetUserContext(c).ID

		role, err := projectUserService.GetUserRole(projectID, userID)
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
