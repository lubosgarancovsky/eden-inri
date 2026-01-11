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
		clients := protected.Group("/clients")
		{
			clients.POST("", c.ClientHandler.Create)
			clients.PUT("/:clientId", c.ClientHandler.Update)
			clients.DELETE("/:clientId", c.ClientHandler.Delete)
			clients.GET("/:clientId", c.ClientHandler.FindByID)
			clients.GET("", c.ClientHandler.List)
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
			projects.POST("/:projectId/favourite", c.ProjectHandler.Favourite)

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

			boards := projects.Group("/:projectId/boards")
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
		}
	}

	return router
}
