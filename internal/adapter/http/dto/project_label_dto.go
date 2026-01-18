package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateProjectLabelReq struct {
	UserID      string `header:"X-User-ID"`
	ProjectID   string `uri:"projectId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type UpdateProjectLabelReq struct {
	UserID      string `header:"X-User-ID"`
	ProjectID   string `uri:"projectId"`
	LabelID     string `uri:"labelId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type DeleteProjectLabelReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	LabelID   string `uri:"labelId"`
}

type FindProjectLabelByIDReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	LabelID   string `uri:"labelId"`
}

type ListProjectLabelsReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
}

type ProjectLabelRes struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"projectId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"createdAt"`
}
