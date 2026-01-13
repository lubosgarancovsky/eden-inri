package project_attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListProjectAttachments struct {
	repo         ports.PersistAttachmentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewListProjectAttachments(repo ports.PersistAttachmentPort, isMemberRepo ports.IsProjectMemberPort) *ListProjectAttachments {
	return &ListProjectAttachments{repo, isMemberRepo}
}

func (s *ListProjectAttachments) Execute(ctx context.Context, query *query.ScopedListQuery) (*[]entity.Attachment, int64, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, query.UserID, query.ScopeID)
	if err != nil {
		return nil, 0, err
	}

	if !isMember {
		return nil, 0, app_error.ErrNotAMember
	}

	return s.repo.ListByModelID(ctx, query.ScopeID, query.ListingQuery)
}
