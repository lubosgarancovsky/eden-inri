package app

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handler"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres"
	"github.com/lubosgarancovsky/eden-inri/internal/app/attachment"
	"github.com/lubosgarancovsky/eden-inri/internal/app/client"
	"github.com/lubosgarancovsky/eden-inri/internal/app/contact_person"
	"github.com/lubosgarancovsky/eden-inri/internal/app/invoice"
	"github.com/lubosgarancovsky/eden-inri/internal/app/invoice_stats"
	"github.com/lubosgarancovsky/eden-inri/internal/app/kanban_board"
	"github.com/lubosgarancovsky/eden-inri/internal/app/kanban_column"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project_attachment"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project_document"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project_invitation"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project_label"
	"github.com/lubosgarancovsky/eden-inri/internal/app/project_user"
	"github.com/lubosgarancovsky/eden-inri/internal/app/story"
	"github.com/lubosgarancovsky/eden-inri/internal/app/story_activity"
	"github.com/lubosgarancovsky/eden-inri/internal/app/story_attachment"
	"github.com/lubosgarancovsky/eden-inri/internal/app/story_label"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type Container struct {
	db     *gorm.DB
	parser *go_kit.Parser

	// -- Repositories --
	txManager                   *postgres.TransactionManager
	userRepository              *postgres.UserRepository
	clientRepository            *postgres.ClientRepository
	invoiceRepository           *postgres.InvoiceRepository
	projectRepository           *postgres.ProjectRepository
	projectUserRepository       *postgres.ProjectUserRepository
	projectLabelRepository      *postgres.ProjectLabelRepository
	projectDocumentRepository   *postgres.ProjectDocumentRepository
	kanbanBoardRepository       *postgres.KanbanBoardRepository
	kanbanColumnRepository      *postgres.KanbanColumnRepository
	contactPersonRepository     *postgres.ContactPersonRepository
	storyRepository             *postgres.StoryRepository
	storyActivityRepository     *postgres.StoryActivityRepository
	attachmentRepository        *postgres.AttachmentRepository
	projectInvitationRepository *postgres.ProjectInvitationRepository
	storyLabelRepository        *postgres.StoryLabelRepository
	invoiceStatsRepository      *postgres.InvoiceStatsRepository

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

	// Invoice stats services
	getInvoiceStatsService *invoice_stats.GetInvoiceStatsService

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
	listKanbanColumnsService  *kanban_column.ListKanbanColumnsService
	createKanbanColumnService *kanban_column.CreateKanbanColumnService
	updateKanbanColumnService *kanban_column.UpdateKanbanColumnService
	deleteKanbanColumnService *kanban_column.DeleteKanbanColumnService

	// Contact Person services
	listContactPersonsService    *contact_person.ListContactPersonsService
	findContactPersonByIDService *contact_person.FindContactPersonByIDService
	createContactPersonService   *contact_person.CreateContactPersonService
	updateContactPersonService   *contact_person.UpdateContactPersonService
	deleteContactPersonService   *contact_person.DeleteContactPersonService

	// Story services
	listStoriesUC         *story.ListStoriesService
	findStoryByIDUC       *story.FindStoryByIDService
	createStoryUC         *story.CreateStoryService
	updateStoryUC         *story.UpdateStoryService
	deleteStoryUC         *story.DeleteStoryService
	listAssignedStoriesUC *story.ListAssignedStoriesService
	changeStoryAssigneeUC *story.ChangeStoryAssigneeService

	// Story Activity services
	listStoryActivitiesUC *story_activity.ListStoryActivitiesService
	createStoryActivityUC *story_activity.CreateStoryActivityService
	updateStoryActivityUC *story_activity.UpdateStoryActivityService
	deleteStoryActivityUC *story_activity.DeleteStoryActivityService

	// Attachment services
	uploadAttachmentUC   *attachment.UploadAttachmentService
	listAttachmentsUC    *attachment.ListAttachmentsService
	findAttachmentByIDUC *attachment.FindAttachmentByIDService
	deleteAttachmentUC   *attachment.DeleteAttachmentService
	updateAttachmentUC   *attachment.UpdateAttachmentService

	// Project Attachment services
	listProjectAttachmentsUC    *project_attachment.ListProjectAttachments
	findProjectAttachmentByIDUC *project_attachment.FindProjectAttachmentByIDService
	deleteProjectAttachmentUC   *project_attachment.DeleteProjectAttachmentService
	updateProjectAttachmentUC   *project_attachment.UpdateProjectAttachmentService

	// Story Attachment services
	listStoryAttachmentsUC    *story_attachment.ListStoryAttachmentsService
	findStoryAttachmentByIDUC *story_attachment.FindStoryAttachmentByIDService
	deleteStoryAttachmentUC   *story_attachment.DeleteStoryAttachmentService

	// Project User services
	listProjectUsersUC      *project_user.ListProjectUserService
	deleteProjectUserUC     *project_user.DeleteProjectUserService
	changeProjectUserRoleUC *project_user.ChangeProjectUserRoleService

	// Project invitation services
	inviteUserUC       *project_invitation.InviteUserToProjectService
	acceptInvitationUC *project_invitation.AcceptProjectInvitationService

	// Story label services
	listStoryLabelsUC    *story_label.ListStoryLabelService
	assignStoryLabelUC   *story_label.AssignStoryLabelService
	unassignStoryLabelUC *story_label.UnAssignStoryLabelService

	// -- Handlers --
	ClientHandler            *handler.ClientHandler
	InvoiceHandler           *handler.InvoiceHandler
	ProjectHandler           *handler.ProjectHandler
	ProjectLabelHandler      *handler.ProjectLabelHandler
	ProjectDocumentHandler   *handler.ProjectDocumentHandler
	KanbanBoardHandler       *handler.KanbanBoardHandler
	KanbanColumnHandler      *handler.KanbanColumnHandler
	ContactPersonHandler     *handler.ContactPersonHandler
	StoryHandler             *handler.StoryHandler
	StoryActivityHandler     *handler.StoryActivityHandler
	AttachmentHandler        *handler.AttachmentHandler
	ProjectAttachmentHandler *handler.ProjectAttachmentHandler
	ProjectUserHandler       *handler.ProjectUserHandler
	StoryAttachmentHandler   *handler.StoryAttachmentHandler
	ProjectInvitationHandler *handler.ProjectInvitationHandler
	StoryLabelHandler        *handler.StoryLabelHandler
	InvoiceStatsHandler      *handler.InvoiceStatsHandler
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
	c.userRepository = postgres.NewUserRepository(c.db)
	c.clientRepository = postgres.NewClientRepository(c.db)
	c.invoiceRepository = postgres.NewInvoiceRepository(c.db)
	c.projectRepository = postgres.NewProjectRepository(c.db)
	c.projectUserRepository = postgres.NewProjectUserRepository(c.db)
	c.projectLabelRepository = postgres.NewProjectLabelRepository(c.db)
	c.projectDocumentRepository = postgres.NewProjectDocumentRepository(c.db)
	c.kanbanBoardRepository = postgres.NewKanbanBoardRepository(c.db)
	c.kanbanColumnRepository = postgres.NewKanbanColumnRepository(c.db)
	c.contactPersonRepository = postgres.NewContactPersonRepository(c.db)
	c.storyRepository = postgres.NewStoryRepository(c.db)
	c.storyLabelRepository = postgres.NewStoryLabelRepository(c.db)
	c.storyActivityRepository = postgres.NewStoryActivityRepository(c.db)
	c.attachmentRepository = postgres.NewAttachmentRepository(c.db)
	c.projectInvitationRepository = postgres.NewProjectInvitationRepository(c.db)
	c.invoiceStatsRepository = postgres.NewInvoiceStatsRepository(c.db)
}

func (c *Container) initServices() {
	// Client services
	c.createClientService = client.NewCreateClientService(c.clientRepository)
	c.updateClientService = client.NewUpdateClientService(c.clientRepository)
	c.deleteClientService = client.NewDeleteClientService(c.clientRepository)
	c.findClientByIDService = client.NewFindClientByIDService(c.clientRepository)
	c.listClientsService = client.NewListClientsService(c.clientRepository)

	// Invoice services
	c.createInvoiceService = invoice.NewCreateInvoiceService(c.invoiceRepository)
	c.updateInvoiceService = invoice.NewUpdateInvoiceService(c.invoiceRepository)
	c.deleteInvoiceService = invoice.NewDeleteInvoiceService(c.invoiceRepository)
	c.findInvoiceByIDService = invoice.NewFindInvoiceService(c.invoiceRepository)
	c.listInvoicesService = invoice.NewListInvoicesService(c.invoiceRepository)

	// Invoice stats services
	c.getInvoiceStatsService = invoice_stats.NewGetInvoiceStatsService(c.invoiceStatsRepository)

	// Project services
	c.createProjectService = project.NewCreateProjectService(c.projectRepository, c.projectUserRepository, c.txManager)
	c.updateProjectService = project.NewUpdateProjectService(c.projectRepository)
	c.deleteProjectService = project.NewDeleteProjectService(c.projectRepository)
	c.findProjectByIDService = project.NewFindProjectByIDService(c.projectRepository)
	c.listProjectsService = project.NewListProjectsService(c.projectRepository)
	c.favouriteProjectService = project.NewFavouriteProjectService(c.projectRepository, c.projectUserRepository)

	// Project Label services
	c.createProjectLabelService = project_label.NewCreateProjectLabelService(c.projectLabelRepository, c.projectUserRepository)
	c.updateProjectLabelService = project_label.NewUpdateProjectLabelService(c.projectLabelRepository, c.projectUserRepository)
	c.deleteProjectLabelService = project_label.NewDeleteProjectLabelService(c.projectLabelRepository, c.projectUserRepository)
	c.findProjectLabelByIDService = project_label.NewFindProjectLabelByIDService(c.projectLabelRepository, c.projectUserRepository)
	c.listProjectLabelsService = project_label.NewListProjectLabelsService(c.projectLabelRepository, c.projectUserRepository)

	// Project Document services
	c.createProjectDocumentService = project_document.NewCreateProjectDocumentService(c.projectDocumentRepository, c.projectUserRepository)
	c.updateProjectDocumentService = project_document.NewUpdateProjectDocumentService(c.projectDocumentRepository, c.projectUserRepository)
	c.deleteProjectDocumentService = project_document.NewDeleteProjectDocumentService(c.projectDocumentRepository, c.projectUserRepository)
	c.findProjectDocumentByIDService = project_document.NewFindProjectDocumentByIDService(c.projectDocumentRepository, c.projectUserRepository)
	c.listProjectDocumentsService = project_document.NewListProjectDocumentsService(c.projectDocumentRepository, c.projectUserRepository)

	// Kanban Board services
	c.createKanbanBoardService = kanban_board.NewCreateKanbanBoardService(c.kanbanBoardRepository, c.projectUserRepository)
	c.updateKanbanBoardService = kanban_board.NewUpdateKanbanBoardService(c.kanbanBoardRepository, c.projectUserRepository)
	c.deleteKanbanBoardService = kanban_board.NewDeleteKanbanBoardService(c.kanbanBoardRepository, c.projectUserRepository)
	c.findKanbanBoardByIDService = kanban_board.NewFindKanbanBoardByIDService(c.kanbanBoardRepository, c.projectUserRepository)
	c.listKanbanBoardsService = kanban_board.NewListKanbanBoardsService(c.kanbanBoardRepository, c.projectUserRepository)

	// Kanban Column services
	c.createKanbanColumnService = kanban_column.NewCreateKanbanColumnService(c.kanbanColumnRepository, c.projectUserRepository)
	c.updateKanbanColumnService = kanban_column.NewUpdateKanbanColumnService(c.kanbanColumnRepository, c.projectUserRepository)
	c.deleteKanbanColumnService = kanban_column.NewDeleteKanbanColumnService(c.kanbanColumnRepository, c.projectUserRepository)
	c.listKanbanColumnsService = kanban_column.NewListKanbanColumnsService(c.kanbanColumnRepository, c.projectUserRepository)

	// Contact Person services
	c.createContactPersonService = contact_person.NewCreateContactPersonService(c.contactPersonRepository)
	c.updateContactPersonService = contact_person.NewUpdateContactPersonService(c.contactPersonRepository)
	c.deleteContactPersonService = contact_person.NewDeleteContactPersonService(c.contactPersonRepository)
	c.findContactPersonByIDService = contact_person.NewFindContactPersonByIDService(c.contactPersonRepository)
	c.listContactPersonsService = contact_person.NewListContactPersonsService(c.contactPersonRepository)

	// Story services
	c.createStoryUC = story.NewCreateStoryService(c.storyRepository, c.projectRepository, c.projectUserRepository, c.txManager)
	c.updateStoryUC = story.NewUpdateStoryService(c.storyRepository, c.projectUserRepository)
	c.deleteStoryUC = story.NewDeleteStoryService(c.storyRepository, c.projectUserRepository)
	c.findStoryByIDUC = story.NewFindStoryByIDService(c.storyRepository, c.projectUserRepository)
	c.listStoriesUC = story.NewListStoriesService(c.storyRepository, c.projectUserRepository)
	c.listAssignedStoriesUC = story.NewListAssignedStoriesService(c.storyRepository)
	c.changeStoryAssigneeUC = story.NewChangeStoryAssigneeService(c.storyRepository, c.projectUserRepository)

	// Story label services
	c.listStoryLabelsUC = story_label.NewListStoryLabelService(c.storyLabelRepository, c.projectUserRepository)
	c.assignStoryLabelUC = story_label.NewAssignStoryLabelService(c.storyLabelRepository, c.projectUserRepository)
	c.unassignStoryLabelUC = story_label.NewUnassignStoryLabelService(c.storyLabelRepository, c.projectUserRepository)

	// Story Activity services
	c.createStoryActivityUC = story_activity.NewCreateStoryActivityService(c.storyActivityRepository)
	c.updateStoryActivityUC = story_activity.NewUpdateStoryActivityService(c.storyActivityRepository)
	c.deleteStoryActivityUC = story_activity.NewDeleteStoryActivityService(c.storyActivityRepository)
	c.listStoryActivitiesUC = story_activity.NewListStoryActivitiesService(c.storyActivityRepository)

	// Attachment services
	c.listAttachmentsUC = attachment.NewListAttachmentsService(c.attachmentRepository)
	c.findAttachmentByIDUC = attachment.NewFindAttachmentByIDService(c.attachmentRepository)
	c.deleteAttachmentUC = attachment.NewDeleteAttachmentService(c.attachmentRepository)
	c.updateAttachmentUC = attachment.NewUpdateAttachmentService(c.attachmentRepository)
	c.uploadAttachmentUC = attachment.NewUploadAttachmentService(c.attachmentRepository, c.txManager)

	// Project attachment services
	c.listProjectAttachmentsUC = project_attachment.NewListProjectAttachments(c.attachmentRepository, c.projectUserRepository)
	c.findProjectAttachmentByIDUC = project_attachment.NewFindProjectAttachmentByIDService(c.attachmentRepository, c.projectUserRepository)
	c.updateProjectAttachmentUC = project_attachment.NewUpdateProjectAttachmentService(c.attachmentRepository, c.projectUserRepository)
	c.deleteProjectAttachmentUC = project_attachment.NewDeleteProjectAttachmentService(c.attachmentRepository, c.projectUserRepository)

	// Story attachment services
	c.listStoryAttachmentsUC = story_attachment.NewListStoryAttachmentsService(c.attachmentRepository, c.projectUserRepository)
	c.findStoryAttachmentByIDUC = story_attachment.NewFindStoryAttachmentByIDService(c.attachmentRepository, c.projectUserRepository)
	c.deleteStoryAttachmentUC = story_attachment.NewDeleteStoryAttachmentService(c.attachmentRepository, c.projectUserRepository)

	// Project User services
	c.listProjectUsersUC = project_user.NewListProjectUserService(c.projectUserRepository, c.projectUserRepository)
	c.deleteProjectUserUC = project_user.NewDeleteProjectUserService(c.projectUserRepository, c.projectUserRepository)
	c.changeProjectUserRoleUC = project_user.NewChangeProjectUserRoleService(c.projectUserRepository, c.projectUserRepository)

	// Project Invitation services
	c.acceptInvitationUC = project_invitation.NewAcceptProjectInvitationService(c.projectInvitationRepository, c.projectUserRepository, c.txManager)
	c.inviteUserUC = project_invitation.NewInviteUserToProjectService(c.projectInvitationRepository, c.userRepository, c.projectUserRepository)

	// Story Label services
	c.listStoryLabelsUC = story_label.NewListStoryLabelService(c.storyLabelRepository, c.projectUserRepository)
	c.assignStoryLabelUC = story_label.NewAssignStoryLabelService(c.storyLabelRepository, c.projectUserRepository)
	c.unassignStoryLabelUC = story_label.NewUnassignStoryLabelService(c.storyLabelRepository, c.projectUserRepository)
}

func (c *Container) initHandlers() {
	c.ClientHandler = handler.NewClientHandler(c.createClientService, c.updateClientService, c.deleteClientService, c.findClientByIDService, c.listClientsService, c.parser)
	c.InvoiceHandler = handler.NewInvoiceHandler(c.createInvoiceService, c.updateInvoiceService, c.deleteInvoiceService, c.findInvoiceByIDService, c.listInvoicesService, c.parser)
	c.ProjectHandler = handler.NewProjectHandler(c.createProjectService, c.updateProjectService, c.deleteProjectService, c.findProjectByIDService, c.listProjectsService, c.favouriteProjectService, c.parser)
	c.ProjectLabelHandler = handler.NewProjectLabelHandler(c.createProjectLabelService, c.updateProjectLabelService, c.deleteProjectLabelService, c.findProjectLabelByIDService, c.listProjectLabelsService, c.parser)
	c.ProjectDocumentHandler = handler.NewProjectDocumentHandler(c.createProjectDocumentService, c.updateProjectDocumentService, c.deleteProjectDocumentService, c.findProjectDocumentByIDService, c.listProjectDocumentsService, c.parser)
	c.KanbanBoardHandler = handler.NewKanbanBoardHandler(c.createKanbanBoardService, c.updateKanbanBoardService, c.deleteKanbanBoardService, c.findKanbanBoardByIDService, c.listKanbanBoardsService, c.parser)
	c.KanbanColumnHandler = handler.NewKanbanColumnHandler(c.createKanbanColumnService, c.updateKanbanColumnService, c.deleteKanbanColumnService, c.listKanbanColumnsService)
	c.ContactPersonHandler = handler.NewContactPersonHandler(c.createContactPersonService, c.updateContactPersonService, c.deleteContactPersonService, c.findContactPersonByIDService, c.listContactPersonsService, c.parser)
	c.StoryHandler = handler.NewStoryHandler(c.createStoryUC, c.updateStoryUC, c.deleteStoryUC, c.findStoryByIDUC, c.listStoriesUC, c.listAssignedStoriesUC, c.changeStoryAssigneeUC, c.parser)
	c.StoryLabelHandler = handler.NewStoryLabelHandler(c.listStoryLabelsUC, c.assignStoryLabelUC, c.unassignStoryLabelUC)
	c.StoryActivityHandler = handler.NewStoryActivityHandler(c.createStoryActivityUC, c.updateStoryActivityUC, c.deleteStoryActivityUC, c.listStoryActivitiesUC, c.parser)
	c.AttachmentHandler = handler.NewAttachmentHandler(c.listAttachmentsUC, c.findAttachmentByIDUC, c.deleteAttachmentUC, c.uploadAttachmentUC, c.updateAttachmentUC, c.parser)
	c.ProjectAttachmentHandler = handler.NewProjectAttachmentHandler(c.listProjectAttachmentsUC, c.findProjectAttachmentByIDUC, c.updateProjectAttachmentUC, c.deleteProjectAttachmentUC, c.parser)
	c.ProjectUserHandler = handler.NewProjectUserHandler(c.listProjectUsersUC, c.deleteProjectUserUC, c.changeProjectUserRoleUC, c.parser)
	c.StoryAttachmentHandler = handler.NewStoryAttachmentHandler(c.listStoryAttachmentsUC, c.findStoryAttachmentByIDUC, c.deleteStoryAttachmentUC, c.parser)
	c.ProjectInvitationHandler = handler.NewProjectInvitationHandler(c.inviteUserUC, c.acceptInvitationUC)
	c.StoryLabelHandler = handler.NewStoryLabelHandler(c.listStoryLabelsUC, c.assignStoryLabelUC, c.unassignStoryLabelUC)
	c.InvoiceStatsHandler = handler.NewInvoiceStatsHandler(c.getInvoiceStatsService, c.parser)
}
