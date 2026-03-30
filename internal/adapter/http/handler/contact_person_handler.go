package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/go-kit"
)

type ContactPersonHandler struct {
	createUC   ports.CreateContactPersonUseCase
	updateUC   ports.UpdateContactPersonUseCase
	deleteUC   ports.DeleteContactPersonUseCase
	findByIDUC ports.FindContactPersonByIDUseCase
	listUC     ports.ListContactPersonsUseCase
	parser     *go_kit.Parser
}

func NewContactPersonHandler(
	createUC ports.CreateContactPersonUseCase,
	updateUC ports.UpdateContactPersonUseCase,
	deleteUC ports.DeleteContactPersonUseCase,
	findByIDUC ports.FindContactPersonByIDUseCase,
	listUC ports.ListContactPersonsUseCase,
	parser *go_kit.Parser,
) *ContactPersonHandler {
	return &ContactPersonHandler{
		createUC,
		updateUC,
		deleteUC,
		findByIDUC,
		listUC,
		parser,
	}
}

type ContactPersonListingAttributes struct {
	Name      string `rsql:"filter,sort"`
	Email     string `rsql:"filter,sort"`
	Phone     string `rsql:"filter,sort"`
	CreatedAt string `rsql:"filter,sort"`
}

func (h *ContactPersonHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &ContactPersonListingAttributes{})

	req := &dto.ListContactPersonsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListContactPersonsQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	items, total, err := h.listUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(items, converter.ToContactPersonResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *ContactPersonHandler) FindByID(c *gin.Context) {
	req := &dto.FindContactPersonByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToFindContactPersonByIDQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findByIDUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToContactPersonResponse(item))
}

func (h *ContactPersonHandler) Create(c *gin.Context) {
	req := &dto.CreateContactPersonReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCreateContactPersonCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, converter.ToContactPersonResponse(created))
}

func (h *ContactPersonHandler) Update(c *gin.Context) {
	req := &dto.UpdateContactPersonReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToUpdateContactPersonCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToContactPersonResponse(updated))
}

func (h *ContactPersonHandler) Delete(c *gin.Context) {
	req := &dto.DeleteContactPersonReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToDeleteContactPersonCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.deleteUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
