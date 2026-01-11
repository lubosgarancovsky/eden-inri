package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectDocumentReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Tags      []string  `json:"tags"`
}

type UpdateProjectDocumentReq struct {
	ProjectID  uuid.UUID `uri:"projectId" binding:"required"`
	DocumentID uuid.UUID `uri:"documentId" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	Tags       []string  `json:"tags"`
}

type DeleteProjectDocumentReq struct {
	ProjectID  uuid.UUID `uri:"projectId" binding:"required"`
	DocumentID uuid.UUID `uri:"documentId" binding:"required"`
}

type FindProjectDocumentByIDReq struct {
	ProjectID  uuid.UUID `uri:"projectId" binding:"required"`
	DocumentID uuid.UUID `uri:"documentId" binding:"required"`
}

type ListProjectDocumentsReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
}

type ProjectDocumentRes struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"projectId"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (r CreateProjectDocumentReq) GetID() uuid.UUID        { return uuid.Nil }
func (r CreateProjectDocumentReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r CreateProjectDocumentReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r UpdateProjectDocumentReq) GetID() uuid.UUID        { return r.DocumentID }
func (r UpdateProjectDocumentReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r UpdateProjectDocumentReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r DeleteProjectDocumentReq) GetID() uuid.UUID        { return r.DocumentID }
func (r DeleteProjectDocumentReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r DeleteProjectDocumentReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r FindProjectDocumentByIDReq) GetID() uuid.UUID        { return r.DocumentID }
func (r FindProjectDocumentByIDReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r FindProjectDocumentByIDReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r ListProjectDocumentsReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r ListProjectDocumentsReq) GetProjectID() uuid.UUID { return r.ProjectID }
