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

var LabelListConfig = &types.ListConfig{
	Filter: map[string]string{
		"id":        "id",
		"projectId": "project_id",
		"boardId":   "board_id",
		"name":      "name",
		"color":     "color",
		"createdAt": "created_at",
	},
	Sort: map[string]string{
		"name":      "name",
		"createdAt": "created_at",
	},
}

type LabelHandler struct {
	parser *rsql.Parser
	s      *service.LabelService
}

type LabelPage struct {
	Items      []model.Label `json:"items"`
	Page       int64         `json:"page"`
	PageSize   int64         `json:"pageSize"`
	TotalCount int64         `json:"totalCount"`
}

func NewLabelHandler(parser *rsql.Parser, s *service.LabelService) *LabelHandler {
	return &LabelHandler{parser, s}
}

// FindAll @Summary      List labels
// @Description  Returns all labels for a project/board
// @Tags         Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        boardId    path      string  false "Board ID"
// @Success      200  {object}  LabelPage
// @Router       /v1/inri/projects/{projectId}/labels [get]
func (h *LabelHandler) FindAll(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")

	lq := helpers.CreateListingQuery(c, h.parser, LabelListConfig.Filter, LabelListConfig.Sort)
	labels, err := h.s.FindAll(c.Request.Context(), projectID, lq)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, labels)
}

// FindByID @Summary      Get a label
// @Description  Returns a label by ID
// @Tags         Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        labelId    path      string  true  "Label ID"
// @Success      200  {object}  model.Label
// @Router       /v1/inri/projects/{projectId}/labels/{labelId} [get]
func (h *LabelHandler) FindByID(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	labelID := helpers.ExtractID(c, "labelId")

	label, err := h.s.FindByID(c.Request.Context(), projectID, labelID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, label)
}

// Insert @Summary      Create a label
// @Description  Creates a new label
// @Tags         Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        body       body      model.LabelRequest true "Label data"
// @Success      201  {object}  model.Label
// @Router       /v1/inri/projects/{projectId}/labels [post]
func (h *LabelHandler) Insert(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")

	var req model.LabelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	label, err := h.s.Insert(c.Request.Context(), projectID, &req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, label)
}

// Update @Summary      Update a label
// @Description  Updates a label
// @Tags         Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        labelId    path      string  true  "Label ID"
// @Param        body       body      model.LabelRequest true "Updated label data"
// @Success      200  {object}  model.Label
// @Router       /v1/inri/projects/{projectId}/labels/{labelId} [put]
func (h *LabelHandler) Update(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	labelID := helpers.ExtractID(c, "labelId")

	var req model.LabelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	label, err := h.s.Update(c.Request.Context(), projectID, labelID, &req)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, label)
}

// Delete @Summary      Delete a label
// @Description  Deletes a label
// @Tags         Labels
// @Accept       json
// @Produce      json
// @Param        projectId  path      string  true  "Project ID"
// @Param        labelId    path      string  true  "Label ID"
// @Success      204  {string} string "No Content"
// @Router       /v1/inri/projects/{projectId}/labels/{labelId} [delete]
func (h *LabelHandler) Delete(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	labelID := helpers.ExtractID(c, "labelId")

	if err := h.s.Delete(c.Request.Context(), projectID, labelID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
