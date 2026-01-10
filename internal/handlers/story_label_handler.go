package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/services"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
)

type StoryLabelHandler struct {
	s *services.StoryLabelService
}

func NewStoryLabelHandler(s *services.StoryLabelService) *StoryLabelHandler {
	return &StoryLabelHandler{s: s}
}

// AssignLabel @Summary      Assign label to a story
// @Description  Assigns a label to a story
// @Tags         Kanban Story Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path  string  true "Project ID"
// @Param        storyId    path  string  true "Story ID"
// @Param        labelId    path  string  true "Label ID"
// @Success      204  {string} string "No Content"
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/labels [post]
// @security GatewayAuth
func (h *StoryLabelHandler) AssignLabel(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	storyID := helpers.ExtractID(c, "storyId")

	var input models.StoryLabelRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.s.AssignLabel(c.Request.Context(), projectID, storyID, &input); err != nil {
		c.Error(err)
		return
	}
	c.Status(204)
}

// UnassignLabel @Summary      Remove label from a story
// @Description  Removes a label from a story
// @Tags         Kanban Story Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path  string  true "Project ID"
// @Param        storyId    path  string  true "Story ID"
// @Param        labelId    path  string  true "Label ID"
// @Success      204  {string} string "No Content"
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/labels/{labelId} [delete]
// @security GatewayAuth
func (h *StoryLabelHandler) UnassignLabel(c *gin.Context) {
	storyID := helpers.ExtractID(c, "storyId")
	labelID := helpers.ExtractID(c, "labelId")

	if err := h.s.UnassignLabel(c.Request.Context(), storyID, labelID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}

// ListLabels @Summary      List labels for a story
// @Description  Returns all labels assigned to a story
// @Tags         Kanban Story Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path  string  true "Project ID"
// @Param        storyId    path  string  true "Story ID"
// @Success      200  {array}  models.Label
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/labels [get]
// @security GatewayAuth
func (h *StoryLabelHandler) ListLabels(c *gin.Context) {
	storyID := helpers.ExtractID(c, "storyId")

	labels, err := h.s.ListLabels(c.Request.Context(), storyID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, labels)
}
