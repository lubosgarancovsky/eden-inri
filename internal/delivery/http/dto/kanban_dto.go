package dto

import (
	"time"

	"github.com/google/uuid"
)

type KanbanRes struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	ProjectID      uuid.UUID `json:"projectId"`
	CreatedAt      time.Time `json:"createdAt"`
	LastActivityAt time.Time `json:"lastActivityAt"`
}

type CreateKanbanReq struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type UpdateKanbanReq struct {
	KanbanID uuid.UUID `uri:"kanbanId"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
}
