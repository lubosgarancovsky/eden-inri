package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
)

type KanbanBoardService struct {
	repo               *repository.KanbanBoardRepository
	projectUserService *ProjectUserService
}

func NewKanbanBoardService(repo *repository.KanbanBoardRepository, projectUserService *ProjectUserService) *KanbanBoardService {
	return &KanbanBoardService{repo: repo, projectUserService: projectUserService}
}

// List all boards in project (caller must be member)
func (s *KanbanBoardService) FindAll(userID, projectID uuid.UUID) ([]model.KanbanBoard, error) {
	_, err := s.projectUserService.GetProjectUserIfMember(projectID, userID)
	if err != nil {
		return nil, err
	}

	return s.repo.FindAll(projectID)
}

// Get board detail
func (s *KanbanBoardService) FindByID(userID, projectID, boardID uuid.UUID) (*model.KanbanBoard, error) {
	_, err := s.projectUserService.GetProjectUserIfMember(projectID, userID)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(boardID, projectID)
}

// Create board (requires user to be member of project)
func (s *KanbanBoardService) Insert(userID, projectID uuid.UUID) (*model.KanbanBoard, error) {
	_, err := s.projectUserService.GetProjectUserIfMember(projectID, userID)
	if err != nil {
		return nil, err
	}

	board := &model.KanbanBoard{
		ProjectID: projectID,
	}

	return s.repo.Create(board)
}

// Delete board (caller must be member)
func (s *KanbanBoardService) Delete(userID, projectID, boardID uuid.UUID) error {
	_, err := s.projectUserService.GetProjectUserIfMember(projectID, userID)
	if err != nil {
		return err
	}

	return s.repo.Delete(boardID, projectID)
}
