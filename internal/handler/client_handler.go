package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/listing"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

type ClientHandler struct {
	s      *service.ClientService
	parser *rsql.Parser
}

type ClientPage struct {
	Items      []model.ClientListItem
	Page       int
	PageSize   int
	TotalCount int64
}

func NewClientHandler(parser *rsql.Parser, s *service.ClientService) *ClientHandler {
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
func (h *ClientHandler) FindAll(c *gin.Context) {
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.ClientFilter, listing.ClientSort)
	if apiErr != nil {
		c.Error(apiErr)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.FindAll(user.ID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// FindByID @Summary      Get client by ID
// @Description  Returns a client by its ID
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Success      200  {object}   model.Client
// @Router       /v1/inri/clients/{clientId} [get]
func (h *ClientHandler) FindByID(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.FindByID(user.ID, UID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Create @Summary      Create a new client
// @Description  Registers a new OAuth client in the IAM system
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        client  body  model.ClientRequest  true  "Client data"
// @Success      201  {object}  model.Client
// @Router       /v1/inri/clients [post]
func (h *ClientHandler) Create(c *gin.Context) {
	var input model.ClientRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Create(user.ID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
}

// Update @Summary      Update a client
// @Description  Updates an existing OAuth client in the IAM system
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        client  body  model.ClientRequest  true  "Client data"
// @Param        clientId   path      string  true  "Client ID"
// @Success      200  {object}  model.Client
// @Router       /v1/inri/clients/{clientId} [put]
func (h *ClientHandler) Update(c *gin.Context) {
	var input model.ClientRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	clientID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Update(user.ID, clientID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Delete @Summary      Delete a client
// @Description  Deletes the OAuth client from the IAM system
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/clients/{clientId} [delete]
func (h *ClientHandler) Delete(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	_, err = h.s.Delete(user.ID, UID)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

// FindAllContactPersons @Summary      List contacts
// @Description  Returns a paginated list of all contacts registered for the client
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}   ContactPersonPage
// @Router       /v1/inri/clients/{clientId}/contact-persons [get]
func (h *ClientHandler) FindAllContactPersons(c *gin.Context) {
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.ContactPersonFilter, listing.ContactPersonSort)
	if apiErr != nil {
		c.Error(apiErr)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	UID, err := helpers.ExtractID(c, "clientId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.FindAllContactPersons(user.ID, UID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}
