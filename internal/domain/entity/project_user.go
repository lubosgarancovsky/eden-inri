package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProjectUser struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	Role      ProjectRole
	IsStarred bool
	JoinedAt  time.Time
}

func NewProjectOwner(userID, projectID uuid.UUID) *ProjectUser {
	return &ProjectUser{
		ProjectID: projectID,
		UserID:    userID,
		Role:      ProjectRoleOwner,
		IsStarred: false,
		JoinedAt:  time.Now(),
	}
}
