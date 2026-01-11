package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectLabelReq struct {
	ProjectID   uuid.UUID `uri:"projectId" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Color       string    `json:"color" binding:"required"`
}

type UpdateProjectLabelReq struct {
	ProjectID   uuid.UUID `uri:"projectId" binding:"required"`
	LabelID     uuid.UUID `uri:"labelId" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Color       string    `json:"color" binding:"required"`
}

type DeleteProjectLabelReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
	LabelID   uuid.UUID `uri:"labelId" binding:"required"`
}

type FindProjectLabelByIDReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
	LabelID   uuid.UUID `uri:"labelId" binding:"required"`
}

type ListProjectLabelsReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
}

type ProjectLabelRes struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"projectId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (r CreateProjectLabelReq) GetID() uuid.UUID        { return uuid.Nil }
func (r CreateProjectLabelReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r CreateProjectLabelReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r UpdateProjectLabelReq) GetID() uuid.UUID        { return r.LabelID }
func (r UpdateProjectLabelReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r UpdateProjectLabelReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r DeleteProjectLabelReq) GetID() uuid.UUID        { return r.LabelID }
func (r DeleteProjectLabelReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r DeleteProjectLabelReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r FindProjectLabelByIDReq) GetID() uuid.UUID        { return r.LabelID }
func (r FindProjectLabelByIDReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r FindProjectLabelByIDReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r ListProjectLabelsReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r ListProjectLabelsReq) GetProjectID() uuid.UUID { return r.ProjectID }
