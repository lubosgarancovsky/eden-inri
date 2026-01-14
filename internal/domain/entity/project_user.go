package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProjectUser struct {
	ProjectID uuid.UUID
	User      *User
	Role      ProjectRole
	IsStarred bool
	JoinedAt  time.Time
}

func NewProjectOwner(userID, projectID uuid.UUID) *ProjectUser {
	return &ProjectUser{
		ProjectID: projectID,
		User:      &User{ID: userID},
		Role:      ProjectRoleOwner,
		IsStarred: false,
		JoinedAt:  time.Now(),
	}
}

func NewProjectMember(userID, projectID uuid.UUID, role ProjectRole) *ProjectUser {
	return &ProjectUser{
		ProjectID: projectID,
		User:      &User{ID: userID},
		Role:      role,
		IsStarred: false,
		JoinedAt:  time.Now(),
	}
}
