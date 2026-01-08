package dto

import (
	"time"

	"github.com/google/uuid"
)

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

type CreateKanbanColumnReq struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Color    string `json:"color"`
	Position int    `json:"position"`
	Type     string `json:"type"`
}

type UpdateKanbanColumnReq struct {
	ColumnID uuid.UUID `uri:"columnId"`
	CreateKanbanColumnReq
}
