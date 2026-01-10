package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type ProjectDocumentHandler struct {
	createUC ports.CreateProjectDocumentUseCase
	updateUC ports.UpdateProjectDocumentUseCase
	deleteUC ports.DeleteProjectDocumentUseCase
	findUC   ports.FindProjectDocumentByIDUseCase
	listUC   ports.ListProjectDocumentsUseCase
	parser   *go_kit.Parser
}

func NewProjectDocumentHandler(createUC ports.CreateProjectDocumentUseCase, updateUC ports.UpdateProjectDocumentUseCase, deleteUC ports.DeleteProjectDocumentUseCase, findUC ports.FindProjectDocumentByIDUseCase, listUC ports.ListProjectDocumentsUseCase, parser *go_kit.Parser) *ProjectDocumentHandler {
	return &ProjectDocumentHandler{createUC, updateUC, deleteUC, findUC, listUC, parser}
}

type ProjectDocumentListingAttributes struct {
	Name      string `rsql:"filter,sort"`
	CreatedAt string `rsql:"filter,sort"`
	UpdatedAt string `rsql:"filter,sort"`
}

func (h *ProjectDocumentHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &ProjectDocumentListingAttributes{})
	req := &dto.ListProjectDocumentsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	items, total, err := h.listUC.Execute(c.Request.Context(), req.ProjectID, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}
	handle.Page(c, listingQuery, total, converter.ToProjectDocumentListResponse(items))
}

func (h *ProjectDocumentHandler) FindByID(c *gin.Context) {
	req := &dto.FindProjectDocumentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	doc, err := h.findUC.Execute(c.Request.Context(), req.ProjectID, req.DocumentID)
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToProjectDocumentResponse(doc))
}

func (h *ProjectDocumentHandler) Create(c *gin.Context) {
	req := &dto.CreateProjectDocumentReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	created, err := h.createUC.Execute(c.Request.Context(), req.ProjectID, converter.ToProjectDocumentEntityFromCreate(req))
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
	updated, err := h.updateUC.Execute(c.Request.Context(), req.ProjectID, req.DocumentID, converter.ToProjectDocumentEntityFromUpdate(req))
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToProjectDocumentResponse(updated))
}

func (h *ProjectDocumentHandler) Delete(c *gin.Context) {
	req := &dto.FindProjectDocumentByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	if err := h.deleteUC.Execute(c.Request.Context(), req.ProjectID, req.DocumentID); err != nil {
		handle.Error(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
