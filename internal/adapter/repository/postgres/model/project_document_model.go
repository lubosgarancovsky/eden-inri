package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type ProjectDocument struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ProjectID uuid.UUID
	Name      string
	Content   string
	Tags      pq.StringArray `gorm:"type:text[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ProjectDocument) TableName() string { return "inri_project_documents" }

func (m *ProjectDocument) ToDomain() *entity.ProjectDocument {
	return &entity.ProjectDocument{
		ID:        m.ID,
		ProjectID: m.ProjectID,
		Name:      m.Name,
		Content:   m.Content,
		Tags:      []string(m.Tags),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ProjectDocumentFromDomain(e *entity.ProjectDocument) *ProjectDocument {
	var tags pq.StringArray
	if e.Tags != nil {
		tags = pq.StringArray(e.Tags)
	}
	return &ProjectDocument{
		ID:        e.ID,
		ProjectID: e.ProjectID,
		Name:      e.Name,
		Content:   e.Content,
		Tags:      tags,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
