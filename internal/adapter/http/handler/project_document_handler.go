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

type ProjectDocumentHandler struct {
	createUC   ports.CreateProjectDocumentUseCase
	updateUC   ports.UpdateProjectDocumentUseCase
	deleteUC   ports.DeleteProjectDocumentUseCase
	findByIDUC ports.FindProjectDocumentByIDUseCase
	listUC     ports.ListProjectDocumentsUseCase
	parser     *go_kit.Parser
}

func NewProjectDocumentHandler(
	createUC ports.CreateProjectDocumentUseCase,
	updateUC ports.UpdateProjectDocumentUseCase,
	deleteUC ports.DeleteProjectDocumentUseCase,
	findByIDUC ports.FindProjectDocumentByIDUseCase,
	listUC ports.ListProjectDocumentsUseCase,
	parser *go_kit.Parser,
) *ProjectDocumentHandler {
	return &ProjectDocumentHandler{
		createUC,
		updateUC,
		deleteUC,
		findByIDUC,
		listUC,
		parser,
	}
}

type ProjectDocumentListingAttributes struct {
	Name      string `rsql:"filter,sort"`
	CreatedAt string `rsql:"filter,sort"`
	UpdatedAt string `rsql:"filter,sort"`
	CreatedBy string `rsql:"filter"`
}

func (h *ProjectDocumentHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &ProjectDocumentListingAttributes{})

	req := &dto.ListProjectDocumentsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	items, total, err := h.listUC.Execute(c.Request.Context(), converter.ToListProjectDocumentsQuery(req, listingQuery))
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(items, converter.ToProjectDocumentResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *ProjectDocumentHandler) FindByID(c *gin.Context) {
	req := &dto.FindProjectDocumentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findByIDUC.Execute(c.Request.Context(), converter.ToFindProjectDocumentByIDQuery(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToProjectDocumentResponse(item))
}

func (h *ProjectDocumentHandler) Create(c *gin.Context) {
	req := &dto.CreateProjectDocumentReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createUC.Execute(c.Request.Context(), converter.ToCreateProjectDocumentCommand(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, converter.ToProjectDocumentResponse(created))
}

func (h *ProjectDocumentHandler) Update(c *gin.Context) {
	req := &dto.UpdateProjectDocumentReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateUC.Execute(c.Request.Context(), converter.ToUpdateProjectDocumentCommand(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToProjectDocumentResponse(updated))
}

func (h *ProjectDocumentHandler) Delete(c *gin.Context) {
	req := &dto.DeleteProjectDocumentReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.deleteUC.Execute(c.Request.Context(), converter.ToDeleteProjectDocumentCommand(req)); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
