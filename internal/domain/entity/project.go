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
	Description    *string
	Status         string
	Tags           []string
	Slug           string
	RepositoryURL  *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastActivityAt time.Time
	StorySequence  int
	Role           ProjectRole // From ProjectUser
	IsStarred      bool        // From ProjectUser
}

func (p *Project) IsOwner() bool {
	return p.Role == ProjectRoleOwner
}

func (p *Project) CanMutate() bool {
	return p.Role == ProjectRoleAdmin || p.Role == ProjectRoleOwner
}

func (p *Project) ToggleIsStarred() bool {
	p.IsStarred = !p.IsStarred
	return p.IsStarred
}
