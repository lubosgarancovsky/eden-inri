package model

import (
	"time"

	"github.com/google/uuid"
)

type KanbanBoard struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ProjectID uuid.UUID `json:"projectId"`
	CreatedAt time.Time `json:"createdAt"`
}

func (kb *KanbanBoard) TableName() string {
	return "inri_kanban_boards"
}
