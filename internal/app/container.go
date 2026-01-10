package app

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handler"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres"
	"github.com/lubosgarancovsky/eden-inri/internal/app/client"
	"gorm.io/gorm"
)

type Container struct {
	db *gorm.DB

	// -- Repositories --
	clientRepository *postgres.ClientRepository

	// -- Services --
	createClientService *client.CreateClientService
	updateClientService *client.UpdateClientService
	deleteClientService *client.DeleteClientService

	// -- Handlers --
	ClientHandler *handler.ClientHandler
}

func NewContainer(db *gorm.DB) *Container {
	c := &Container{
		db: db,
	}

	c.initRepositories()
	c.initServices()

	return c
}

func (c *Container) initRepositories() {
	c.clientRepository = postgres.NewClientRepository(c.db)
}

func (c *Container) initServices() {
	c.createClientService = client.NewCreateClientService(c.clientRepository)
	c.updateClientService = client.NewUpdateClientService(c.clientRepository)
	c.deleteClientService = client.NewDeleteClientService(c.clientRepository)
}

func (c *Container) initHandlers() {
	c.ClientHandler = handler.NewClientHandler(c.createClientService, c.updateClientService, c.deleteClientService)
}
