package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProjectInvitationReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type AcceptInvitationReq struct {
	UserID string `header:"X-User-ID"`
	Token  string `uri:"token"`
}

type ProjectInvitationRes struct {
	ID         uuid.UUID  `json:"id"`
	ProjectID  uuid.UUID  `json:"projectId"`
	UserID     uuid.UUID  `json:"userId"`
	InvitedBy  uuid.UUID  `json:"invitedBy"`
	Role       string     `json:"role"`
	CreatedAt  time.Time  `json:"createdAt"`
	ExpiresAt  time.Time  `json:"expiresAt"`
	AcceptedAt *time.Time `json:"acceptedAt"`
}
