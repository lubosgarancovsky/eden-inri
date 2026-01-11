package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProjectDocument struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	Name      string
	Content   string
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
