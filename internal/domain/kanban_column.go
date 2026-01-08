package domain

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
	ID        uuid.UUID
	BoardID   uuid.UUID
	Type      ColumnType
	Key       string
	Name      string
	Color     string
	Position  int
	CreatedAt time.Time
}
