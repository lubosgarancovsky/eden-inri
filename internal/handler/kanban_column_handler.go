package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type KanbanColumnHandler struct {
	s *service.KanbanColumnService
}

func NewKanbanColumnHandler(s *service.KanbanColumnService) *KanbanColumnHandler {
	return &KanbanColumnHandler{s: s}
}

// FindAll @Summary      List columns of a kanban board
// @Description  Returns all columns for a given kanban board
// @Tags         Kanban Columns
// @Accept       json
// @Produce      json
// @Param        kanbanId   path      string  true  "Kanban board ID"
// @Success      200  {array}  model.KanbanColumn
// @Router       /v1/inri/kanban/{kanbanId}/columns [get]
// @security GatewayAuth
func (h *KanbanColumnHandler) FindAll(c *gin.Context) {
	boardID := helpers.ExtractID(c, "kanbanId")
	cols, err := h.s.FindAll(c.Request.Context(), boardID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, cols)
}

// FindByID @Summary      Get a column
// @Description  Returns a specific column by ID
// @Tags         Kanban Columns
// @Accept       json
// @Produce      json
// @Param        kanbanId   path      string  true  "Kanban board ID"
// @Param        columnId   path      string  true  "Column ID"
// @Success      200  {object}  model.KanbanColumn
// @Router       /v1/inri/kanban/{kanbanId}/columns/{columnId} [get]
// @security GatewayAuth
func (h *KanbanColumnHandler) FindByID(c *gin.Context) {
	columnID := helpers.ExtractID(c, "columnId")
	boardID := helpers.ExtractID(c, "kanbanId")

	col, err := h.s.FindByID(c.Request.Context(), boardID, columnID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, col)
}

// Insert @Summary      Create a new column
// @Description  Creates a new column in the kanban board
// @Tags         Kanban Columns
// @Accept       json
// @Produce      json
// @Param        kanbanId   path      string  true  "Kanban board ID"
// @Param        body      body      model.KanbanColumnRequest true "Column data"
// @Success      201  {object}  model.KanbanColumn
// @Router       /v1/inri/kanban/{kanbanId}/columns [post]
// @security GatewayAuth
func (h *KanbanColumnHandler) Insert(c *gin.Context) {
	boardID := helpers.ExtractID(c, "kanbanId")

	var req model.KanbanColumnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	col, err := h.s.Insert(c.Request.Context(), boardID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, col)
}

// Update @Summary      Update a column
// @Description  Updates an existing column
// @Tags         Kanban Columns
// @Accept       json
// @Produce      json
// @Param        kanbanId   path      string  true  "Kanban board ID"
// @Param        columnId   path      string  true  "Column ID"
// @Param        body      body      model.KanbanColumnRequest true "Updated column data"
// @Success      200  {object}  model.KanbanColumn
// @Router       /v1/inri/kanban/{kanbanId}/columns/{columnId} [put]
// @security GatewayAuth
func (h *KanbanColumnHandler) Update(c *gin.Context) {
	columnID := helpers.ExtractID(c, "columnId")
	boardID := helpers.ExtractID(c, "kanbanId")

	var req model.KanbanColumnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	col, err := h.s.Update(c.Request.Context(), boardID, columnID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, col)
}

// Delete @Summary      Delete a column
// @Description  Deletes a column from the kanban board
// @Tags         Kanban Columns
// @Accept       json
// @Produce      json
// @Param        kanbanId   path      string  true  "Kanban board ID"
// @Param        columnId   path      string  true  "Column ID"
// @Success      204  {string}  string "No Content"
// @Router       /v1/inri/kanban/{kanbanId}/columns/{columnId} [delete]
// @security GatewayAuth
func (h *KanbanColumnHandler) Delete(c *gin.Context) {
	columnID := helpers.ExtractID(c, "columnId")
	boardID := helpers.ExtractID(c, "kanbanId")

	if err := h.s.Delete(c.Request.Context(), boardID, columnID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
