package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/middleware"
	"github.com/lubosgarancovsky/eden-inri/internal/app"
)

func NewServerRoute(c *app.Container) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.ErrorHandlerMiddleware())

	v1 := router.Group("/v1/inri")
	protected := v1.Group("", middleware.AuthMiddleware())

	{
		attachments := protected.Group("/attachments")
		{
			attachments.GET("", c.AttachmentHandler.List)
			attachments.GET("/:attachmentId", c.AttachmentHandler.FindByID)
			attachments.PUT("/:attachmentId", c.AttachmentHandler.Update)
			attachments.DELETE("/:attachmentId", c.AttachmentHandler.Delete)
			attachments.POST("", c.AttachmentHandler.Upload)
			attachments.GET("/:attachmentId/download", c.AttachmentHandler.Download)
		}

		clients := protected.Group("/clients")
		{
			clients.POST("", c.ClientHandler.Create)
			clients.PUT("/:clientId", c.ClientHandler.Update)
			clients.DELETE("/:clientId", c.ClientHandler.Delete)
			clients.GET("/:clientId", c.ClientHandler.FindByID)
			clients.GET("", c.ClientHandler.List)

			contactPersons := clients.Group("/:clientId/contact-persons")
			{
				contactPersons.POST("", c.ContactPersonHandler.Create)
				contactPersons.PUT("/:contactPersonId", c.ContactPersonHandler.Update)
				contactPersons.DELETE("/:contactPersonId", c.ContactPersonHandler.Delete)
				contactPersons.GET("/:contactPersonId", c.ContactPersonHandler.FindByID)
				contactPersons.GET("", c.ContactPersonHandler.List)
			}
		}

		invoices := protected.Group("/invoices")
		{
			invoices.POST("", c.InvoiceHandler.Create)
			invoices.PUT("/:invoiceId", c.InvoiceHandler.Update)
			invoices.DELETE("/:invoiceId", c.InvoiceHandler.Delete)
			invoices.GET("/:invoiceId", c.InvoiceHandler.FindByID)
			invoices.GET("", c.InvoiceHandler.List)
		}

		projects := protected.Group("/projects")
		{
			projects.POST("", c.ProjectHandler.Create)
			projects.PUT("/:projectId", c.ProjectHandler.Update)
			projects.DELETE("/:projectId", c.ProjectHandler.Delete)
			projects.GET("/:projectId", c.ProjectHandler.FindByID)
			projects.GET("", c.ProjectHandler.List)
			projects.PUT("/:projectId/favourite", c.ProjectHandler.Favourite)

			projects.POST("/:projectId/invite", c.ProjectInvitationHandler.Invite)
			projects.POST("/accept-invitation", c.ProjectInvitationHandler.Accept)

			projectAttachments := projects.Group("/:projectId/attachments")
			{
				projectAttachments.GET("/:attachmentId/download", c.ProjectAttachmentHandler.Download)
				projectAttachments.PUT("/:attachmentId", c.ProjectAttachmentHandler.Update)
				projectAttachments.DELETE("/:attachmentId", c.ProjectAttachmentHandler.Delete)
				projectAttachments.GET("/:attachmentId", c.ProjectAttachmentHandler.FindByID)
				projectAttachments.GET("", c.ProjectAttachmentHandler.List)
			}

			projectUsers := projects.Group("/:projectId/members")
			{
				projectUsers.GET("", c.ProjectUserHandler.List)
				projectUsers.PUT("/:memberId", c.ProjectUserHandler.ChangeRole)
				projectUsers.DELETE("/:memberId", c.ProjectUserHandler.Delete)
			}

			labels := projects.Group("/:projectId/labels")
			{
				labels.POST("", c.ProjectLabelHandler.Create)
				labels.PUT("/:labelId", c.ProjectLabelHandler.Update)
				labels.DELETE("/:labelId", c.ProjectLabelHandler.Delete)
				labels.GET("/:labelId", c.ProjectLabelHandler.FindByID)
				labels.GET("", c.ProjectLabelHandler.List)
			}

			documents := projects.Group("/:projectId/documents")
			{
				documents.POST("", c.ProjectDocumentHandler.Create)
				documents.PUT("/:documentId", c.ProjectDocumentHandler.Update)
				documents.DELETE("/:documentId", c.ProjectDocumentHandler.Delete)
				documents.GET("/:documentId", c.ProjectDocumentHandler.FindByID)
				documents.GET("", c.ProjectDocumentHandler.List)
			}

			boards := projects.Group("/:projectId/kanban")
			{
				boards.POST("", c.KanbanBoardHandler.Create)
				boards.PUT("/:boardId", c.KanbanBoardHandler.Update)
				boards.DELETE("/:boardId", c.KanbanBoardHandler.Delete)
				boards.GET("/:boardId", c.KanbanBoardHandler.FindByID)
				boards.GET("", c.KanbanBoardHandler.List)

				columns := boards.Group("/:boardId/columns")
				{
					columns.POST("", c.KanbanColumnHandler.Create)
					columns.PUT("/:columnId", c.KanbanColumnHandler.Update)
					columns.DELETE("/:columnId", c.KanbanColumnHandler.Delete)
					columns.GET("", c.KanbanColumnHandler.List)
				}
			}

			stories := projects.Group("/:projectId/stories")
			{
				stories.POST("", c.StoryHandler.Create)
				stories.PUT("/:storyId", c.StoryHandler.Update)
				stories.DELETE("/:storyId", c.StoryHandler.Delete)
				stories.GET("/:storyId", c.StoryHandler.FindByID)
				stories.GET("", c.StoryHandler.List)
				stories.PUT("/:storyId/assignee", c.StoryHandler.ChangeAssignee)

				activities := stories.Group("/:storyId/activities")
				{
					activities.POST("", c.StoryActivityHandler.Create)
					activities.PUT("/:activityId", c.StoryActivityHandler.Update)
					activities.DELETE("/:activityId", c.StoryActivityHandler.Delete)
					activities.GET("", c.StoryActivityHandler.List)
				}

				storyAttachments := stories.Group("/:storyId/attachments")
				{
					storyAttachments.GET("", c.StoryAttachmentHandler.List)
					storyAttachments.DELETE("/:attachmentId", c.StoryAttachmentHandler.Delete)
					storyAttachments.GET("/:attachmentId/download", c.StoryAttachmentHandler.Download)
				}
			}
		}

		protected.GET("/stories", c.StoryHandler.ListAssigned)
	}

	return router
}
