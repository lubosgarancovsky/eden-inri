package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type LabelHandler struct {
	s *service.LabelService
}

func NewLabelHandler(s *service.LabelService) *LabelHandler {
	return &LabelHandler{s: s}
}

// FindAll @Summary      List labels
// @Description  Returns all labels for a project/board
// @Tags         Kanban Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        boardId    path      string  false "Board ID"
// @Success      200  {array}  []model.KanbanLabel
// @Router       /v1/inri/projects/{projectId}/labels [get]
func (h *LabelHandler) FindAll(c *gin.Context) {
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
	labels, err := h.s.FindAll(user.ID, projectID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, labels)
}

// FindByID @Summary      Get a label
// @Description  Returns a label by ID
// @Tags         Kanban Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        labelId    path      string  true  "Label ID"
// @Success      200  {object}  model.KanbanLabel
// @Router       /v1/inri/projects/{projectId}/labels/{labelId} [get]
func (h *LabelHandler) FindByID(c *gin.Context) {
	labelID, err := helpers.ExtractID(c, "labelId")
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
	label, err := h.s.FindByID(user.ID, labelID, projectID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, label)
}

// Insert @Summary      Create a label
// @Description  Creates a new label
// @Tags         Kanban Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        body       body      model.KanbanLabelRequest true "Label data"
// @Success      201  {object}  model.KanbanLabel
// @Router       /v1/inri/projects/{projectId}/labels [post]
func (h *LabelHandler) Insert(c *gin.Context) {
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
	var req model.LabelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}
	label, err := h.s.Insert(user.ID, projectID, &req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, label)
}

// Update @Summary      Update a label
// @Description  Updates a label
// @Tags         Kanban Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        labelId    path      string  true  "Label ID"
// @Param        body       body      model.KanbanLabelRequest true "Updated label data"
// @Success      200  {object}  model.KanbanLabel
// @Router       /v1/inri/projects/{projectId}/labels/{labelId} [put]
func (h *LabelHandler) Update(c *gin.Context) {
	labelID, err := helpers.ExtractID(c, "labelId")
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
	var req model.LabelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}
	label, err := h.s.Update(user.ID, labelID, projectID, &req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, label)
}

// Delete @Summary      Delete a label
// @Description  Deletes a label
// @Tags         Kanban Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        labelId    path      string  true  "Label ID"
// @Success      204  {string} string "No Content"
// @Router       /v1/inri/projects/{projectId}/labels/{labelId} [delete]
func (h *LabelHandler) Delete(c *gin.Context) {
	labelID, err := helpers.ExtractID(c, "labelId")
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
	if err := h.s.Delete(user.ID, labelID, projectID); err != nil {
		c.Error(err)
		return
	}
	c.Status(204)
}
