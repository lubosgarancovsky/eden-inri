package model

import (
	"time"

	"github.com/google/uuid"
)

type ProjectUser struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	Role      string
	IsStarred bool
	JoinedAt  time.Time
}

func (ProjectUser) TableName() string { return "inri_project_users" }
