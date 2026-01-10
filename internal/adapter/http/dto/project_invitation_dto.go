package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProjectInvitationRes struct {
	ID         uuid.UUID `json:"id"`
	ProjectID  uuid.UUID `json:"projectId"`
	UserID     uuid.UUID `json:"userId"`
	Role       string    `json:"role"`
	InvitedBy  uuid.UUID `json:"invitedBy"`
	CreatedAt  time.Time `json:"cratedAt"`
	ExpiresAt  time.Time `json:"expiresAt"`
	AcceptedAt time.Time `json:"acceptedAt"`
}

type CreateProjectInvitationReq struct {
	Role  string `json:"role"`
	Email string `json:"email"`
}

type AcceptProjectInvitationReq struct {
	Token string `json:"token"`
}
