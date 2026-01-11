package entity

import (
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Model        string
	ModelID      string
	OriginalName string
	ServerName   string
	MimeType     string
	Size         int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
