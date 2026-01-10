package app

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handler"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres"
	"github.com/lubosgarancovsky/eden-inri/internal/app/client"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type Container struct {
	db *gorm.DB

	parser *go_kit.Parser

	// -- Repositories --
	clientRepository *postgres.ClientRepository

	// -- Services --
	listClientsService    *client.ListClientsService
	findClientByIDService *client.FindClientByIDService
	createClientService   *client.CreateClientService
	updateClientService   *client.UpdateClientService
	deleteClientService   *client.DeleteClientService

	// -- Handlers --
	ClientHandler *handler.ClientHandler
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
}

func (c *Container) initServices() {
	c.createClientService = client.NewCreateClientService(c.clientRepository)
	c.updateClientService = client.NewUpdateClientService(c.clientRepository)
	c.deleteClientService = client.NewDeleteClientService(c.clientRepository)
	c.findClientByIDService = client.NewFindClientByIDService(c.clientRepository)
}

func (c *Container) initHandlers() {
	c.ClientHandler = handler.NewClientHandler(c.createClientService, c.updateClientService, c.deleteClientService, c.findClientByIDService, c.listClientsService, c.parser)
}
