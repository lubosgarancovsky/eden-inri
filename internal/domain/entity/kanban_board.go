package entity

import (
	"time"

	"github.com/google/uuid"
)

type KanbanBoard struct {
	ID             uuid.UUID
	Name           string
	Status         string
	ProjectID      uuid.UUID
	CreatedAt      time.Time
	LastActivityAt time.Time
}
