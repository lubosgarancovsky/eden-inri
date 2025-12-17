package service

import (
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

// Assign a label to a story
func (s *StoryLabelService) AssignLabel(userID, storyID, kanbanID, labelID uuid.UUID) error {
	projectID, err := s.kanbanService.GetProjectIDByBoardID(kanbanID)
	if err != nil {
		return err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return err
	}
	return s.repo.AssignLabel(storyID, labelID)
}

// Unassign a label from a story
func (s *StoryLabelService) UnassignLabel(userID, storyID, kanbanID, labelID uuid.UUID) error {
	projectID, err := s.kanbanService.GetProjectIDByBoardID(kanbanID)
	if err != nil {
		return err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return err
	}
	return s.repo.UnassignLabel(storyID, labelID)
}

// ListLabels for a story
func (s *StoryLabelService) ListLabels(userID, storyID, kanbanID uuid.UUID) ([]model.Label, error) {
	projectID, err := s.kanbanService.GetProjectIDByBoardID(kanbanID)
	if err != nil {
		return nil, err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}
	return s.repo.ListLabels(storyID)
}
