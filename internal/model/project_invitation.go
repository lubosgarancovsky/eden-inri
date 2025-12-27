package model

import (
	"time"

	"github.com/google/uuid"
)

type ProjectInvitation struct {
	ID         uuid.UUID   `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ProjectID  uuid.UUID   `json:"projectId"`
	UserID     uuid.UUID   `json:"-"`
	Role       ProjectRole `json:"role"`
	Token      string      `json:"_"`
	InvitedBy  uuid.UUID   `json:"invitedBy"`
	CreatedAt  time.Time   `json:"cratedAt"`
	ExpiresAt  time.Time   `json:"expiresAt"`
	AcceptedAt time.Time   `json:"acceptedAt"`
}

type ProjectInvitationRequest struct {
	Role  ProjectRole `json:"role"`
	Email string      `json:"email"`
}

type ProjectInvitationTemplate struct {
	URL         string
	ProjectName string
	UserName    string
}
type AcceptInvitationRequest struct {
	Token string `json:"token"`
}

func (c *ProjectInvitation) TableName() string {
	return "inri_project_invitations"
}
