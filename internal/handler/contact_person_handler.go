package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/listing"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

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
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}   ContactPersonPage
// @Router       /v1/inri/contact-persons [get]
func (h *ContactPersonHandler) FindAll(c *gin.Context) {
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

	result, err := h.s.FindAll(user.ID, lq)
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
// @Param        contactPersonId   path      string  true  "Contact Person ID"
// @Success      200  {object}   model.ContactPerson
// @Router       /v1/inri/contact-persons/{contactPersonId} [get]
func (h *ContactPersonHandler) FindByID(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "contactPersonId")
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

// Create @Summary      Create a new contact person
// @Description  Creates a new contact person
// @Tags         ContactPersons
// @Accept       json
// @Produce      json
// @Param        contactPerson  body  model.ContactPersonRequest  true  "Contact person data"
// @Success      201  {object}  model.ContactPerson
// @Router       /v1/inri/contact-persons [post]
func (h *ContactPersonHandler) Create(c *gin.Context) {
	var input model.ContactPersonRequest
	if err := c.ShouldBindJSON(&input); err != nil {
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

// Update @Summary      Update a contact person
// @Description  Updates an existing contact person
// @Tags         ContactPersons
// @Accept       json
// @Produce      json
// @Param        contactPerson  body  model.ContactPersonRequest  true  "Contact person data"
// @Param        contactPersonId   path      string  true  "Contact Person ID"
// @Success      200  {object}  model.ContactPerson
// @Router       /v1/inri/contact-persons/{contactPersonId} [put]
func (h *ContactPersonHandler) Update(c *gin.Context) {
	var input model.ContactPersonRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	UID, err := helpers.ExtractID(c, "contactPersonId")
	if err != nil {
		c.Error(err)
		return
	}

	result, err := h.s.Update(user.ID, UID, &input)
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
// @Param        contactPersonId   path      string  true  "Contact Person ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/contact-persons/{contactPersonId} [delete]
func (h *ContactPersonHandler) Delete(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "contactPersonId")
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
