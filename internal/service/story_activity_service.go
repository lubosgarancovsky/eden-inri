package service

import (
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
	kanbanService      *KanbanBoardService
}

func NewStoryActivityService(repo *repository.StoryActivityRepository, pus *ProjectUserService, ks *KanbanBoardService) *StoryActivityService {
	return &StoryActivityService{
		repo:               repo,
		projectUserService: pus,
		kanbanService:      ks,
	}
}

// InsertActivity adds an activity event to a story
func (s *StoryActivityService) InsertActivity(userID, kanbanID, storyID uuid.UUID, eventType model.ActivityType, payload json.RawMessage) (*model.StoryActivity, error) {
	projectID, err := s.kanbanService.GetProjectIDByBoardID(kanbanID)
	if err != nil {
		return nil, err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}

	activity := &model.StoryActivity{
		StoryID:   storyID,
		ActorID:   userID,
		Type:      eventType,
		Payload:   payload,
		CreatedAt: time.Now(),
	}

	return s.repo.Insert(activity)
}

// ListActivities returns all activities for a story
func (s *StoryActivityService) ListActivities(userID, kanbanID, storyID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.StoryActivity], error) {
	projectID, err := s.kanbanService.GetProjectIDByBoardID(kanbanID)
	if err != nil {
		return nil, err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}

	items, totalCount, err := s.repo.ListActivities(storyID, lq)
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
