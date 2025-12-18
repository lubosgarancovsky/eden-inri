package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ProjectDocument struct {
	ID        uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ProjectID uuid.UUID      `json:"-"`
	Name      string         `json:"name"`
	Content   string         `json:"content"`
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]" swaggertype:"array,string"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

type ProjectDocumentRequest struct {
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (p *ProjectDocument) TableName() string {
	return "inri_project_documents"
}
