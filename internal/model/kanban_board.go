package model

import (
	"time"

	"github.com/google/uuid"
)

type KanbanBoard struct {
	ID             uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	ProjectID      uuid.UUID `json:"projectId"`
	CreatedAt      time.Time `json:"createdAt"`
	LastActivityAt time.Time `json:"lastActivityAt"`
}

type KanbanBoardRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (kb *KanbanBoard) TableName() string {
	return "inri_kanban_boards"
}
