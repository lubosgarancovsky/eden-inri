package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type KanbanStatus string

var (
	KanbanStatusOpen   KanbanStatus = "open"
	KanbanStatusClosed KanbanStatus = "closed"
)

type Kanban struct {
	ID             uuid.UUID
	ProjectID      uuid.UUID
	Name           string
	Status         KanbanStatus
	CreatedAt      time.Time
	LastActivityAt time.Time
}

func NewKanban(projectID uuid.UUID, name string, status KanbanStatus) (*Kanban, error) {
	if projectID == uuid.Nil {
		return nil, errors.New("projectID is required")
	}

	if name == "" {
		return nil, errors.New("name is required")
	}

	if status == "" {
		return nil, errors.New("status is required")
	}

	now := time.Now()

	return &Kanban{
		ID:             uuid.New(),
		ProjectID:      projectID,
		Name:           name,
		Status:         status,
		CreatedAt:      now,
		LastActivityAt: now,
	}, nil
}
