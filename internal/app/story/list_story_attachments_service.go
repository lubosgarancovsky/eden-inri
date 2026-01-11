package story

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type ListStoryAttachmentsService struct {
	repo         ports.PersistAttachmentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewListStoryAttachmentsService(repo ports.PersistAttachmentPort, isMemberRepo ports.IsProjectMemberPort) *ListStoryAttachmentsService {
	return &ListStoryAttachmentsService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *ListStoryAttachmentsService) Execute(ctx context.Context, projectIDStr, storyIDStr, userIDStr string) ([]entity.Attachment, error) {
	projectID := uuid.MustParse(projectIDStr)
	storyID := uuid.MustParse(storyIDStr)
	userID := uuid.MustParse(userIDStr)

	isMember, err := s.isMemberRepo.IsMember(ctx, userID, projectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, nil
	}

	return s.repo.FindByModelID(ctx, userID, storyID, "story")
}
