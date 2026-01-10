package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/services"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

var ClientListConfig = types.ListConfig{
	Filter: map[string]string{
		"clientType":   "client_type",
		"contractType": "contract_type",
		"name":         "name",
	},
	Sort: map[string]string{
		"startedAt":  "started_at",
		"finishedAt": "finished_at",
		"createdAt":  "created_at",
		"name":       "name",
	},
}

type ClientHandler struct {
	clientService *services.ClientService
	parser        *rsql.Parser
}

type ClientPage struct {
	Items      []models.ClientListItem
	Page       int
	PageSize   int
	TotalCount int64
}

func NewClientHandler(parser *rsql.Parser, s *services.ClientService) *ClientHandler {
	return &ClientHandler{s, parser}
}

// FindAll @Summary      List clients
// @Description  Returns a paginated list of all clients
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}   ClientPage
// @Router       /v1/inri/clients [get]
// @security GatewayAuth
func (h *ClientHandler) FindAll(c *gin.Context) {
	helpers.HandleList(c, h.parser, ClientListConfig, h.clientService.FindAll)
}

// FindByID @Summary      Get client by ID
// @Description  Returns a client by its ID
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Success      200  {object}   models.Client
// @Router       /v1/inri/clients/{clientId} [get]
// @security GatewayAuth
func (h *ClientHandler) FindByID(c *gin.Context) {
	helpers.HandleFindByID(c, "clientId", h.clientService.FindByID)
}

// Create @Summary      Create a new client
// @Description  Registers a new OAuth client in the IAM system
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        client  body  models.ClientRequest  true  "Client data"
// @Success      201  {object}  models.Client
// @Router       /v1/inri/clients [post]
// @security GatewayAuth
func (h *ClientHandler) Create(c *gin.Context) {
	helpers.HandleCreate(c, h.clientService.Create)
}

// Update @Summary      Update a client
// @Description  Updates an existing OAuth client in the IAM system
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        client  body  models.ClientRequest  true  "Client data"
// @Param        clientId   path      string  true  "Client ID"
// @Success      200  {object}  models.Client
// @Router       /v1/inri/clients/{clientId} [put]
// @security GatewayAuth
func (h *ClientHandler) Update(c *gin.Context) {
	helpers.HandleUpdate(c, "clientId", h.clientService.Update)
}

// Delete @Summary      Delete a client
// @Description  Deletes the OAuth client from the IAM system
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/clients/{clientId} [delete]
// @security GatewayAuth
func (h *ClientHandler) Delete(c *gin.Context) {
	helpers.HandleDelete(c, "clientId", h.clientService.Delete)
}
