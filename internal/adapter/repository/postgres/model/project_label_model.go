package model

import (
	"time"

	"github.com/google/uuid"
)

type ProjectLabel struct {
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ProjectID   uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"type:text;not null"`
	Description *string   `gorm:"type:text"`
	Color       string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null"`
}

func (ProjectLabel) TableName() string { return "inri_labels" }
