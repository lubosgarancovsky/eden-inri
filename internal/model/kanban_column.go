package model

import (
	"time"

	"github.com/google/uuid"
)

type ColumnType string

const (
	ColumnNormal  ColumnType = "normal"
	ColumnDone    ColumnType = "done"
	ColumnBlocked ColumnType = "blocked"
)

type KanbanColumn struct {
	ID        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	BoardID   uuid.UUID  `gorm:"type:uuid;not null" json:"boardId"`
	Key       string     `json:"key"`
	Name      string     `json:"name"`
	Type      ColumnType `json:"type"`
	Color     string     `json:"color"`
	Position  int        `json:"position"`
	CreatedAt time.Time  `json:"createdAt"`
}

type KanbanColumnRequest struct {
	Key      string     `json:"key"`
	Name     string     `json:"name"`
	Color    string     `json:"color"`
	Position int        `json:"position"`
	Type     ColumnType `json:"type"`
}

func (KanbanColumn) TableName() string {
	return "inri_kanban_columns"
}
