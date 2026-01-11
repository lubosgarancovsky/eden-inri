package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type KanbanColumn struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	BoardID   uuid.UUID `gorm:"type:uuid;not null"`
	Key       string    `gorm:"type:string;not null"`
	Name      string    `gorm:"type:string;not null"`
	Type      string    `gorm:"type:string;not null"`
	Color     string    `gorm:"type:string"`
	Position  int       `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime;not null"`
}

func (KanbanColumn) TableName() string {
	return "inri_kanban_columns"
}

func (m KanbanColumn) ToDomain() *entity.KanbanColumn {
	return &entity.KanbanColumn{
		ID:        m.ID,
		BoardID:   m.BoardID,
		Key:       m.Key,
		Name:      m.Name,
		Type:      entity.ColumnType(m.Type),
		Color:     m.Color,
		Position:  m.Position,
		CreatedAt: m.CreatedAt,
	}
}
