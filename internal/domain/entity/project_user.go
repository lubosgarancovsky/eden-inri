package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProjectUserRole string

const (
	ProjectUserRoleOwner     ProjectUserRole = "owner"
	ProjectUserRoleAdmin     ProjectUserRole = "admin"
	ProjectUserRoleDeveloper ProjectUserRole = "developer"
	ProjectUserRoleGuest     ProjectUserRole = "guest"
)

type ProjectUser struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	Role      string
	IsStarred bool
	JoinedAt  time.Time
}
