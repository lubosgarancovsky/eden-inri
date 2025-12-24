package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
)

type StoryActivityService struct {
	repo               *repository.StoryActivityRepository
	projectUserService *ProjectUserService
}

func NewStoryActivityService(repo *repository.StoryActivityRepository, pus *ProjectUserService) *StoryActivityService {
	return &StoryActivityService{
		repo:               repo,
		projectUserService: pus,
	}
}

func (s *StoryActivityService) InsertActivity(ctx context.Context, userID, storyID uuid.UUID, eventType model.ActivityType, payload json.RawMessage) (*model.StoryActivity, error) {
	activity := &model.StoryActivity{
		StoryID:   storyID,
		ActorID:   userID,
		Type:      eventType,
		Payload:   payload,
		CreatedAt: time.Now(),
	}

	return s.repo.Insert(ctx, activity)
}

func (s *StoryActivityService) ListActivities(ctx context.Context, storyID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.StoryActivity], error) {
	items, totalCount, err := s.repo.ListActivities(ctx, storyID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.StoryActivity]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}
