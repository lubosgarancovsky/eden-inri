package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Project struct {
	ID             uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Status         string         `json:"status"`
	Tags           pq.StringArray `gorm:"type:text[]" json:"tags" swaggertype:"array,string"`
	Slug           string         `json:"slug"`
	IsStarred      bool           `json:"isStarred"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	LastActivityAt time.Time      `json:"lastActivityAt"`
	StorySequence  int            `json:"storySequence"`
	Role           ProjectRole    `json:"role"`
}

type ProjectRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	IsStarred   bool     `json:"isStarred"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
	Slug        string   `json:"slug"`
}

type ProjectUser struct {
	ProjectID uuid.UUID   `json:"projectId"`
	UserID    uuid.UUID   `json:"-"`
	User      User        `json:"user"`
	Role      ProjectRole `json:"role"`
	JoinedAt  time.Time   `json:"joinedAt"`
}

type UpdateProjectUserRequest struct {
	Role ProjectRole `json:"role"`
}

type ProjectRole string

const (
	Owner     ProjectRole = "owner"
	Admin     ProjectRole = "admin"
	Developer ProjectRole = "developer"
	Guest     ProjectRole = "guest"
)

func (p *Project) TableName() string {
	return "inri_projects"
}

func (ProjectUser) TableName() string {
	return "inri_project_users"
}
