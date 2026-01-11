package app

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handler"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres"
	"github.com/lubosgarancovsky/eden-inri/internal/app/client"
	"github.com/lubosgarancovsky/eden-inri/internal/app/invoice"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type Container struct {
	db     *gorm.DB
	parser *go_kit.Parser

	// -- Repositories --
	txManager         *postgres.TransactionManager
	clientRepository  *postgres.ClientRepository
	invoiceRepository *postgres.InvoiceRepository

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

	// -- Handlers --
	ClientHandler  *handler.ClientHandler
	InvoiceHandler *handler.InvoiceHandler
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
}

func (c *Container) initHandlers() {
	c.ClientHandler = handler.NewClientHandler(c.createClientService, c.updateClientService, c.deleteClientService, c.findClientByIDService, c.listClientsService, c.parser)
	c.InvoiceHandler = handler.NewInvoiceHandler(c.createInvoiceService, c.updateInvoiceService, c.deleteInvoiceService, c.findInvoiceByIDService, c.listInvoicesService, c.parser)
}
