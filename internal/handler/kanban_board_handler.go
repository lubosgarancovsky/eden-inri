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

var KanbanBoardListConfig = types.ListConfig{
	Filter: map[string]string{
		"id":             "id",
		"name":           "name",
		"status":         "status",
		"projectId":      "project_id",
		"createdAt":      "created_at",
		"lastActivityAt": "last_activity_at",
	},
	Sort: map[string]string{
		"name":           "name",
		"createdAt":      "created_at",
		"lastActivityAt": "last_activity_at",
	},
}

type KanbanBoardHandler struct {
	s      *service.KanbanBoardService
	parser *rsql.Parser
}

func NewKanbanBoardHandler(s *service.KanbanBoardService, parser *rsql.Parser) *KanbanBoardHandler {
	return &KanbanBoardHandler{s, parser}
}

// FindAll @Summary      List project boards
// @Description  Returns all boards of a project
// @Tags         Kanban boards
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {array}  model.KanbanBoard
// @Router       /v1/inri/projects/{projectId}/kanban [get]
// @security GatewayAuth
func (h *KanbanBoardHandler) FindAll(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")

	lq := helpers.CreateListingQuery(c, h.parser, KanbanBoardListConfig.Filter, KanbanBoardListConfig.Sort)
	boards, err := h.s.FindAll(c.Request.Context(), projectID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, boards)
}

// FindByID @Summary      Get board detail
// @Description  Returns a single Kanban board
// @Tags         Kanban boards
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        boardId     path      string  true  "Board ID"
// @Success      200  {object}  model.KanbanBoard
// @Router       /v1/inri/projects/{projectId}/kanban/{kanbanId} [get]
// @security GatewayAuth
func (h *KanbanBoardHandler) FindByID(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	boardID := helpers.ExtractID(c, "kanbanId")

	board, err := h.s.FindByID(c.Request.Context(), projectID, boardID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, board)
}

// Insert @Summary      Create a Kanban board
// @Description  Creates a new Kanban board for a project
// @Tags         Kanban boards
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Success      201  {object}  model.KanbanBoard
// @Router       /v1/inri/projects/{projectId}/kanban [post]
// @security GatewayAuth
func (h *KanbanBoardHandler) Insert(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")

	var req model.KanbanBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	board, err := h.s.Insert(c.Request.Context(), projectID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, board)
}

// Update @Summary      Update the Kanban board
// @Description  Updates the Kanban board
// @Tags         Kanban boards
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        kanbanId   path      string  true  "Kanban ID"
// @Success      201  {object}  model.KanbanBoard
// @Router       /v1/inri/projects/{projectId}/kanban/{kanbanId} [put]
// @security GatewayAuth
func (h *KanbanBoardHandler) Update(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	boardID := helpers.ExtractID(c, "kanbanId")

	var req model.KanbanBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	board, err := h.s.Update(c.Request.Context(), projectID, boardID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, board)
}

// Delete @Summary      Delete a board
// @Description  Deletes a Kanban board
// @Tags         Kanban boards
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        boardId     path      string  true  "Board ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/projects/{projectId}/kanban/{kanbanId} [delete]
// @security GatewayAuth
func (h *KanbanBoardHandler) Delete(c *gin.Context) {
	projectID := helpers.ExtractID(c, "projectId")
	boardID := helpers.ExtractID(c, "kanbanId")

	if err := h.s.Delete(c.Request.Context(), projectID, boardID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
