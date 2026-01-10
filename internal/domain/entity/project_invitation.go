package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProjectInvitation struct {
	ID         uuid.UUID
	ProjectID  uuid.UUID
	UserID     uuid.UUID
	InvitedBy  uuid.UUID
	Token      string
	CreatedAt  time.Time
	ExpiresAt  time.Time
	AcceptedAt time.Time
}
