package models

import (
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID `json:"-"`
	Model        string    `json:"model"`
	ModelID      string    `json:"modelId"`
	OriginalName string    `json:"originalName"`
	ServerName   string    `json:"serverName"`
	MimeType     string    `json:"mimeType"`
	Size         int64     `json:"size"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (a *Attachment) TableName() string {
	return "inri_attachments"
}
