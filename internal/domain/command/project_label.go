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
	Description *string
	Color       string
}

func (c *CreateProjectLabelCommand) ToDomain() *entity.ProjectLabel {
	now := time.Now()
	return &entity.ProjectLabel{
		ID:          uuid.New(),
		ProjectID:   c.ProjectID,
		Name:        c.Name,
		Description: c.Description,
		Color:       c.Color,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

type UpdateProjectLabelCommand struct {
	ProjectID   uuid.UUID
	UserID      uuid.UUID
	Name        string
	Description *string
	Color       string
}

func (c *UpdateProjectLabelCommand) Apply(pl *entity.ProjectLabel) {
	now := time.Now()
	pl.Name = c.Name
	pl.Description = c.Description
	pl.Color = c.Color
	pl.UpdatedAt = now
}

type DeleteProjectLabelCommand struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	ID        uuid.UUID
}
