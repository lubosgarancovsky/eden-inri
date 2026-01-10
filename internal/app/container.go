package app

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handler"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres"
	"github.com/lubosgarancovsky/eden-inri/internal/app/client"
	app_invoice "github.com/lubosgarancovsky/eden-inri/internal/app/invoice"
	app_label "github.com/lubosgarancovsky/eden-inri/internal/app/label"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project"
	project_document_uc "github.com/lubosgarancovsky/eden-inri/internal/app/project_document"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type Container struct {
	db *gorm.DB

	parser *go_kit.Parser

	// -- Repositories --
	clientRepository          *postgres.ClientRepository
	projectRepository         *postgres.ProjectRepository
	projectUserRepository     *postgres.ProjectUserRepository
	projectDocumentRepository *postgres.ProjectDocumentRepository
	labelRepository           *postgres.LabelRepository
	invoiceRepository         *postgres.InvoiceRepository
	txManager                 *postgres.TransactionManager

	// -- Services --
	listClientsService    *client.ListClientsService
	findClientByIDService *client.FindClientByIDService
	createClientService   *client.CreateClientService
	updateClientService   *client.UpdateClientService
	deleteClientService   *client.DeleteClientService

	// Project services
	listProjectsService    *project.ListProjectsService
	findProjectByIDService *project.FindProjectByIDService
	createProjectService   *project.CreateProjectService
	updateProjectService   *project.UpdateProjectService
	deleteProjectService   *project.DeleteProjectService

	// ProjectDocument services
	listProjectDocumentsService    *project_document_uc.ListProjectDocumentsService
	findProjectDocumentByIDService *project_document_uc.FindProjectDocumentService
	createProjectDocumentService   *project_document_uc.CreateProjectDocumentService
	updateProjectDocumentService   *project_document_uc.UpdateProjectDocumentService
	deleteProjectDocumentService   *project_document_uc.DeleteProjectDocumentService

	// Label services
	listLabelsService    *app_label.ListLabelsService
	findLabelByIDService *app_label.FindLabelService
	createLabelService   *app_label.CreateLabelService
	updateLabelService   *app_label.UpdateLabelService
	deleteLabelService   *app_label.DeleteLabelService

	// Invoice services
	listInvoicesService    *app_invoice.ListInvoicesService
	findInvoiceByIDService *app_invoice.FindInvoiceService
	createInvoiceService   *app_invoice.CreateInvoiceService
	updateInvoiceService   *app_invoice.UpdateInvoiceService
	deleteInvoiceService   *app_invoice.DeleteInvoiceService

	// -- Handlers --
	ClientHandler          *handler.ClientHandler
	ProjectHandler         *handler.ProjectHandler
	ProjectDocumentHandler *handler.ProjectDocumentHandler
	LabelHandler           *handler.LabelHandler
	InvoiceHandler         *handler.InvoiceHandler
}

func NewContainer(db *gorm.DB, parser *go_kit.Parser) *Container {
	c := &Container{
		db:     db,
		parser: parser,
	}

	c.initRepositories()
	c.initServices()
	c.initHandlers()

	return c
}

func (c *Container) initRepositories() {
	c.clientRepository = postgres.NewClientRepository(c.db)
	c.projectRepository = postgres.NewProjectRepository(c.db)
	c.projectUserRepository = postgres.NewProjectUserRepository(c.db)
	c.projectDocumentRepository = postgres.NewProjectDocumentRepository(c.db)
	c.labelRepository = postgres.NewLabelRepository(c.db)
	c.invoiceRepository = postgres.NewInvoiceRepository(c.db)
	c.txManager = postgres.NewTransactionManager(c.db)
}

func (c *Container) initServices() {
	// Client services
	c.createClientService = client.NewCreateClientService(c.clientRepository)
	c.updateClientService = client.NewUpdateClientService(c.clientRepository)
	c.deleteClientService = client.NewDeleteClientService(c.clientRepository)
	c.findClientByIDService = client.NewFindClientByIDService(c.clientRepository)

	// Project services
	c.createProjectService = project.NewCreateProjectService(c.projectRepository, c.projectUserRepository, c.txManager)
	c.updateProjectService = project.NewUpdateProjectService(c.projectRepository, c.txManager)
	c.deleteProjectService = project.NewDeleteProjectService(c.projectRepository, c.txManager)
	c.findProjectByIDService = project.NewFindProjectByIDService(c.projectRepository)
	c.listProjectsService = project.NewListProjectsService(c.projectRepository)

	// ProjectDocument services
	c.createProjectDocumentService = project_document_uc.NewCreateProjectDocumentService(c.projectDocumentRepository, c.txManager)
	c.updateProjectDocumentService = project_document_uc.NewUpdateProjectDocumentService(c.projectDocumentRepository, c.txManager)
	c.deleteProjectDocumentService = project_document_uc.NewDeleteProjectDocumentService(c.projectDocumentRepository, c.txManager)
	c.findProjectDocumentByIDService = project_document_uc.NewFindProjectDocumentService(c.projectDocumentRepository)
	c.listProjectDocumentsService = project_document_uc.NewListProjectDocumentsService(c.projectDocumentRepository)

	// Label services
	c.createLabelService = app_label.NewCreateLabelService(c.labelRepository, c.txManager)
	c.updateLabelService = app_label.NewUpdateLabelService(c.labelRepository, c.txManager)
	c.deleteLabelService = app_label.NewDeleteLabelService(c.labelRepository, c.txManager)
	c.findLabelByIDService = app_label.NewFindLabelService(c.labelRepository)
	c.listLabelsService = app_label.NewListLabelsService(c.labelRepository)

	// Invoice services
	c.createInvoiceService = app_invoice.NewCreateInvoiceService(c.invoiceRepository, c.txManager)
	c.updateInvoiceService = app_invoice.NewUpdateInvoiceService(c.invoiceRepository, c.txManager)
	c.deleteInvoiceService = app_invoice.NewDeleteInvoiceService(c.invoiceRepository, c.txManager)
	c.findInvoiceByIDService = app_invoice.NewFindInvoiceService(c.invoiceRepository)
	c.listInvoicesService = app_invoice.NewListInvoicesService(c.invoiceRepository)
}

func (c *Container) initHandlers() {
	c.ClientHandler = handler.NewClientHandler(c.createClientService, c.updateClientService, c.deleteClientService, c.findClientByIDService, c.listClientsService, c.parser)
	c.ProjectHandler = handler.NewProjectHandler(c.createProjectService, c.updateProjectService, c.deleteProjectService, c.findProjectByIDService, c.listProjectsService, c.parser)
	c.ProjectDocumentHandler = handler.NewProjectDocumentHandler(c.createProjectDocumentService, c.updateProjectDocumentService, c.deleteProjectDocumentService, c.findProjectDocumentByIDService, c.listProjectDocumentsService, c.parser)
	c.LabelHandler = handler.NewLabelHandler(c.createLabelService, c.updateLabelService, c.deleteLabelService, c.findLabelByIDService, c.listLabelsService, c.parser)
	c.InvoiceHandler = handler.NewInvoiceHandler(c.createInvoiceService, c.updateInvoiceService, c.deleteInvoiceService, c.findInvoiceByIDService, c.listInvoicesService, c.parser)
}
