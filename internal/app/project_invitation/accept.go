package project_invitation

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type AcceptProjectInvitationService struct {
	repo            ports.PersistProjectInvitationPort
	projectUserRepo ports.PersistProjectUserPort
	txManager       ports.TransactionManager
}

func NewAcceptProjectInvitationService(repo ports.PersistProjectInvitationPort, projectUserRepo ports.PersistProjectUserPort, txManager ports.TransactionManager) *AcceptProjectInvitationService {
	return &AcceptProjectInvitationService{repo, projectUserRepo, txManager}
}

func (s *AcceptProjectInvitationService) Execute(ctx context.Context, cmd *command.AcceptProjectInvitationCommand) (*entity.ProjectInvitation, error) {
	invitation, err := s.repo.FindByToken(ctx, cmd.Token)
	if err != nil {
		return nil, err
	}

	if invitation.UserID != cmd.UserID {
		return nil, go_kit.ErrForbidden.WithMessage("You can't accept invitation for another user")
	}

	invitation.Accept()
	projectUser := entity.NewProjectMember(invitation.UserID, invitation.ProjectID, invitation.Role)

	if err := s.txManager.WithTransaction(ctx, func(txCtx context.Context) error {
		if err := s.projectUserRepo.Insert(txCtx, projectUser); err != nil {
			return err
		}

		if err := s.repo.Save(ctx, invitation); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return invitation, nil
}
