package dto

import (
	"time"

	"github.com/google/uuid"
)

type AttachmentRes struct {
	ID           uuid.UUID `json:"id"`
	Model        string    `json:"model"`
	ModelID      string    `json:"modelId"`
	OriginalName string    `json:"originalName"`
	ServerName   string    `json:"serverName"`
	MimeType     string    `json:"mimeType"`
	Size         int64     `json:"size"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
