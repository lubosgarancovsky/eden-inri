package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/listing"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

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
// @Param        projectId  path      string  true  "Project ID"
// @Success      200  {array}  StoryPage
// @Router       /v1/inri/kanban/{kanbanId}/columns/{columnId}/stories [get]
func (h *StoryHandler) FindAll(c *gin.Context) {
	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.StoryFilter, listing.StorySort)
	if apiErr != nil {
		c.Error(apiErr)
		return
	}

	columnID, err := helpers.ExtractID(c, "columnId")
	if err != nil {
		c.Error(err)
		return
	}
	kanbanID, err := helpers.ExtractID(c, "kanbanId")
	if err != nil {
		c.Error(err)
		return
	}
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	stories, err := h.s.FindAll(user.ID, kanbanID, columnID, lq)
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
// @Param        projectId  path      string  true  "Project ID"
// @Success      200  {object}  model.Story
// @Router       /v1/inri/kanban/{kanbanId}/stories/{storyId} [get]
func (h *StoryHandler) FindByID(c *gin.Context) {
	storyID, err := helpers.ExtractID(c, "storyId")
	if err != nil {
		c.Error(err)
		return
	}
	kanbanID, err := helpers.ExtractID(c, "kanbanId")
	if err != nil {
		c.Error(err)
		return
	}
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	story, err := h.s.FindByID(user.ID, storyID, kanbanID)
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
// @Param        kanbanId   path      string  true  "Kanban board ID"
// @Param        columnId   path      string  true  "Column ID"
// @Param        projectId  path      string  true  "Project ID"
// @Param        body       body      model.StoryRequest true "Story data"
// @Success      201  {object}  model.Story
// @Router       /v1/inri/kanban/{kanbanId}/stories [post]
func (h *StoryHandler) Insert(c *gin.Context) {
	kanbanID, err := helpers.ExtractID(c, "kanbanId")
	if err != nil {
		c.Error(err)
		return
	}
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	var req model.StoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	story, err := h.s.Insert(user.ID, kanbanID, &req)
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
// @Param        projectId  path      string  true  "Project ID"
// @Param        body       body      model.StoryRequest true "Updated story data"
// @Success      200  {object}  model.Story
// @Router       /v1/inri/kanban/{kanbanId}/stories [put]
func (h *StoryHandler) Update(c *gin.Context) {
	storyID, err := helpers.ExtractID(c, "storyId")
	if err != nil {
		c.Error(err)
		return
	}
	kanbanID, err := helpers.ExtractID(c, "kanbanId")
	if err != nil {
		c.Error(err)
		return
	}
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	var req model.StoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	story, err := h.s.Update(user.ID, storyID, kanbanID, &req)
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
// @Param        projectId  path      string  true  "Project ID"
// @Success      204  {string} string "No Content"
// @Router       /v1/inri/kanban/{kanbanId}/stories/{storyId} [delete]
func (h *StoryHandler) Delete(c *gin.Context) {
	storyID, err := helpers.ExtractID(c, "storyId")
	if err != nil {
		c.Error(err)
		return
	}
	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.s.Delete(user.ID, storyID, projectID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
