package project_attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindProjectAttachmentByIDService struct {
	repo         ports.PersistAttachmentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewFindProjectAttachmentByIDService(repo ports.PersistAttachmentPort, isMemberRepo ports.IsProjectMemberPort) *FindProjectAttachmentByIDService {
	return &FindProjectAttachmentByIDService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *FindProjectAttachmentByIDService) Execute(ctx context.Context, q *query.ScopedQuery) (*entity.Attachment, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, q.UserID, q.ScopeID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_err.ErrNotAMember
	}

	attachment, err := s.repo.FindByID(ctx, q.ID)
	if err != nil {
		return nil, err
	}

	return attachment, nil
}
