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

	// Labels
	labelRepo := repository.NewLabelRepository(db)
	labelService := service.NewLabelService(labelRepo, projectUserService)
	labelHandelr := handler.NewLabelHandler(labelService)

	labels := protected.Group("/:projectId/labels")
	{
		labels.GET("", labelHandelr.FindAll)
		labels.GET("/:labelId", labelHandelr.FindByID)
		labels.POST("", labelHandelr.Insert)
		labels.PUT("/:labelId", labelHandelr.Update)
		labels.DELETE("/:labelId", labelHandelr.Delete)
	}

	// Kanban boards
	kanbanRepo := repository.NewKanbanBoardRepository(db)
	kanbanService := service.NewKanbanBoardService(kanbanRepo, projectUserService)
	kanbanHandler := handler.NewKanbanBoardHandler(kanbanService)

	{
		projects.GET("/:projectId/kanban", kanbanHandler.FindAll)
		projects.POST("/:projectId/kanban", projectHandler.Create)
		projects.GET("/:projectId/kanban/:kanbanId", kanbanHandler.FindByID)
		projects.DELETE("/:projectId/kanban/:kanbanId", kanbanHandler.Delete)
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

	// Kanban columns
	kanbanColumnsRepo := repository.NewKanbanColumnRepository(db)
	kanbanColumnsService := service.NewKanbanColumnService(kanbanColumnsRepo, projectUserService, kanbanService)
	kanbanColumnsHandler := handler.NewKanbanColumnHandler(kanbanColumnsService)

	kanban := r.Group("/kanban")
	{
		kanban.GET("/:kanbanId/columns", kanbanColumnsHandler.FindAll)
		kanban.GET("/:kanbanId/columns/:columnId", kanbanColumnsHandler.FindByID)
		kanban.POST("/:kanbanId/columns", kanbanColumnsHandler.Insert)
		kanban.PUT("/:kanbanId/columns/:columnId", kanbanColumnsHandler.Update)
		kanban.DELETE("/:kanbanId/columns/:columnId", kanbanColumnsHandler.Delete)
	}

	// Stories
	storyRepo := repository.NewStoryRepository(db)
	storyService := service.NewStoryService(storyRepo, projectService, projectUserService, kanbanService)
	storyHandler := handler.NewStoryHandler(parser, storyService)

	stories := kanban.Group("/:kanbanId/stories")

	{
		kanban.GET("/:kanbanId/columns/:columnId/stories", storyHandler.FindAll)
		kanban.GET("/:kanbanId/stories/:storyId", storyHandler.FindByID)
		stories.POST("", storyHandler.Insert)
		stories.PUT("/:storyId", storyHandler.Update)
		stories.DELETE("/:storyId", storyHandler.Delete)
	}

	// Story activity
	storyActivityRepo := repository.NewStoryActivityRepository(db)
	storyActivityService := service.NewStoryActivityService(storyActivityRepo, projectUserService, kanbanService)
	storyActivityHandler := handler.NewStoryActivityHandler(parser, storyActivityService)

	{
		stories.GET("/:storyId/activities", storyActivityHandler.ListActivities)
		stories.POST("/:storyId/activities", storyActivityHandler.InsertActivity)

	}

	// Story labels
	storyLabelRepo := repository.NewStoryLabelRepository(db)
	storyLabelService := service.NewStoryLabelService(storyLabelRepo, projectUserService, kanbanService)
	storyLabelHandler := handler.NewStoryLabelHandler(storyLabelService)

	storyLabels := kanban.Group("/:kanbanId/stories/:storyId/labels")
	{
		storyLabels.GET("", storyLabelHandler.ListLabels)
		storyLabels.GET("/:labelId", storyLabelHandler.AssignLabel)
		storyLabels.POST("/:labelId", storyLabelHandler.UnassignLabel)

	}

	// TODO: Add time log API

	return r
}
