package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
)

type StoryActivityService struct {
	repo               *repositories.StoryActivityRepository
	projectUserService *ProjectUserService
}

func NewStoryActivityService(repo *repositories.StoryActivityRepository, pus *ProjectUserService) *StoryActivityService {
	return &StoryActivityService{
		repo:               repo,
		projectUserService: pus,
	}
}

func (s *StoryActivityService) InsertActivity(ctx context.Context, userID, storyID uuid.UUID, eventType models.ActivityType, payload json.RawMessage) (*models.StoryActivity, error) {
	activity := &models.StoryActivity{
		StoryID:   storyID,
		ActorID:   userID,
		Type:      eventType,
		Payload:   payload,
		CreatedAt: time.Now(),
	}

	return s.repo.Insert(ctx, activity)
}

func (s *StoryActivityService) UpdateActivity(ctx context.Context, userID, storyID, activityID uuid.UUID, input *models.StoryActivityUpdateRequest) (*models.StoryActivity, error) {
	existing, err := s.repo.FindByID(ctx, userID, storyID, activityID)
	if err != nil {
		return nil, err
	}

	if existing.ActorID != userID {
		return nil, api_err.ErrForbidden.WithMessage("Activity can only be updated by its creator")
	}

	if existing.Type != "comment" {
		return nil, api_err.ErrBadRequest.WithMessage("This type of activity cannot be updated")
	}

	activity := &models.StoryActivity{
		ID:      existing.ID,
		StoryID: existing.StoryID,
		ActorID: existing.ActorID,
		Type:    existing.Type,
		Payload: input.Payload,
	}

	return s.repo.Update(ctx, activity)
}

func (s *StoryActivityService) ListActivities(ctx context.Context, storyID uuid.UUID, lq *list.ListingQuery) (*list.Page[models.StoryActivity], error) {
	items, totalCount, err := s.repo.ListActivities(ctx, storyID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[models.StoryActivity]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}
