package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type Project struct {
	ID             uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name           string         `gorm:"type:string;not null"`
	Description    *string        `gorm:"type:string"`
	Status         string         `gorm:"type:string"`
	Tags           pq.StringArray `gorm:"type:text[]"`
	Slug           string         `gorm:"type:string;unique"`
	RepositoryURL  *string        `gorm:"type:string"`
	CreatedAt      time.Time      `gorm:"type:timestamptz;autoCreateTime;not null"`
	UpdatedAt      time.Time      `gorm:"type:timestamptz;autoUpdateTime;not null"`
	LastActivityAt time.Time      `gorm:"type:timestamptz"`
	StorySequence  int            `gorm:"type:int;default:0"`
	DeletedAt      gorm.DeletedAt `gorm:"type:timestamptz;index"`

	// Virtual fields from joins
	Role      string `gorm:"->"`
	IsStarred bool   `gorm:"->"`
}

func (Project) TableName() string {
	return "inri_projects"
}

func (p Project) ToDomain() *entity.Project {
	return &entity.Project{
		ID:             p.ID,
		Name:           p.Name,
		Description:    p.Description,
		Status:         p.Status,
		Tags:           p.Tags,
		Slug:           p.Slug,
		RepositoryURL:  p.RepositoryURL,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		LastActivityAt: p.LastActivityAt,
		StorySequence:  p.StorySequence,
		Role:           entity.ProjectRole(p.Role),
		IsStarred:      p.IsStarred,
	}
}

type ProjectUser struct {
	ProjectID uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserID    uuid.UUID `gorm:"primaryKey;type:uuid"`
	Role      string    `gorm:"type:string;not null"`
	IsStarred bool      `gorm:"type:boolean;default:false"`
	JoinedAt  time.Time `gorm:"type:timestamptz;autoCreateTime"`

	// Relations
	User User `gorm:"->"`
}

func (ProjectUser) TableName() string {
	return "inri_project_users"
}

func (pu ProjectUser) ToDomain() *entity.ProjectUser {
	return &entity.ProjectUser{
		ProjectID: pu.ProjectID,
		Role:      entity.ProjectRole(pu.Role),
		IsStarred: pu.IsStarred,
		JoinedAt:  pu.JoinedAt,
		User:      pu.User.ToDomain(),
	}
}
