package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProjectLabelRes struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"projectId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateLabelReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Color       string `json:"color" binding:"required"`
}

type UpdateLabelReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Color       string `json:"color" binding:"required"`
}
