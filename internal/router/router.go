package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/handler"
	"github.com/lubosgarancovsky/eden-inri/internal/middleware"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/go-kit/rsql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, cfg *config.Config, db *gorm.DB) *gin.Engine {
	parser := rsql.New()
	v1 := r.Group("/v1/inri", middleware.ErrorMiddleware())
	protected := v1.Group("", middleware.AuthMiddleware())

	v1.GET("/internal/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.GET("/internal/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Attachments
	attachmentRepo := repository.NewAttachmentRepository(db)
	attachmentService := service.NewAttachmentService(cfg, attachmentRepo)
	attachmentHandler := handler.NewAttachmentHandler(parser, attachmentService)

	// Clients
	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handler.NewClientHandler(parser, clientService)

	clients := protected.Group("/clients")
	{
		clients.GET("", clientHandler.FindAll)
		clients.GET("/:clientId", clientHandler.FindByID)
		clients.POST("", clientHandler.Create)
		clients.PUT("/:clientId", clientHandler.Update)
		clients.DELETE("/:clientId", clientHandler.Delete)
	}

	// Contact persons
	cpRepo := repository.NewContactPersonRepository(db)
	cpService := service.NewContactPersonService(cpRepo)
	cpHandler := handler.NewContactPersonHandler(parser, cpService)

	contactPersons := protected.Group("/contact-persons")
	{
		contactPersons.GET("", cpHandler.FindAll)
		contactPersons.GET("/:contactPersonId", cpHandler.FindByID)
		contactPersons.POST("", cpHandler.Create)
		contactPersons.PUT("/:contactPersonId", cpHandler.Update)
		contactPersons.DELETE("/:contactPersonId", cpHandler.Delete)
	}

	// Invoices
	invRepo := repository.NewInvoiceRepository(db)
	invService := service.NewInvoiceService(invRepo, attachmentService)
	invHandler := handler.NewInvoiceHandler(parser, invService)

	invoices := protected.Group("/invoices")
	{
		invoices.GET("", invHandler.FindAll)
		invoices.GET("/:invoiceId", invHandler.FindByID)
		invoices.POST("", invHandler.Create)
		invoices.PUT("/:invoiceId", invHandler.Update)
		invoices.DELETE("/:invoiceId", invHandler.Delete)
		invoices.POST("/:invoiceId/attachments", invHandler.UploadAttachments)
		invoices.GET("/:invoiceId/attachments", invHandler.ListAttachments)
	}

	// Project users
	projectUserRepo := repository.NewProjectUserRepository(db)
	projectUserService := service.NewProjectUserService(projectUserRepo)
	projectUserHandler := handler.NewProjectUserHandler(parser, projectUserService)

	projects := protected.Group("/projects")
	{
		projects.GET("/:projectId/members", projectUserHandler.FindAll)
		projects.DELETE("/:projectId/members/:memberId", projectUserHandler.Delete)
	}

	// Projects
	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo, attachmentService, projectUserService)
	projectHandler := handler.NewProjectHandler(parser, projectService)

	{
		projects.GET("", projectHandler.FindAll)
		projects.GET("/:projectId", projectHandler.FindByID)
		projects.POST("", projectHandler.Create)
		projects.PUT("/:projectId", projectHandler.Update)
		projects.DELETE("/:projectId", projectHandler.Delete)
		projects.POST("/:projectId/attachments", projectHandler.UploadAttachments)
		projects.GET("/:projectId/attachments", projectHandler.ListAttachments)
	}

	// Kanban boards
	kanbanRepo := repository.NewKanbanBoardRepository(db)
	kanbanService := service.NewKanbanBoardService(kanbanRepo, projectUserService)
	kanbanHandler := handler.NewKanbanBoardHandler(kanbanService)

	{
		projects.GET("/:projectId/boards", kanbanHandler.FindAll)
		projects.POST("/:projectId/boards", projectHandler.Create)
		projects.GET("/:projectId/boards/:boardId", kanbanHandler.FindByID)
		projects.DELETE("/:projectId/boards/:boardId", kanbanHandler.Delete)
	}

	// Project Documents
	projectDocRepo := repository.NewProjectDocumentRepository(db)
	projectDocService := service.NewProjectDocumentService(projectDocRepo, projectRepo)
	projectDocHandler := handler.NewProjectDocumentHandler(parser, projectDocService)

	projectDocuments := projects.Group(":projectId/documents")
	{
		projectDocuments.GET("", projectDocHandler.FindAll)
		projectDocuments.GET("/:documentId", projectDocHandler.FindByID)
		projectDocuments.POST("", projectDocHandler.Create)
		projectDocuments.PUT("/:documentId", projectDocHandler.Update)
		projectDocuments.DELETE("/:documentId", projectDocHandler.Delete)
	}

	// Attachments
	attachments := protected.Group("/attachments")
	{
		attachments.GET("", attachmentHandler.FindAll)
		attachments.GET("/:attachmentId", attachmentHandler.FindByID)
		attachments.DELETE("/:attachmentId", attachmentHandler.Delete)
		attachments.GET("/:attachmentId/download", attachmentHandler.Download)
	}

	return r
}
