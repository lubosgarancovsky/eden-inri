package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProjectLabel struct {
	ID          uuid.UUID
	ProjectID   uuid.UUID
	Name        string
	Description string
	Color       string
	CreatedAt   time.Time
}
