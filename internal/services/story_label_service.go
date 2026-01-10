package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
)

type StoryLabelService struct {
	repo               *repositories.StoryLabelRepository
	projectUserService *ProjectUserService
	kanbanService      *KanbanBoardService
}

func NewStoryLabelService(repo *repositories.StoryLabelRepository, pus *ProjectUserService, ks *KanbanBoardService) *StoryLabelService {
	return &StoryLabelService{repo: repo, projectUserService: pus, kanbanService: ks}
}

func (s *StoryLabelService) AssignLabel(ctx context.Context, projectID, storyID uuid.UUID, input *models.StoryLabelRequest) error {
	storyLabel := &models.StoryLabel{
		StoryID: storyID,
		LabelID: input.LabelID,
	}
	return s.repo.AssignLabel(ctx, projectID, storyLabel)
}

func (s *StoryLabelService) UnassignLabel(ctx context.Context, storyID, labelID uuid.UUID) error {
	return s.repo.UnassignLabel(ctx, storyID, labelID)
}

func (s *StoryLabelService) ListLabels(ctx context.Context, storyID uuid.UUID) ([]models.Label, error) {
	return s.repo.ListLabels(ctx, storyID)
}
