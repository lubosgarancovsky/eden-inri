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
	ProjectID uuid.UUID `uri:"projectId" binding:"required,uuid"`
	Name      string    `json:"name" binding:"required,min=1,max=100"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
}

type UpdateProjectDocumentReq struct {
	DocumentID uuid.UUID `uri:"documentId" binding:"required,uuid"`
	ProjectID  uuid.UUID `uri:"projectId" binding:"required,uuid"`
	Name       string    `json:"name" binding:"required,min=1,max=100"`
	Content    string    `json:"content"`
	Tags       []string  `json:"tags"`
}

type FindProjectDocumentByIDReq struct {
	ProjectID  uuid.UUID `uri:"projectId" binding:"required,uuid"`
	DocumentID uuid.UUID `uri:"documentId" binding:"required,uuid"`
}

type ListProjectDocumentsReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required,uuid"`
}
