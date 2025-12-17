package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
)

type KanbanBoardHandler struct {
	s *service.KanbanBoardService
}

func NewKanbanBoardHandler(s *service.KanbanBoardService) *KanbanBoardHandler {
	return &KanbanBoardHandler{s: s}
}

// FindAll @Summary      List project boards
// @Description  Returns all boards of a project
// @Tags         Boards
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Success      200  {array}  model.KanbanBoard
// @Router       /v1/inri/projects/{projectId}/boards [get]
func (h *KanbanBoardHandler) FindAll(c *gin.Context) {
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

	boards, err := h.s.FindAll(user.ID, projectID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, boards)
}

// FindByID @Summary      Get board detail
// @Description  Returns a single Kanban board
// @Tags         Boards
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        boardId     path      string  true  "Board ID"
// @Success      200  {object}  model.KanbanBoard
// @Router       /v1/inri/projects/{projectId}/boards/{boardId} [get]
func (h *KanbanBoardHandler) FindByID(c *gin.Context) {
	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}

	boardID, err := helpers.ExtractID(c, "boardId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	board, err := h.s.FindByID(user.ID, projectID, boardID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, board)
}

// Insert @Summary      Create a Kanban board
// @Description  Creates a new Kanban board for a project
// @Tags         Boards
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Success      201  {object}  model.KanbanBoard
// @Router       /v1/inri/projects/{projectId}/boards [post]
func (h *KanbanBoardHandler) Insert(c *gin.Context) {
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

	board, err := h.s.Insert(user.ID, projectID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, board)
}

// Delete @Summary      Delete a board
// @Description  Deletes a Kanban board
// @Tags         Boards
// @Accept       json
// @Produce      json
// @Param        projectId   path      string  true  "Project ID"
// @Param        boardId     path      string  true  "Board ID"
// @Success      204  {string}  string  "No Content"
// @Router       /v1/inri/projects/{projectId}/boards/{boardId} [delete]
func (h *KanbanBoardHandler) Delete(c *gin.Context) {
	projectID, err := helpers.ExtractID(c, "projectId")
	if err != nil {
		c.Error(err)
		return
	}

	boardID, err := helpers.ExtractID(c, "boardId")
	if err != nil {
		c.Error(err)
		return
	}

	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.s.Delete(user.ID, projectID, boardID); err != nil {
		c.Error(err)
		return
	}

	c.Status(204)
}
