package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

var StoryListConfig = types.ListConfig{
	Filter: map[string]string{
		"id":         "id",
		"projectId":  "project_id",
		"boardId":    "board_id",
		"columnId":   "column_id",
		"slug":       "slug",
		"title":      "title",
		"kind":       "kind",
		"assigneeId": "assignee_id",
		"priority":   "priority",
		"startDate":  "start_date",
		"endDate":    "end_date",
		"createdAt":  "created_at",
	},
	Sort: map[string]string{
		"position":  "position",
		"priority":  "priority",
		"createdAt": "created_at",
		"updatedAt": "updated_at",
		"title":     "title",
		"slug":      "slug",
	},
}

type StoryHandler struct {
	parser *rsql.Parser
	s      *service.StoryService
}

type StoryPage struct {
	Items      []model.Story
	Page       int
	PageSize   int
	TotalCount int64
}

func NewStoryHandler(p *rsql.Parser, s *service.StoryService) *StoryHandler {
	return &StoryHandler{s: s, parser: p}
}

// FindAll @Summary      List stories in a column
// @Description  Returns all stories for a column
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        kanbanId   path      string  true  "Kanban board ID"
// @Param        columnId   path      string  true  "Column ID"
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {array}  StoryPage
// @Router       /v1/inri/projects/{projectId}/stories [get]
// @security GatewayAuth
func (h *StoryHandler) FindAll(c *gin.Context) {
	lq := helpers.CreateListingQuery(c, h.parser, StoryListConfig.Filter, StoryListConfig.Sort)
	projectID := helpers.ExtractID(c, "projectId")

	stories, err := h.s.FindAll(c.Request.Context(), projectID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, stories)
}

// FindByID @Summary      Get story by ID
// @Description  Returns a specific story by ID
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        storyId    path      string  true  "Story ID"
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {object}  model.Story
// @Router       /v1/inri/projects/{projectId}/stories/{storyId} [get]
// @security GatewayAuth
func (h *StoryHandler) FindByID(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	story, err := h.s.FindByID(c.Request.Context(), projectID, storyID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, story)
}

// Insert @Summary      Create a new story
// @Description  Creates a new story in a column
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        body       body      model.StoryRequest true "Story data"
// @Success      201  {object}  model.Story
// @Router       /v1/inri/projects/{projectId}/stories [post]
// @security GatewayAuth
func (h *StoryHandler) Insert(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")

	var req model.StoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	story, err := h.s.Insert(c.Request.Context(), projectID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, story)
}

// Update @Summary      Update a story
// @Description  Updates an existing story
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        storyId    path      string  true  "Story ID"
// @Param        projectId   path      string  true  "Project ID"
// @Param        body       body      model.StoryRequest true "Updated story data"
// @Success      200  {object}  model.Story
// @Router       /v1/inri/projects/{projectId}/stories/{storyId} [put]
// @security GatewayAuth
func (h *StoryHandler) Update(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	var req model.StoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	story, err := h.s.Update(c.Request.Context(), projectID, storyID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, story)
}

// Delete @Summary      Delete a story
// @Description  Deletes a story
// @Tags         Kanban Stories
// @Accept       json
// @Produce      json
// @Param        storyId    path      string  true  "Story ID"
// @Param        projectId   path      string  true  "Project ID"
// @Success      204  {string} string "No Content"
// @Router       /v1/inri/projects/{projectId}/stories/{storyId} [delete]
// @security GatewayAuth
func (h *StoryHandler) Delete(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	if err := h.s.Delete(c.Request.Context(), projectID, storyID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
