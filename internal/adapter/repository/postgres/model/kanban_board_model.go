package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type KanbanBoard struct {
	ID             uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name           string    `gorm:"type:string;not null"`
	Status         string    `gorm:"type:string"`
	ProjectID      uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt      time.Time `gorm:"type:timestamptz;autoCreateTime;not null"`
	LastActivityAt time.Time `gorm:"type:timestamptz;autoUpdateTime;not null"`
}

func (KanbanBoard) TableName() string {
	return "inri_kanban_boards"
}

func (m KanbanBoard) ToDomain() *entity.KanbanBoard {
	return &entity.KanbanBoard{
		ID:             m.ID,
		Name:           m.Name,
		Status:         m.Status,
		ProjectID:      m.ProjectID,
		CreatedAt:      m.CreatedAt,
		LastActivityAt: m.LastActivityAt,
	}
}
