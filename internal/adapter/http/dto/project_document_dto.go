package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProjectDocumentRes struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"projectId"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateProjectDocumentReq struct {
	ProjectID string   `uri:"projectId"`
	Name      string   `json:"name"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
}

type UpdateProjectDocumentReq struct {
	ID        string   `uri:"projectDocumentId"`
	ProjectID string   `uri:"projectId"`
	Name      string   `json:"name"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
}
