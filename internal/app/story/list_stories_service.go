package story

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListStoriesService struct {
	repo       ports.PersistStoryPort
	memberRepo ports.MemberHasRolePort
}

func NewListStoriesService(repo ports.PersistStoryPort, memberRepo ports.MemberHasRolePort) *ListStoriesService {
	return &ListStoriesService{repo: repo, memberRepo: memberRepo}
}

func (s *ListStoriesService) Execute(ctx context.Context, q *query.ListStoriesQuery) (*[]entity.Story, int64, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin, entity.ProjectRoleDeveloper, entity.ProjectRoleGuest}
	hasRole, err := s.memberRepo.HasRole(ctx, q.UserID, q.ProjectID, roles)
	if err != nil {
		return nil, 0, err
	}
	if !hasRole {
		return nil, 0, app_err.ErrInsufficientProjectRole
	}

	return s.repo.List(ctx, q.ProjectID, q.ListingQuery)
}
