package project_user

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListProjectUserService struct {
	repo         ports.PersistProjectUserPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewListProjectUserService(repo ports.PersistProjectUserPort, isMemberRepo ports.IsProjectMemberPort) *ListProjectUserService {
	return &ListProjectUserService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *ListProjectUserService) Execute(ctx context.Context, query *query.ScopedListQuery) (*[]entity.ProjectUser, int64, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, query.UserID, query.ScopeID)
	if err != nil {
		return nil, 0, err
	}

	if !isMember {
		return nil, 0, app_error.ErrNotAMember
	}

	return s.repo.List(ctx, query.ScopeID, query.ListingQuery)
}
