package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type PersistProjectInvitationPort interface {
	FindByToken(ctx context.Context, token string) (*entity.ProjectInvitation, error)
	Save(ctx context.Context, project *entity.ProjectInvitation) error
}

type AcceptProjectInvitationPort interface {
	Execute(ctx context.Context, cmd *command.AcceptProjectInvitationCommand) (*entity.ProjectInvitation, error)
}

type InviteProjectUserPort interface {
	Execute(ctx context.Context, cmd *command.InviteUserToProjectCommand) error
}
