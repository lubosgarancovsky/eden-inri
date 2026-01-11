package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateKanbanColumnCommand struct {
	UserID    uuid.UUID
	ProjectID uuid.UUID
	BoardID   uuid.UUID
	Key       string
	Name      string
	Type      string
	Color     string
	Position  int
}

func (c *CreateKanbanColumnCommand) ToDomain() *entity.KanbanColumn {
	return &entity.KanbanColumn{
		ID:        uuid.New(),
		BoardID:   c.BoardID,
		Key:       c.Key,
		Name:      c.Name,
		Type:      entity.ColumnType(c.Type),
		Color:     c.Color,
		Position:  c.Position,
		CreatedAt: time.Now(),
	}
}

type UpdateKanbanColumnCommand struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProjectID uuid.UUID
	BoardID   uuid.UUID
	Key       string
	Name      string
	Type      string
	Color     string
	Position  int
}

func (c *UpdateKanbanColumnCommand) Apply(e *entity.KanbanColumn) {
	e.Key = c.Key
	e.Name = c.Name
	e.Type = entity.ColumnType(c.Type)
	e.Color = c.Color
	e.Position = c.Position
}

type DeleteKanbanColumnCommand struct {
	ID        uuid.UUID
	BoardID   uuid.UUID
	UserID    uuid.UUID
	ProjectID uuid.UUID
}
