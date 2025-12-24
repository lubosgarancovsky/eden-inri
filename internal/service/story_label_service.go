package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
)

type StoryLabelService struct {
	repo               *repository.StoryLabelRepository
	projectUserService *ProjectUserService
	kanbanService      *KanbanBoardService
}

func NewStoryLabelService(repo *repository.StoryLabelRepository, pus *ProjectUserService, ks *KanbanBoardService) *StoryLabelService {
	return &StoryLabelService{repo: repo, projectUserService: pus, kanbanService: ks}
}

func (s *StoryLabelService) AssignLabel(ctx context.Context, storyID, labelID uuid.UUID) error {
	return s.repo.AssignLabel(ctx, storyID, labelID)
}

func (s *StoryLabelService) UnassignLabel(ctx context.Context, storyID, labelID uuid.UUID) error {
	return s.repo.UnassignLabel(ctx, storyID, labelID)
}

func (s *StoryLabelService) ListLabels(ctx context.Context, storyID uuid.UUID) ([]model.Label, error) {
	return s.repo.ListLabels(ctx, storyID)
}
