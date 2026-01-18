package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateKanbanBoardReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

type UpdateKanbanBoardReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	BoardID   string `uri:"boardId"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

type DeleteKanbanBoardReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	BoardID   string `uri:"boardId"`
}

type FindKanbanBoardByIDReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	BoardID   string `uri:"boardId"`
}

type ListKanbanBoardsReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
}

type KanbanBoardRes struct {
	ID             uuid.UUID `json:"id"`
	ProjectID      uuid.UUID `json:"projectId"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	LastActivityAt time.Time `json:"lastActivityAt"`
}
