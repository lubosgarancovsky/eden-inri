package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateKanbanColumnReq struct {
	BoardID  string `uri:"boardId"`
	Key      string `json:"key"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Color    string `json:"color"`
	Position int    `json:"position"`
}

type UpdateKanbanColumnReq struct {
	BoardID  string `uri:"boardId"`
	ColumnID string `uri:"columnId"`
	Key      string `json:"key"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Color    string `json:"color"`
	Position int    `json:"position"`
}

type DeleteKanbanColumnReq struct {
	BoardID   string `uri:"boardId"`
	ColumnID  string `uri:"columnId"`
	ProjectID string `uri:"projectId"`
	UserID    string `header:"X-User-ID"`
}

type FindKanbanColumnByIDReq struct {
	BoardID  string `uri:"boardId"`
	ColumnID string `uri:"columnId"`
}

type ListKanbanColumnsReq struct {
	BoardID   string `uri:"boardId"`
	ProjectID string `uri:"projectId"`
	UserID    string `header:"X-User-ID"`
}

type KanbanColumnRes struct {
	ID        uuid.UUID `json:"id"`
	BoardID   uuid.UUID `json:"boardId"`
	Key       string    `json:"key"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Color     string    `json:"color"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"createdAt"`
}
