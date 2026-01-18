package command

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type InviteUserToProjectCommand struct {
	UserID    uuid.UUID
	ProjectID uuid.UUID
	Role      entity.ProjectRole
	Email     string
}

type AcceptProjectInvitationCommand struct {
	UserID uuid.UUID
	Token  string
}
