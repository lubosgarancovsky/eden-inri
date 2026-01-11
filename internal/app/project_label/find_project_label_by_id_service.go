package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindProjectLabelByIDService struct {
	repo         ports.PersistProjectLabelPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewFindProjectLabelByIDService(repo ports.PersistProjectLabelPort, isMemberRepo ports.IsProjectMemberPort) *FindProjectLabelByIDService {
	return &FindProjectLabelByIDService{repo: repo}
}

func (s *FindProjectLabelByIDService) Execute(ctx context.Context, q *query.FindProjectLabelByIDQuery) (*entity.ProjectLabel, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, q.UserID, q.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_error.ErrNotAMember
	}

	return s.repo.FindByID(ctx, q.ProjectID, q.ID)
}
