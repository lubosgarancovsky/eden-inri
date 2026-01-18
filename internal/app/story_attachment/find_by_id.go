package story_attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindStoryAttachmentByIDService struct {
	repo         ports.PersistAttachmentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewFindStoryAttachmentByIDService(repo ports.PersistAttachmentPort, isMemberRepo ports.IsProjectMemberPort) *FindStoryAttachmentByIDService {
	return &FindStoryAttachmentByIDService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *FindStoryAttachmentByIDService) Execute(ctx context.Context, query *query.FindStoryAttachmentByIDQuery) (*entity.Attachment, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, query.UserID, query.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_error.ErrNotAMember
	}

	return s.repo.FindByID(ctx, query.AttachmentID)
}
