package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/handler"
	"github.com/lubosgarancovsky/eden-inri/internal/middleware"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/go-kit/rsql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, cfg *config.Config, db *gorm.DB) *gin.Engine {
	parser := rsql.New()
	v1 := r.Group("/v1/inri", middleware.ErrorMiddleware(), middleware.AppStateValidationMiddleware(db))
	protected := v1.Group("", middleware.AuthMiddleware())

	v1.GET("/internal/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO: Create better health service and make app unhealthy if db == nil
	v1.GET("/internal/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	emailService := service.NewEmailService(cfg)

	// Attachments
	attachmentRepo := repository.NewAttachmentRepository(db)
	attachmentService := service.NewAttachmentService(cfg, attachmentRepo)
	attachmentHandler := handler.NewAttachmentHandler(parser, attachmentService)

	// Contact persons
	cpRepo := repository.NewContactPersonRepository(db)
	cpService := service.NewContactPersonService(cpRepo)
	cpHandler := handler.NewContactPersonHandler(parser, cpService)

	clients := protected.Group("/clients")
	contactPersons := clients.Group("/:clientId/contact-persons")
	{
		contactPersons.GET("", cpHandler.FindAll)
		contactPersons.GET("/:contactPersonId", cpHandler.FindByID)
		contactPersons.POST("", cpHandler.Create)
		contactPersons.PUT("/:contactPersonId", cpHandler.Update)
		contactPersons.DELETE("/:contactPersonId", cpHandler.Delete)
	}

	// Clients
	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepo)
	clientHandler := handler.NewClientHandler(parser, clientService)
	{
		clients.GET("", clientHandler.FindAll)
		clients.GET("/:clientId", clientHandler.FindByID)
		clients.POST("", clientHandler.Create)
		clients.PUT("/:clientId", clientHandler.Update)
		clients.DELETE("/:clientId", clientHandler.Delete)
	}

	// Invoices
	invRepo := repository.NewInvoiceRepository(db)
	invService := service.NewInvoiceService(invRepo, attachmentService)
	invHandler := handler.NewInvoiceHandler(parser, invService)

	invoices := protected.Group("/invoices")
	{
		invoices.GET("", invHandler.FindAll)
		invoices.GET("/revenue", invHandler.TotalRevenue)
		invoices.GET("/revenue/graph", invHandler.RevenueGraph)
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
		projects.GET("/:projectId/members", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), projectUserHandler.FindAll)
		projects.PUT("/:projectId/members/:memberId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), projectUserHandler.Update)
		projects.DELETE("/:projectId/members/:memberId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), projectUserHandler.Delete)
	}

	// Projects
	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo, attachmentService, projectUserService)
	projectHandler := handler.NewProjectHandler(parser, projectService)

	{
		projects.GET("", projectHandler.FindAll)
		projects.GET("/:projectId", projectHandler.FindByID)
		projects.POST("", projectHandler.Create)
		projects.PUT("/:projectId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), projectHandler.Update)
		projects.DELETE("/:projectId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner), projectHandler.Delete)
		projects.POST("/:projectId/favourite", projectHandler.Favourite)

		// Project attachments
		projects.POST("/:projectId/attachments", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), projectHandler.UploadAttachments)
		projects.GET("/:projectId/attachments", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), projectHandler.ListAttachments)
		projects.GET("/:projectId/attachments/:attachmentId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), projectHandler.DownloadAttachment)
		projects.DELETE("/:projectId/attachments/:attachmentId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), projectHandler.DeleteAttachment)
	}

	// Labels
	labelRepo := repository.NewLabelRepository(db)
	labelService := service.NewLabelService(labelRepo, projectUserService)
	labelHandler := handler.NewLabelHandler(parser, labelService)

	labels := projects.Group("/:projectId/labels")
	{
		labels.GET("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), labelHandler.FindAll)
		labels.GET("/:labelId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), labelHandler.FindByID)
		labels.POST("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), labelHandler.Insert)
		labels.PUT("/:labelId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), labelHandler.Update)
		labels.DELETE("/:labelId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), labelHandler.Delete)
	}

	// Kanban boards
	kanbanRepo := repository.NewKanbanBoardRepository(db)
	kanbanService := service.NewKanbanBoardService(kanbanRepo, projectUserService)
	kanbanHandler := handler.NewKanbanBoardHandler(kanbanService, parser)

	kanban := projects.Group("/:projectId/kanban")
	{
		kanban.GET("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), kanbanHandler.FindAll)
		kanban.POST("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), kanbanHandler.Insert)
		kanban.PUT("/:kanbanId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), kanbanHandler.Update)
		kanban.GET("/:kanbanId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), kanbanHandler.FindByID)
		kanban.DELETE("/:kanbanId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), kanbanHandler.Delete)
	}

	// Project Documents
	projectDocRepo := repository.NewProjectDocumentRepository(db)
	projectDocService := service.NewProjectDocumentService(projectDocRepo, projectUserService)
	projectDocHandler := handler.NewProjectDocumentHandler(parser, projectDocService)

	projectDocuments := projects.Group(":projectId/documents")
	{
		projectDocuments.GET("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), projectDocHandler.FindAll)
		projectDocuments.GET("/:documentId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), projectDocHandler.FindByID)
		projectDocuments.POST("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), projectDocHandler.Create)
		projectDocuments.PUT("/:documentId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), projectDocHandler.Update)
		projectDocuments.DELETE("/:documentId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), projectDocHandler.Delete)
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

	column := kanban.Group("/:kanbanId/columns")
	{
		column.GET("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), kanbanColumnsHandler.FindAll)
		column.GET("/:columnId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), kanbanColumnsHandler.FindByID)
		column.POST("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), kanbanColumnsHandler.Insert)
		column.PUT("/:columnId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), kanbanColumnsHandler.Update)
		column.DELETE("/:columnId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), kanbanColumnsHandler.Delete)
	}

	// Stories
	storyRepo := repository.NewStoryRepository(db)
	storyService := service.NewStoryService(storyRepo, projectService, projectUserService, kanbanService, attachmentService)
	storyHandler := handler.NewStoryHandler(parser, storyService)

	stories := projects.Group("/:projectId/stories")

	{
		stories.GET("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), storyHandler.FindAll)
		stories.POST("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyHandler.Insert)
		stories.GET("/slug/:slug", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), storyHandler.FindBySlug)
		stories.GET("/:storyId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), storyHandler.FindByID)
		stories.PUT("/:storyId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyHandler.Update)
		stories.DELETE("/:storyId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyHandler.Delete)
		stories.GET("/:storyId/assignee", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), storyHandler.GetAssignee)
		stories.PUT("/:storyId/assignee", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyHandler.ChangeAssignee)

		// Story attachments
		stories.POST("/:storyId/attachments", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyHandler.UploadAttachments)
		stories.GET("/:storyId/attachments", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), storyHandler.ListAttachments)
		stories.GET("/:storyId/attachments/:attachmentId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), storyHandler.DownloadAttachment)
		stories.DELETE("/:storyId/attachments/:attachmentId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyHandler.DeleteAttachment)
	}

	// Story activity
	storyActivityRepo := repository.NewStoryActivityRepository(db)
	storyActivityService := service.NewStoryActivityService(storyActivityRepo, projectUserService)
	storyActivityHandler := handler.NewStoryActivityHandler(parser, storyActivityService)

	{
		stories.GET("/:storyId/activities", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), storyActivityHandler.ListActivities)
		stories.POST("/:storyId/activities", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyActivityHandler.InsertActivity)
		stories.PUT("/:storyId/activities/:activityId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyActivityHandler.UpdateActivity)
	}

	// Story labels
	storyLabelRepo := repository.NewStoryLabelRepository(db)
	storyLabelService := service.NewStoryLabelService(storyLabelRepo, projectUserService, kanbanService)
	storyLabelHandler := handler.NewStoryLabelHandler(storyLabelService)

	storyLabels := stories.Group("/:storyId/labels")
	{
		storyLabels.GET("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer, model.Guest), storyLabelHandler.ListLabels)
		storyLabels.POST("", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyLabelHandler.AssignLabel)
		storyLabels.DELETE("/:labelId", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin, model.Developer), storyLabelHandler.UnassignLabel)

	}

	// Project invitation
	invitationRepo := repository.NewProjectInvitationRepository(db)
	invitationService := service.NewProjectInvitationService(cfg, invitationRepo, projectService, projectUserService, emailService)
	invitationHandler := handler.NewProjectInvitationsHandler(invitationService)
	{
		projects.POST("/:projectId/invite", middleware.ProjectRoleMiddleware(projectUserService, model.Owner, model.Admin), invitationHandler.Create)
		projects.POST("/accept-invitation", invitationHandler.Accept)
	}

	// TODO: Add time log API

	return r
}
