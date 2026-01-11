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

type ProjectLabelHandler struct {
	createUC   ports.CreateProjectLabelUseCase
	updateUC   ports.UpdateProjectLabelUseCase
	deleteUC   ports.DeleteProjectLabelUseCase
	findByIDUC ports.FindProjectLabelByIDUseCase
	listUC     ports.ListProjectLabelsUseCase
	parser     *go_kit.Parser
}

func NewProjectLabelHandler(
	createUC ports.CreateProjectLabelUseCase,
	updateUC ports.UpdateProjectLabelUseCase,
	deleteUC ports.DeleteProjectLabelUseCase,
	findByIDUC ports.FindProjectLabelByIDUseCase,
	listUC ports.ListProjectLabelsUseCase,
	parser *go_kit.Parser,
) *ProjectLabelHandler {
	return &ProjectLabelHandler{
		createUC,
		updateUC,
		deleteUC,
		findByIDUC,
		listUC,
		parser,
	}
}

type ProjectLabelListingAttributes struct {
	Name      string `rsql:"filter,sort"`
	CreatedAt string `rsql:"filter,sort"`
}

func (h *ProjectLabelHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &ProjectLabelListingAttributes{})

	req := &dto.ListProjectLabelsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListProjectScopedQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	items, total, err := h.listUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(items, converter.ToProjectLabelResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *ProjectLabelHandler) FindByID(c *gin.Context) {
	req := &dto.FindProjectLabelByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToFindByIDProjectScopedQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	item, err := h.findByIDUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToProjectLabelResponse(item))
}

func (h *ProjectLabelHandler) Create(c *gin.Context) {
	req := &dto.CreateProjectLabelReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createUC.Execute(c.Request.Context(), converter.ToCreateProjectLabelCommand(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusCreated, converter.ToProjectLabelResponse(created))
}

func (h *ProjectLabelHandler) Update(c *gin.Context) {
	req := &dto.UpdateProjectLabelReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateUC.Execute(c.Request.Context(), converter.ToUpdateProjectLabelCommand(req))
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToProjectLabelResponse(updated))
}

func (h *ProjectLabelHandler) Delete(c *gin.Context) {
	req := &dto.DeleteProjectLabelReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToDeleteProjectScopedCommand(req)
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
