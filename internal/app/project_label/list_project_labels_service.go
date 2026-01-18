package project_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListProjectLabelsService struct {
	repo         ports.PersistProjectLabelPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewListProjectLabelsService(repo ports.PersistProjectLabelPort, isMemberRepo ports.IsProjectMemberPort) *ListProjectLabelsService {
	return &ListProjectLabelsService{repo, isMemberRepo}
}

func (s *ListProjectLabelsService) Execute(ctx context.Context, q *query.ListProjectLabelsQuery) (*[]entity.ProjectLabel, int64, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, q.UserID, q.ProjectID)
	if err != nil {
		return nil, 0, err
	}

	if !isMember {
		return nil, 0, app_error.ErrNotAMember
	}

	return s.repo.List(ctx, q.ProjectID, q.ListingQuery)
}
