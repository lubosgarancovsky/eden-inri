package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/services"
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
	s      *services.LabelService
}

type LabelPage struct {
	Items      []models.Label `json:"items"`
	Page       int64          `json:"page"`
	PageSize   int64          `json:"pageSize"`
	TotalCount int64          `json:"totalCount"`
}

func NewLabelHandler(parser *rsql.Parser, s *services.LabelService) *LabelHandler {
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
// @security GatewayAuth
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
// @Success      200  {object}  models.Label
// @Router       /v1/inri/projects/{projectId}/labels/{labelId} [get]
// @security GatewayAuth
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
// @Param        body       body      models.LabelRequest true "Label data"
// @Success      201  {object}  models.Label
// @Router       /v1/inri/projects/{projectId}/labels [post]
// @security GatewayAuth
func (h *LabelHandler) Insert(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")

	var req models.LabelRequest
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
// @Param        body       body      models.LabelRequest true "Updated label data"
// @Success      200  {object}  models.Label
// @Router       /v1/inri/projects/{projectId}/labels/{labelId} [put]
// @security GatewayAuth
func (h *LabelHandler) Update(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	labelID := helpers.ExtractID(c, "labelId")

	var req models.LabelRequest
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
// @security GatewayAuth
func (h *LabelHandler) Delete(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	labelID := helpers.ExtractID(c, "labelId")

	if err := h.s.Delete(c.Request.Context(), projectID, labelID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
