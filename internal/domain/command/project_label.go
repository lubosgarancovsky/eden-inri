package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateProjectLabelCommand struct {
	ProjectID   uuid.UUID
	UserID      uuid.UUID
	Name        string
	Description string
	Color       string
}

func (c *CreateProjectLabelCommand) ToDomain() *entity.ProjectLabel {
	return &entity.ProjectLabel{
		ID:          uuid.New(),
		ProjectID:   c.ProjectID,
		Name:        c.Name,
		Description: c.Description,
		Color:       c.Color,
		CreatedAt:   time.Now(),
	}
}

type UpdateProjectLabelCommand struct {
	ID          uuid.UUID
	ProjectID   uuid.UUID
	UserID      uuid.UUID
	Name        string
	Description string
	Color       string
}

func (c *UpdateProjectLabelCommand) Apply(e *entity.ProjectLabel) {
	e.Name = c.Name
	e.Description = c.Description
	e.Color = c.Color
}

type DeleteProjectLabelCommand struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	UserID    uuid.UUID
}
