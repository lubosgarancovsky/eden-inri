package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectDocumentReq struct {
	UserID    string   `header:"X-User-ID"`
	ProjectID string   `uri:"projectId"`
	Name      string   `json:"name"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
}

type UpdateProjectDocumentReq struct {
	UserID     string   `header:"X-User-ID"`
	ProjectID  string   `uri:"projectId"`
	DocumentID string   `uri:"documentId"`
	Name       string   `json:"name"`
	Content    string   `json:"content"`
	Tags       []string `json:"tags"`
}

type DeleteProjectDocumentReq struct {
	UserID     string `header:"X-User-ID"`
	ProjectID  string `uri:"projectId"`
	DocumentID string `uri:"documentId"`
}

type FindProjectDocumentByIDReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	ID        string `uri:"documentId"`
}

type ListProjectDocumentsReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
}

type ProjectDocumentRes struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"projectId"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	CreatedBy *UserRes  `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
