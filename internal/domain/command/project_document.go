package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateProjectDocumentCommand struct {
	ProjectID uuid.UUID
	Name      string
	Content   string
	Tags      []string
}

func (c *CreateProjectDocumentCommand) ToDomain() *entity.ProjectDocument {
	now := time.Now()
	return &entity.ProjectDocument{
		ID:        uuid.New(),
		ProjectID: c.ProjectID,
		Name:      c.Name,
		Content:   c.Content,
		Tags:      c.Tags,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type UpdateProjectDocumentCommand struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	Name      string
	Content   string
	Tags      []string
}

func (c *UpdateProjectDocumentCommand) Apply(e *entity.ProjectDocument) {
	e.Name = c.Name
	e.Content = c.Content
	e.Tags = c.Tags
	e.UpdatedAt = time.Now()
}
