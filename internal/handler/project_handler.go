package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/listing"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

type ProjectHandler struct {
	s      *service.ProjectService
	parser *rsql.Parser
}

type ProjectPage struct {
	Items      []model.Project
	Page       int
	PageSize   int
	TotalCount int64
}

func NewProjectHandler(parser *rsql.Parser, s *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{s, parser}
}

// FindAll @Summary      List projects
// @Description  Returns a paginated list of all projects
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}   ProjectPage
// @Router       /v1/inri/projects [get]
func (h *ProjectHandler) FindAll(c *gin.Context) {
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.ProjectFilter, listing.ProjectSort)
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

// FindByID @Summary      Get project by ID
// @Description  Returns a project by its ID
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object}   model.Project
// @Router       /v1/inri/projects/{projectId} [get]
func (h *ProjectHandler) FindByID(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "projectId")
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

// Create @Summary      Create a new project
// @Description  Creates a new project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project  body  model.ProjectRequest  true  "Project data"
// @Success      201  {object}  model.Project
// @Router       /v1/inri/projects [post]
func (h *ProjectHandler) Create(c *gin.Context) {
	var input model.ProjectRequest
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

// Update @Summary      Update a project
// @Description  Updates an existing project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        project  body  model.ProjectRequest  true  "Project data"
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object}  model.Project
// @Router       /v1/inri/projects/{projectId} [put]
func (h *ProjectHandler) Update(c *gin.Context) {
	var input model.ProjectRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	UID, err := helpers.ExtractID(c, "projectId")
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

// Delete @Summary      Delete a project
// @Description  Deletes the project
// @Tags         Projects
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/projects/{projectId} [delete]
func (h *ProjectHandler) Delete(c *gin.Context) {
	UID, err := helpers.ExtractID(c, "projectId")
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
