package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProjectRole string

const (
	ProjectRoleOwner     ProjectRole = "owner"
	ProjectRoleAdmin     ProjectRole = "admin"
	ProjectRoleDeveloper ProjectRole = "developer"
	ProjectRoleGuest     ProjectRole = "guest"
)

type Project struct {
	ID             uuid.UUID
	Name           string
	Description    string
	Status         string
	Tags           []string
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastActivityAt time.Time
	StorySequence  int
	Role           ProjectRole // From ProjectUser
	IsStarred      bool        // From ProjectUser
}

type ProjectUser struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	Role      ProjectRole
	IsStarred bool
	JoinedAt  time.Time
}
