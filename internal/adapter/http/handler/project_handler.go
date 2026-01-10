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

type ProjectHandler struct {
	createUC ports.CreateProjectUseCase
	updateUC ports.UpdateProjectUseCase
	deleteUC ports.DeleteProjectUseCase
	findUC   ports.FindProjectByIDUseCase
	listUC   ports.ListProjectsUseCase
	parser   *go_kit.Parser
}

func NewProjectHandler(createUC ports.CreateProjectUseCase, updateUC ports.UpdateProjectUseCase, deleteUC ports.DeleteProjectUseCase, findUC ports.FindProjectByIDUseCase, listUC ports.ListProjectsUseCase, parser *go_kit.Parser) *ProjectHandler {
	return &ProjectHandler{createUC, updateUC, deleteUC, findUC, listUC, parser}
}

type ProjectListingAttributes struct {
	Name           string `rsql:"filter,sort"`
	Status         string `rsql:"filter"`
	CreatedAt      string `rsql:"filter,sort"`
	UpdatedAt      string `rsql:"filter,sort"`
	LastActivityAt string `rsql:"filter,sort"`
	Slug           string `rsql:"filter,sort"`
}

func (h *ProjectHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &ProjectListingAttributes{})

	req := &dto.ListProjectsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	projects, total, err := h.listUC.Execute(c.Request.Context(), converter.ToListProjectsQuery(req, listingQuery))
	if err != nil {
		handle.Error(c, err)
		return
	}
	handle.Page(c, listingQuery, total, converter.ToProjectListResponse(projects))
}

func (h *ProjectHandler) FindByID(c *gin.Context) {
	req := &dto.FindProjectByIDReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	project, err := h.findUC.Execute(c.Request.Context(), converter.ToFindProjectByIDQuery(req))
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToProjectResponse(project))
}

func (h *ProjectHandler) Create(c *gin.Context) {
	req := &dto.CreateProjectReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	created, err := h.createUC.Execute(c.Request.Context(), converter.ToCreateProjectCommand(req, req.UserID))
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusCreated, converter.ToProjectResponse(created))
}

func (h *ProjectHandler) Update(c *gin.Context) {
	req := &dto.UpdateProjectReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	updated, err := h.updateUC.Execute(c.Request.Context(), converter.ToUpdateProjectCommand(req, req.UserID))
	if err != nil {
		handle.Error(c, err)
		return
	}
	c.JSON(http.StatusOK, converter.ToProjectResponse(updated))
}

func (h *ProjectHandler) Delete(c *gin.Context) {
	req := &dto.DeleteProjectReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}
	if err := h.deleteUC.Execute(c.Request.Context(), converter.ToDeleteProjectCommand(req)); err != nil {
		handle.Error(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
