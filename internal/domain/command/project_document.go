package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateProjectDocumentCommand struct {
	UserID    uuid.UUID
	ProjectID uuid.UUID
	Name      string
	Content   *string
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
	UserID    uuid.UUID
	ProjectID uuid.UUID
	Name      string
	Content   *string
	Tags      []string
}

func (c UpdateProjectDocumentCommand) Apply(document *entity.ProjectDocument) {
	document.Name = c.Name
	document.Content = c.Content
	document.Tags = c.Tags
	document.UpdatedAt = time.Now()
}

type DeleteProjectDocumentCommand struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ProjectID uuid.UUID
}
