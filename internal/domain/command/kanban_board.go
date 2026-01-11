package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateKanbanBoardCommand struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	Name      string
	Status    string
}

func (c *CreateKanbanBoardCommand) ToDomain() *entity.KanbanBoard {
	now := time.Now()
	return &entity.KanbanBoard{
		ID:             uuid.New(),
		ProjectID:      c.ProjectID,
		Name:           c.Name,
		Status:         c.Status,
		CreatedAt:      now,
		LastActivityAt: now,
	}
}

type UpdateKanbanBoardCommand struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	UserID    uuid.UUID
	Name      string
	Status    string
}

func (c *UpdateKanbanBoardCommand) Apply(e *entity.KanbanBoard) {
	e.Name = c.Name
	e.Status = c.Status
	e.LastActivityAt = time.Now()
}

type DeleteKanbanBoardCommand struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	UserID    uuid.UUID
}
