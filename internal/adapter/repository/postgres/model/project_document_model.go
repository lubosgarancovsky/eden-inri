package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type ProjectDocument struct {
	ID        uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	ProjectID uuid.UUID      `gorm:"type:uuid;not null"`
	Name      string         `gorm:"type:string;not null"`
	Content   string         `gorm:"type:text"`
	Tags      pq.StringArray `gorm:"type:text[]"`
	CreatedBy uuid.UUID      `gorm:"type:uuid;not null"`
	CreatedAt time.Time      `gorm:"type:timestamptz;autoCreateTime;not null"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz;index"`
	User      *User          `gorm:"->"`
}

func (ProjectDocument) TableName() string {
	return "inri_project_documents"
}

func (m ProjectDocument) ToDomain() *entity.ProjectDocument {
	return &entity.ProjectDocument{
		ID:        m.ID,
		ProjectID: m.ProjectID,
		Name:      m.Name,
		Content:   m.Content,
		Tags:      m.Tags,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		CreatedBy: m.User.ToDomain(),
	}
}
