package app

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handler"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres"
	"github.com/lubosgarancovsky/eden-inri/internal/app/client"
	"github.com/lubosgarancovsky/eden-inri/internal/app/invoice"
	"github.com/lubosgarancovsky/eden-inri/internal/app/kanban_board"
	"github.com/lubosgarancovsky/eden-inri/internal/app/kanban_column"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project_document"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project_label"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type Container struct {
	db     *gorm.DB
	parser *go_kit.Parser

	// -- Repositories --
	txManager                 *postgres.TransactionManager
	clientRepository          *postgres.ClientRepository
	invoiceRepository         *postgres.InvoiceRepository
	projectRepository         *postgres.ProjectRepository
	projectLabelRepository    *postgres.ProjectLabelRepository
	projectDocumentRepository *postgres.ProjectDocumentRepository
	kanbanBoardRepository     *postgres.KanbanBoardRepository
	kanbanColumnRepository    *postgres.KanbanColumnRepository

	// -- Client services --
	listClientsService    *client.ListClientsService
	findClientByIDService *client.FindClientByIDService
	createClientService   *client.CreateClientService
	updateClientService   *client.UpdateClientService
	deleteClientService   *client.DeleteClientService

	// Invoice services
	listInvoicesService    *invoice.ListInvoicesService
	findInvoiceByIDService *invoice.FindInvoiceService
	createInvoiceService   *invoice.CreateInvoiceService
	updateInvoiceService   *invoice.UpdateInvoiceService
	deleteInvoiceService   *invoice.DeleteInvoiceService

	// Project services
	listProjectsService     *project.ListProjectsService
	findProjectByIDService  *project.FindProjectByIDService
	createProjectService    *project.CreateProjectService
	updateProjectService    *project.UpdateProjectService
	deleteProjectService    *project.DeleteProjectService
	favouriteProjectService *project.FavouriteProjectService

	// Project Label services
	listProjectLabelsService    *project_label.ListProjectLabelsService
	findProjectLabelByIDService *project_label.FindProjectLabelByIDService
	createProjectLabelService   *project_label.CreateProjectLabelService
	updateProjectLabelService   *project_label.UpdateProjectLabelService
	deleteProjectLabelService   *project_label.DeleteProjectLabelService

	// Project Document services
	listProjectDocumentsService    *project_document.ListProjectDocumentsService
	findProjectDocumentByIDService *project_document.FindProjectDocumentByIDService
	createProjectDocumentService   *project_document.CreateProjectDocumentService
	updateProjectDocumentService   *project_document.UpdateProjectDocumentService
	deleteProjectDocumentService   *project_document.DeleteProjectDocumentService

	// Kanban Board services
	listKanbanBoardsService    *kanban_board.ListKanbanBoardsService
	findKanbanBoardByIDService *kanban_board.FindKanbanBoardByIDService
	createKanbanBoardService   *kanban_board.CreateKanbanBoardService
	updateKanbanBoardService   *kanban_board.UpdateKanbanBoardService
	deleteKanbanBoardService   *kanban_board.DeleteKanbanBoardService

	// Kanban Column services
	listKanbanColumnsService    *kanban_column.ListKanbanColumnsService
	findKanbanColumnByIDService *kanban_column.FindKanbanColumnByIDService
	createKanbanColumnService   *kanban_column.CreateKanbanColumnService
	updateKanbanColumnService   *kanban_column.UpdateKanbanColumnService
	deleteKanbanColumnService   *kanban_column.DeleteKanbanColumnService

	// -- Handlers --
	ClientHandler          *handler.ClientHandler
	InvoiceHandler         *handler.InvoiceHandler
	ProjectHandler         *handler.ProjectHandler
	ProjectLabelHandler    *handler.ProjectLabelHandler
	ProjectDocumentHandler *handler.ProjectDocumentHandler
	KanbanBoardHandler     *handler.KanbanBoardHandler
	KanbanColumnHandler    *handler.KanbanColumnHandler
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
	c.txManager = postgres.NewTransactionManager(c.db)
	c.clientRepository = postgres.NewClientRepository(c.db)
	c.invoiceRepository = postgres.NewInvoiceRepository(c.db)
	c.projectRepository = postgres.NewProjectRepository(c.db)
	c.projectLabelRepository = postgres.NewProjectLabelRepository(c.db)
	c.projectDocumentRepository = postgres.NewProjectDocumentRepository(c.db)
	c.kanbanBoardRepository = postgres.NewKanbanBoardRepository(c.db)
	c.kanbanColumnRepository = postgres.NewKanbanColumnRepository(c.db)
}

func (c *Container) initServices() {
	// Client services
	c.createClientService = client.NewCreateClientService(c.clientRepository)
	c.updateClientService = client.NewUpdateClientService(c.clientRepository)
	c.deleteClientService = client.NewDeleteClientService(c.clientRepository)
	c.findClientByIDService = client.NewFindClientByIDService(c.clientRepository)

	// Invoice services
	c.createInvoiceService = invoice.NewCreateInvoiceService(c.invoiceRepository)
	c.updateInvoiceService = invoice.NewUpdateInvoiceService(c.invoiceRepository)
	c.deleteInvoiceService = invoice.NewDeleteInvoiceService(c.invoiceRepository)
	c.findInvoiceByIDService = invoice.NewFindInvoiceService(c.invoiceRepository)
	c.listInvoicesService = invoice.NewListInvoicesService(c.invoiceRepository)

	// Project services
	c.createProjectService = project.NewCreateProjectService(c.projectRepository, c.txManager)
	c.updateProjectService = project.NewUpdateProjectService(c.projectRepository)
	c.deleteProjectService = project.NewDeleteProjectService(c.projectRepository)
	c.findProjectByIDService = project.NewFindProjectByIDService(c.projectRepository)
	c.listProjectsService = project.NewListProjectsService(c.projectRepository)
	c.favouriteProjectService = project.NewFavouriteProjectService(c.projectRepository)

	// Project Label services
	c.createProjectLabelService = project_label.NewCreateProjectLabelService(c.projectLabelRepository)
	c.updateProjectLabelService = project_label.NewUpdateProjectLabelService(c.projectLabelRepository)
	c.deleteProjectLabelService = project_label.NewDeleteProjectLabelService(c.projectLabelRepository)
	c.findProjectLabelByIDService = project_label.NewFindProjectLabelByIDService(c.projectLabelRepository)
	c.listProjectLabelsService = project_label.NewListProjectLabelsService(c.projectLabelRepository)

	// Project Document services
	c.createProjectDocumentService = project_document.NewCreateProjectDocumentService(c.projectDocumentRepository)
	c.updateProjectDocumentService = project_document.NewUpdateProjectDocumentService(c.projectDocumentRepository)
	c.deleteProjectDocumentService = project_document.NewDeleteProjectDocumentService(c.projectDocumentRepository)
	c.findProjectDocumentByIDService = project_document.NewFindProjectDocumentByIDService(c.projectDocumentRepository)
	c.listProjectDocumentsService = project_document.NewListProjectDocumentsService(c.projectDocumentRepository)

	// Kanban Board services
	c.createKanbanBoardService = kanban_board.NewCreateKanbanBoardService(c.kanbanBoardRepository)
	c.updateKanbanBoardService = kanban_board.NewUpdateKanbanBoardService(c.kanbanBoardRepository)
	c.deleteKanbanBoardService = kanban_board.NewDeleteKanbanBoardService(c.kanbanBoardRepository)
	c.findKanbanBoardByIDService = kanban_board.NewFindKanbanBoardByIDService(c.kanbanBoardRepository)
	c.listKanbanBoardsService = kanban_board.NewListKanbanBoardsService(c.kanbanBoardRepository)

	// Kanban Column services
	c.createKanbanColumnService = kanban_column.NewCreateKanbanColumnService(c.kanbanColumnRepository)
	c.updateKanbanColumnService = kanban_column.NewUpdateKanbanColumnService(c.kanbanColumnRepository)
	c.deleteKanbanColumnService = kanban_column.NewDeleteKanbanColumnService(c.kanbanColumnRepository)
	c.findKanbanColumnByIDService = kanban_column.NewFindKanbanColumnByIDService(c.kanbanColumnRepository)
	c.listKanbanColumnsService = kanban_column.NewListKanbanColumnsService(c.kanbanColumnRepository)
}

func (c *Container) initHandlers() {
	c.ClientHandler = handler.NewClientHandler(c.createClientService, c.updateClientService, c.deleteClientService, c.findClientByIDService, c.listClientsService, c.parser)
	c.InvoiceHandler = handler.NewInvoiceHandler(c.createInvoiceService, c.updateInvoiceService, c.deleteInvoiceService, c.findInvoiceByIDService, c.listInvoicesService, c.parser)
	c.ProjectHandler = handler.NewProjectHandler(c.createProjectService, c.updateProjectService, c.deleteProjectService, c.findProjectByIDService, c.listProjectsService, c.favouriteProjectService, c.parser)
	c.ProjectLabelHandler = handler.NewProjectLabelHandler(c.createProjectLabelService, c.updateProjectLabelService, c.deleteProjectLabelService, c.findProjectLabelByIDService, c.listProjectLabelsService, c.parser)
	c.ProjectDocumentHandler = handler.NewProjectDocumentHandler(c.createProjectDocumentService, c.updateProjectDocumentService, c.deleteProjectDocumentService, c.findProjectDocumentByIDService, c.listProjectDocumentsService, c.parser)
	c.KanbanBoardHandler = handler.NewKanbanBoardHandler(c.createKanbanBoardService, c.updateKanbanBoardService, c.deleteKanbanBoardService, c.findKanbanBoardByIDService, c.listKanbanBoardsService, c.parser)
	c.KanbanColumnHandler = handler.NewKanbanColumnHandler(c.createKanbanColumnService, c.updateKanbanColumnService, c.deleteKanbanColumnService, c.findKanbanColumnByIDService, c.listKanbanColumnsService)
}
