package story_attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type DeleteStoryAttachmentService struct {
	repo         ports.PersistAttachmentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewDeleteStoryAttachmentService(repo ports.PersistAttachmentPort, isMemberRepo ports.IsProjectMemberPort) *DeleteStoryAttachmentService {
	return &DeleteStoryAttachmentService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *DeleteStoryAttachmentService) Execute(ctx context.Context, cmd *command.DeleteStoryAttachmentCommand) error {
	isMember, err := s.isMemberRepo.IsMember(ctx, cmd.UserID, cmd.ProjectID)
	if err != nil {
		return err
	}

	if !isMember {
		return go_kit.ErrForbidden
	}

	return s.repo.Delete(ctx, cmd.AttachmentID)
}
