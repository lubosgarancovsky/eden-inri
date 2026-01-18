package story_activity

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListStoryActivitiesService struct {
	repo ports.PersistStoryActivityPort
}

func NewListStoryActivitiesService(repo ports.PersistStoryActivityPort) *ListStoryActivitiesService {
	return &ListStoryActivitiesService{repo: repo}
}

func (s *ListStoryActivitiesService) Execute(ctx context.Context, q *query.ListStoryActivitiesQuery) (*[]entity.StoryActivity, int64, error) {
	return s.repo.List(ctx, q.StoryID, q.ListingQuery)
}
