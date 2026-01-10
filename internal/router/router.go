package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/config"
	"github.com/lubosgarancovsky/eden-inri/internal/handlers"
	"github.com/lubosgarancovsky/eden-inri/internal/middleware"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/eden-inri/internal/services"
	"github.com/lubosgarancovsky/go-kit/rsql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {
	// TODO: Remove after migrating to GlobalConfig ( temp )
	cfg := config.GlobalConfig

	parser := rsql.New()
	v1 := r.Group("/v1/inri", middleware.ErrorMiddleware(), middleware.AppStateValidationMiddleware(db))
	protected := v1.Group("", middleware.AuthMiddleware())

	v1.GET("/internal/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO: Create better health service and make app unhealthy if db == nil
	v1.GET("/internal/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	emailService := services.NewEmailService(cfg)

	// Attachments
	attachmentRepo := repositories.NewAttachmentRepository(db)
	attachmentService := services.NewAttachmentService(cfg, attachmentRepo)
	attachmentHandler := handlers.NewAttachmentHandler(parser, attachmentService)

	// Contact persons
	cpRepo := repositories.NewContactPersonRepository(db)
	cpService := services.NewContactPersonService(cpRepo)
	cpHandler := handlers.NewContactPersonHandler(parser, cpService)

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
	clientRepo := repositories.NewClientRepository(db)
	clientService := services.NewClientService(clientRepo)
	clientHandler := handlers.NewClientHandler(parser, clientService)
	{
		clients.GET("", clientHandler.FindAll)
		clients.GET("/:clientId", clientHandler.FindByID)
		clients.POST("", clientHandler.Create)
		clients.PUT("/:clientId", clientHandler.Update)
		clients.DELETE("/:clientId", clientHandler.Delete)
	}

	// Invoices
	invRepo := repositories.NewInvoiceRepository(db)
	invService := services.NewInvoiceService(invRepo, attachmentService)
	invHandler := handlers.NewInvoiceHandler(parser, invService)

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
	projectUserRepo := repositories.NewProjectUserRepository(db)
	projectUserService := services.NewProjectUserService(projectUserRepo)
	projectUserHandler := handlers.NewProjectUserHandler(parser, projectUserService)

	projects := protected.Group("/projects")
	{
		projects.GET("/:projectId/members", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), projectUserHandler.FindAll)
		projects.PUT("/:projectId/members/:memberId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), projectUserHandler.Update)
		projects.DELETE("/:projectId/members/:memberId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), projectUserHandler.Delete)
	}

	// Projects
	projectRepo := repositories.NewProjectRepository(db)
	projectService := services.NewProjectService(projectRepo, attachmentService, projectUserService)
	projectHandler := handlers.NewProjectHandler(parser, projectService)

	{
		projects.GET("", projectHandler.FindAll)
		projects.GET("/:projectId", projectHandler.FindByID)
		projects.POST("", projectHandler.Create)
		projects.PUT("/:projectId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), projectHandler.Update)
		projects.DELETE("/:projectId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner), projectHandler.Delete)
		projects.PUT("/:projectId/favourite", projectHandler.Favourite)

		// Project attachments
		projects.POST("/:projectId/attachments", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), projectHandler.UploadAttachments)
		projects.GET("/:projectId/attachments", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), projectHandler.ListAttachments)
		projects.GET("/:projectId/attachments/:attachmentId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), projectHandler.DownloadAttachment)
		projects.DELETE("/:projectId/attachments/:attachmentId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), projectHandler.DeleteAttachment)
	}

	// Labels
	labelRepo := repositories.NewLabelRepository(db)
	labelService := services.NewLabelService(labelRepo, projectUserService)
	labelHandler := handlers.NewLabelHandler(parser, labelService)

	labels := projects.Group("/:projectId/labels")
	{
		labels.GET("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), labelHandler.FindAll)
		labels.GET("/:labelId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), labelHandler.FindByID)
		labels.POST("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), labelHandler.Insert)
		labels.PUT("/:labelId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), labelHandler.Update)
		labels.DELETE("/:labelId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), labelHandler.Delete)
	}

	// Kanban boards
	kanbanRepo := repositories.NewKanbanBoardRepository(db)
	kanbanService := services.NewKanbanBoardService(kanbanRepo, projectUserService)
	kanbanHandler := handlers.NewKanbanBoardHandler(kanbanService, parser)

	kanban := projects.Group("/:projectId/kanban")
	{
		kanban.GET("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), kanbanHandler.FindAll)
		kanban.POST("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), kanbanHandler.Insert)
		kanban.PUT("/:kanbanId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), kanbanHandler.Update)
		kanban.GET("/:kanbanId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), kanbanHandler.FindByID)
		kanban.DELETE("/:kanbanId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), kanbanHandler.Delete)
	}

	// Project Documents
	projectDocRepo := repositories.NewProjectDocumentRepository(db)
	projectDocService := services.NewProjectDocumentService(projectDocRepo, projectUserService)
	projectDocHandler := handlers.NewProjectDocumentHandler(parser, projectDocService)

	projectDocuments := projects.Group(":projectId/documents")
	{
		projectDocuments.GET("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), projectDocHandler.FindAll)
		projectDocuments.GET("/:documentId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), projectDocHandler.FindByID)
		projectDocuments.POST("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), projectDocHandler.Create)
		projectDocuments.PUT("/:documentId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), projectDocHandler.Update)
		projectDocuments.DELETE("/:documentId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), projectDocHandler.Delete)
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
	kanbanColumnsRepo := repositories.NewKanbanColumnRepository(db)
	kanbanColumnsService := services.NewKanbanColumnService(kanbanColumnsRepo, projectUserService, kanbanService)
	kanbanColumnsHandler := handlers.NewKanbanColumnHandler(kanbanColumnsService)

	column := kanban.Group("/:kanbanId/columns")
	{
		column.GET("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), kanbanColumnsHandler.FindAll)
		column.GET("/:columnId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), kanbanColumnsHandler.FindByID)
		column.POST("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), kanbanColumnsHandler.Insert)
		column.PUT("/:columnId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), kanbanColumnsHandler.Update)
		column.DELETE("/:columnId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), kanbanColumnsHandler.Delete)
	}

	// Stories
	storyRepo := repositories.NewStoryRepository(db)
	storyService := services.NewStoryService(storyRepo, projectService, projectUserService, kanbanService, kanbanColumnsService, attachmentService)
	storyHandler := handlers.NewStoryHandler(parser, storyService)

	stories := projects.Group("/:projectId/stories")

	{
		protected.GET("/stories", storyHandler.FindAllAssigned)
		stories.GET("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), storyHandler.FindAll)
		stories.POST("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyHandler.Insert)
		stories.GET("/slug/:slug", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), storyHandler.FindBySlug)
		stories.GET("/:storyId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), storyHandler.FindByID)
		stories.PUT("/:storyId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyHandler.Update)
		stories.DELETE("/:storyId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyHandler.Delete)
		stories.GET("/:storyId/assignee", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), storyHandler.GetAssignee)
		stories.PUT("/:storyId/assignee", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyHandler.ChangeAssignee)

		// Story attachments
		stories.POST("/:storyId/attachments", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyHandler.UploadAttachments)
		stories.GET("/:storyId/attachments", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), storyHandler.ListAttachments)
		stories.GET("/:storyId/attachments/:attachmentId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), storyHandler.DownloadAttachment)
		stories.DELETE("/:storyId/attachments/:attachmentId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyHandler.DeleteAttachment)
	}

	// Story activity
	storyActivityRepo := repositories.NewStoryActivityRepository(db)
	storyActivityService := services.NewStoryActivityService(storyActivityRepo, projectUserService)
	storyActivityHandler := handlers.NewStoryActivityHandler(parser, storyActivityService)

	{
		stories.GET("/:storyId/activities", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), storyActivityHandler.ListActivities)
		stories.POST("/:storyId/activities", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyActivityHandler.InsertActivity)
		stories.PUT("/:storyId/activities/:activityId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyActivityHandler.UpdateActivity)
	}

	// Story labels
	storyLabelRepo := repositories.NewStoryLabelRepository(db)
	storyLabelService := services.NewStoryLabelService(storyLabelRepo, projectUserService, kanbanService)
	storyLabelHandler := handlers.NewStoryLabelHandler(storyLabelService)

	storyLabels := stories.Group("/:storyId/labels")
	{
		storyLabels.GET("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer, models.Guest), storyLabelHandler.ListLabels)
		storyLabels.POST("", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyLabelHandler.AssignLabel)
		storyLabels.DELETE("/:labelId", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin, models.Developer), storyLabelHandler.UnassignLabel)

	}

	// Project invitation
	invitationRepo := repositories.NewProjectInvitationRepository(db)
	invitationService := services.NewProjectInvitationService(cfg, invitationRepo, projectService, projectUserService, emailService)
	invitationHandler := handlers.NewProjectInvitationsHandler(invitationService)
	{
		projects.POST("/:projectId/invite", middleware.ProjectRoleMiddleware(projectUserService, models.Owner, models.Admin), invitationHandler.Create)
		projects.POST("/accept-invitation", invitationHandler.Accept)
	}

	// TODO: Add time log API

	return r
}
