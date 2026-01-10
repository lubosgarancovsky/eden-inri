package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type Project struct {
	ID             uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name           string
	Description    string
	Status         string
	Tags           pq.StringArray `gorm:"type:text[]"`
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastActivityAt time.Time
	StorySequence  int
	IsStarred      bool `gorm:"->"`
}

func (Project) TableName() string { return "inri_projects" }

func (m *Project) ToDomain() *entity.Project {
	return &entity.Project{
		ID:             m.ID,
		Slug:           m.Slug,
		Name:           m.Name,
		Description:    m.Description,
		StorySequence:  m.StorySequence,
		IsStarred:      m.IsStarred,
		Tags:           []string(m.Tags),
		Status:         entity.ProjectStatus(m.Status),
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		LastActivityAt: m.LastActivityAt,
	}
}

func ProjectFromDomain(e *entity.Project) *Project {
	var tags pq.StringArray
	if e.Tags != nil {
		tags = pq.StringArray(e.Tags)
	}
	return &Project{
		ID:             e.ID,
		Name:           e.Name,
		Description:    e.Description,
		Status:         string(e.Status),
		Tags:           tags,
		Slug:           e.Slug,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
		LastActivityAt: e.LastActivityAt,
		StorySequence:  e.StorySequence,
	}
}
