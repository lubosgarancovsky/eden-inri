package story

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindStoryBySlugService struct {
	repo         ports.PersistStoryPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewFindStoryBySlugService(repo ports.PersistStoryPort, isMemberRepo ports.IsProjectMemberPort) *FindStoryBySlugService {
	return &FindStoryBySlugService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *FindStoryBySlugService) Execute(ctx context.Context, query *query.FindStoryBySlugQuery) (*entity.Story, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, query.UserID, query.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_error.ErrNotAMember
	}

	return s.repo.FindBySlug(ctx, query.ProjectID, query.Slug)
}
