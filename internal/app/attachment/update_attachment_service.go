package attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type UpdateAttachmentService struct {
	repo ports.PersistAttachmentPort
}

func NewUpdateAttachmentService(repo ports.PersistAttachmentPort) *UpdateAttachmentService {
	return &UpdateAttachmentService{repo: repo}
}

func (s *UpdateAttachmentService) Execute(ctx context.Context, cmd *command.UpdateAttachmentCommand) (*entity.Attachment, error) {
	attachment, err := s.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(attachment)
	if err = s.repo.Update(ctx, attachment); err != nil {
		return nil, err
	}
	return attachment, nil
}
