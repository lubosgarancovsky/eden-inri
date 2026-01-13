package project_attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type UpdateProjectAttachmentService struct {
	repo         ports.PersistAttachmentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewUpdateProjectAttachmentService(repo ports.PersistAttachmentPort, isMemberRepo ports.IsProjectMemberPort) *UpdateProjectAttachmentService {
	return &UpdateProjectAttachmentService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *UpdateProjectAttachmentService) Execute(ctx context.Context, cmd *command.UpdateProjectAttachmentCommand) (*entity.Attachment, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, cmd.UserID, cmd.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_err.ErrNotAMember
	}

	att, err := s.repo.FindByID(ctx, cmd.AttachmentID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(att)

	if err := s.repo.Update(ctx, att); err != nil {
		return nil, err
	}

	return att, nil
}
