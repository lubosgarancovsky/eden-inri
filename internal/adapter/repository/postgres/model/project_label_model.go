package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type ProjectLabel struct {
	ID          uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ProjectID   *uuid.UUID `gorm:"type:uuid"`
	Name        string     `gorm:"type:string;not null"`
	Description string     `gorm:"type:string"`
	Color       string     `gorm:"type:string;not null"`
	CreatedAt   time.Time  `gorm:"type:timestamptz;autoCreateTime;not null"`
}

func (ProjectLabel) TableName() string {
	return "inri_labels"
}

func (m ProjectLabel) ToDomain() *entity.ProjectLabel {
	return &entity.ProjectLabel{
		ID:          m.ID,
		ProjectID:   m.ProjectID,
		Name:        m.Name,
		Description: m.Description,
		Color:       m.Color,
		CreatedAt:   m.CreatedAt,
	}
}
