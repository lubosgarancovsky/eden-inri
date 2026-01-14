package story_label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListStoryLabelService struct {
	repo         ports.PersistStoryLabelPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewListStoryLabelService(repo ports.PersistStoryLabelPort, isMemberRepo ports.IsProjectMemberPort) *ListStoryLabelService {
	return &ListStoryLabelService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *ListStoryLabelService) Execute(ctx context.Context, query *query.ListStoryLabelsQuery) (*[]entity.StoryLabel, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, query.UserID, query.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_error.ErrNotAMember
	}

	return s.repo.List(ctx, query.ProjectID)
}
