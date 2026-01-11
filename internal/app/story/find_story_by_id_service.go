package story

import (
	"context"
	"fmt"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindStoryByIDService struct {
	repo       ports.PersistStoryPort
	memberRepo ports.MemberHasRolePort
}

func NewFindStoryByIDService(repo ports.PersistStoryPort, memberRepo ports.MemberHasRolePort) *FindStoryByIDService {
	return &FindStoryByIDService{repo: repo, memberRepo: memberRepo}
}

func (s *FindStoryByIDService) Execute(ctx context.Context, q *query.FindStoryByIDQuery) (*entity.Story, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin, entity.ProjectRoleDeveloper, entity.ProjectRoleGuest}
	hasRole, err := s.memberRepo.HasRole(ctx, q.UserID, q.ProjectID, roles)
	if err != nil {
		return nil, err
	}
	if !hasRole {
		return nil, fmt.Errorf("user is not a member of project")
	}

	return s.repo.FindByID(ctx, q.ProjectID, q.ID)
}
