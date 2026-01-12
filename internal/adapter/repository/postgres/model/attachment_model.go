package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type Attachment struct {
	ID           uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null"`
	Model        string         `gorm:"type:string;not null"`
	ModelID      string         `gorm:"type:string;not null"`
	OriginalName string         `gorm:"type:string;not null"`
	ServerName   string         `gorm:"type:string;not null"`
	MimeType     string         `gorm:"type:string;not null"`
	Size         int64          `gorm:"type:bigint;not null"`
	CreatedAt    time.Time      `gorm:"type:timestamptz;autoCreateTime;not null"`
	UpdatedAt    time.Time      `gorm:"type:timestamptz;autoUpdateTime;not null"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamptz;index"`
}

func (Attachment) TableName() string {
	return "inri_attachments"
}

func (a Attachment) ToDomain() *entity.Attachment {
	return &entity.Attachment{
		ID:           a.ID,
		UserID:       a.UserID,
		Model:        a.Model,
		ModelID:      a.ModelID,
		OriginalName: a.OriginalName,
		ServerName:   a.ServerName,
		MimeType:     a.MimeType,
		Size:         a.Size,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
	}
}
