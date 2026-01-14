package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/go-kit"
)

type ProjectInvitation struct {
	ID         uuid.UUID
	ProjectID  uuid.UUID
	UserID     uuid.UUID
	InvitedBy  uuid.UUID
	Role       ProjectRole
	Token      string
	CreatedAt  time.Time
	ExpiresAt  time.Time
	AcceptedAt *time.Time
}

func NewProjectInvitation(projectID, userID, invitedBy uuid.UUID, role ProjectRole) *ProjectInvitation {
	return &ProjectInvitation{
		ID:         uuid.New(),
		ProjectID:  projectID,
		UserID:     userID,
		InvitedBy:  invitedBy,
		Role:       role,
		Token:      go_kit.OpaqueToken(32),
		CreatedAt:  time.Now(),
		ExpiresAt:  time.Now().Add(time.Second * time.Duration(config.GlobalConfig.InvitationTokenExp)),
		AcceptedAt: nil,
	}
}

func (i *ProjectInvitation) Accept() {
	now := time.Now()
	i.AcceptedAt = &now
}
