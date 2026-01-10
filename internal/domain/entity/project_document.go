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

func NewProjectDocument(projectID uuid.UUID, name string, content string, tags []string) *ProjectDocument {
	now := time.Now()
	return &ProjectDocument{
		ProjectID: projectID,
		Name:      name,
		Content:   content,
		Tags:      tags,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
