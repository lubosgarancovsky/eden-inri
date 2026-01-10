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

// Deprecated: use dto.CreateLabelReq and dto.UpdateLabelReq in label_dto.go
