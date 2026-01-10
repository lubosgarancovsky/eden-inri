package dto

import (
	"github.com/google/uuid"
)

type LabelRes struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"projectId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
}

type CreateLabelReq struct {
	ProjectID   uuid.UUID `uri:"projectId" binding:"required,uuid"`
	Name        string    `json:"name" binding:"required,min=1,max=60"`
	Description string    `json:"description"`
	Color       string    `json:"color" binding:"required"`
}

type UpdateLabelReq struct {
	ProjectID   uuid.UUID `uri:"projectId" binding:"required,uuid"`
	LabelID     uuid.UUID `uri:"labelId" binding:"required,uuid"`
	Name        string    `json:"name" binding:"required,min=1,max=60"`
	Description string    `json:"description"`
	Color       string    `json:"color" binding:"required"`
}

type FindLabelByIDReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required,uuid"`
	LabelID   uuid.UUID `uri:"labelId" binding:"required,uuid"`
}

type ListLabelsReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required,uuid"`
}
