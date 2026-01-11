package story

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListAssignedStoriesService struct {
	repo ports.PersistStoryPort
}

func NewListAssignedStoriesService(repo ports.PersistStoryPort) *ListAssignedStoriesService {
	return &ListAssignedStoriesService{repo: repo}
}

func (s *ListAssignedStoriesService) Execute(ctx context.Context, q *query.ListAssignedStoriesQuery) (*[]entity.Story, int64, error) {
	return s.repo.ListAssigned(ctx, q.UserID, q.ListingQuery)
}
