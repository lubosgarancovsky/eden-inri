package project_attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type DeleteProjectAttachmentService struct {
	repo         ports.PersistAttachmentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewDeleteProjectAttachmentService(repo ports.PersistAttachmentPort, isMemberRepo ports.IsProjectMemberPort) *DeleteProjectAttachmentService {
	return &DeleteProjectAttachmentService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *DeleteProjectAttachmentService) Execute(ctx context.Context, cmd *command.ScopedCommand) error {
	isMember, err := s.isMemberRepo.IsMember(ctx, cmd.UserID, cmd.ScopeID)
	if err != nil {
		return err
	}

	if !isMember {
		return app_err.ErrNotAMember
	}

	return s.repo.Delete(ctx, cmd.ID)
}
