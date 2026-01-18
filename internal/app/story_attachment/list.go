package story_attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListStoryAttachmentsService struct {
	repo         ports.PersistAttachmentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewListStoryAttachmentsService(repo ports.PersistAttachmentPort, isMemberRepo ports.IsProjectMemberPort) *ListStoryAttachmentsService {
	return &ListStoryAttachmentsService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *ListStoryAttachmentsService) Execute(ctx context.Context, query *query.ListStoryAttachmentsQuery) (*[]entity.Attachment, int64, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, query.UserID, query.ProjectID)
	if err != nil {
		return nil, 0, err
	}

	if !isMember {
		return nil, 0, app_error.ErrNotAMember
	}

	return s.repo.ListByModelID(ctx, query.StoryID, query.ListingQuery)
}
