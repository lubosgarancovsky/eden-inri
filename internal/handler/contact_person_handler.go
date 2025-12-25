package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

var ContactPersonListConfig = types.ListConfig{
	Filter: map[string]string{
		"name":  "name",
		"email": "email",
		"phone": "phone",
	},
	Sort: map[string]string{
		"createdAt": "created_at",
		"name":      "name",
	},
}

type ContactPersonHandler struct {
	s      *service.ContactPersonService
	parser *rsql.Parser
}

type ContactPersonPage struct {
	Items      []model.ContactPerson
	Page       int
	PageSize   int
	TotalCount int64
}

func NewContactPersonHandler(parser *rsql.Parser, s *service.ContactPersonService) *ContactPersonHandler {
	return &ContactPersonHandler{s, parser}
}

// FindAll @Summary      List contact persons
// @Description  Returns a paginated list of all contact persons
// @Tags         ContactPersons
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}   ContactPersonPage
// @Router       /v1/inri/clients/{clientId}/contact-persons [get]
// @security GatewayAuth
func (h *ContactPersonHandler) FindAll(c *gin.Context) {
	lq := helpers.CreateListingQuery(c, h.parser, ContactPersonListConfig.Filter, ContactPersonListConfig.Sort)
	userID := helpers.GetUserContext(c).ID
	clientID := helpers.ExtractID(c, "clientId")

	result, err := h.s.FindAll(c.Request.Context(), userID, clientID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// FindByID @Summary      Get contact person by ID
// @Description  Returns a contact person by its ID
// @Tags         ContactPersons
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Param        contactPersonId   path      string  true  "Contact Person ID"
// @Success      200  {object}   model.ContactPerson
// @Router       /v1/inri/clients/{clientId}/contact-persons/{contactPersonId} [get]
// @security GatewayAuth
func (h *ContactPersonHandler) FindByID(c *gin.Context) {
	clientID := helpers.ExtractID(c, "clientId")
	contactPersonID := helpers.ExtractID(c, "contactPersonId")
	userID := helpers.GetUserContext(c).ID

	result, err := h.s.FindByID(c.Request.Context(), userID, clientID, contactPersonID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Create @Summary      Create a new contact person
// @Description  Creates a new contact person
// @Tags         ContactPersons
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Param        contactPerson  body  model.ContactPersonRequest  true  "Contact person data"
// @Success      201  {object}  model.ContactPerson
// @Router       /v1/inri/clients/{clientId}/contact-persons [post]
// @security GatewayAuth
func (h *ContactPersonHandler) Create(c *gin.Context) {
	var input model.ContactPersonRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	userID := helpers.GetUserContext(c).ID
	clientID := helpers.ExtractID(c, "clientId")

	result, err := h.s.Create(c.Request.Context(), userID, clientID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
}

// Update @Summary      Update a contact person
// @Description  Updates an existing contact person
// @Tags         ContactPersons
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Param        contactPersonId   path      string  true  "Contact Person ID"
// @Param        contactPerson  body  model.ContactPersonRequest  true  "Contact person data"
// @Success      200  {object}  model.ContactPerson
// @Router       /v1/inri/clients/{clientId}/contact-persons/{contactPersonId} [put]
// @security GatewayAuth
func (h *ContactPersonHandler) Update(c *gin.Context) {
	var input model.ContactPersonRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	userID := helpers.GetUserContext(c).ID
	clientID := helpers.ExtractID(c, "clientId")
	contactPersonID := helpers.ExtractID(c, "contactPersonId")

	result, err := h.s.Update(c.Request.Context(), userID, clientID, contactPersonID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

// Delete @Summary      Delete a contact person
// @Description  Deletes the contact person
// @Tags         ContactPersons
// @Accept       json
// @Produce      json
// @Param        clientId   path      string  true  "Client ID"
// @Param        contactPersonId   path      string  true  "Contact Person ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/clients/{clientId}/contact-persons/{contactPersonId} [delete]
// @security GatewayAuth
func (h *ContactPersonHandler) Delete(c *gin.Context) {
	userID := helpers.GetUserContext(c).ID
	clientID := helpers.ExtractID(c, "clientId")
	contactPersonID := helpers.ExtractID(c, "contactPersonId")

	if err := h.s.Delete(c.Request.Context(), userID, clientID, contactPersonID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
