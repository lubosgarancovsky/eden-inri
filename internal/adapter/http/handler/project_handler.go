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

type ProjectHandler struct {
	createProjectUC    ports.CreateProjectUseCase
	updateProjectUC    ports.UpdateProjectUseCase
	deleteProjectUC    ports.DeleteProjectUseCase
	findProjectByIDUC  ports.FindProjectByIDUseCase
	listProjectsUC     ports.ListProjectsUseCase
	favouriteProjectUC ports.FavouriteProjectUseCase
	parser             *go_kit.Parser
}

func NewProjectHandler(
	createProjectUC ports.CreateProjectUseCase,
	updateProjectUC ports.UpdateProjectUseCase,
	deleteProjectUC ports.DeleteProjectUseCase,
	findProjectByIDUC ports.FindProjectByIDUseCase,
	listProjectsUC ports.ListProjectsUseCase,
	favouriteProjectUC ports.FavouriteProjectUseCase,
	parser *go_kit.Parser,
) *ProjectHandler {
	return &ProjectHandler{
		createProjectUC,
		updateProjectUC,
		deleteProjectUC,
		findProjectByIDUC,
		listProjectsUC,
		favouriteProjectUC,
		parser,
	}
}

type ProjectListingAttributes struct {
	Name           string `rsql:"filter,sort"`
	Status         string `rsql:"filter"`
	Slug           string `rsql:"filter"`
	CreatedAt      string `rsql:"filter,sort"`
	UpdatedAt      string `rsql:"filter,sort"`
	LastActivityAt string `rsql:"filter,sort"`
	IsStarred      string `rsql:"filter"`
	Role           string `rsql:"filter"`
}

func (h *ProjectHandler) List(c *gin.Context) {
	listingQuery := handle.ListingQuery(c, h.parser, &ProjectListingAttributes{})

	req := &dto.ListProjectsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListQuery(req, listingQuery)
	if err != nil {
		handle.Error(c, err)
		return
	}

	projects, total, err := h.listProjectsUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	responseItems := converter.ToListResponse(projects, converter.ToProjectResponse)
	handle.Page(c, listingQuery, total, responseItems)
}

func (h *ProjectHandler) FindByID(c *gin.Context) {
	req := &dto.FindProjectByIDReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToQuery(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	project, err := h.findProjectByIDUC.Execute(c.Request.Context(), query)
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

	cmd, err := converter.ToCreateProjectCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	created, err := h.createProjectUC.Execute(c.Request.Context(), cmd)
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

	cmd, err := converter.ToUpdateProjectCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.updateProjectUC.Execute(c.Request.Context(), cmd)
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

	cmd, err := converter.ToCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	if err := h.deleteProjectUC.Execute(c.Request.Context(), cmd); err != nil {
		handle.Error(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ProjectHandler) Favourite(c *gin.Context) {
	req := &dto.FavouriteProjectReq{}

	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	cmd, err := converter.ToCommand(req)
	if err != nil {
		handle.Error(c, err)
		return
	}

	updated, err := h.favouriteProjectUC.Execute(c.Request.Context(), cmd)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToProjectResponse(updated))
}
