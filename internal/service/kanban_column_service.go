package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
)

type KanbanColumnService struct {
	repo               *repository.KanbanColumnRepository
	projectUserService *ProjectUserService
	kanbanBoardService *KanbanBoardService
}

func NewKanbanColumnService(repo *repository.KanbanColumnRepository, pus *ProjectUserService, kbs *KanbanBoardService) *KanbanColumnService {
	return &KanbanColumnService{repo: repo, projectUserService: pus, kanbanBoardService: kbs}
}

// FindAll columns of a board (caller must be member)
func (s *KanbanColumnService) FindAll(userID, boardID uuid.UUID) ([]model.KanbanColumn, error) {
	projectID, err := s.kanbanBoardService.GetProjectIDByBoardID(boardID)
	if err != nil {
		return nil, err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}
	return s.repo.FindAll(boardID)
}

// FindByID column
func (s *KanbanColumnService) FindByID(userID, boardID, columnID uuid.UUID) (*model.KanbanColumn, error) {
	projectID, err := s.kanbanBoardService.GetProjectIDByBoardID(boardID)
	if err != nil {
		return nil, err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}
	return s.repo.FindByID(columnID)
}

// Insert column
func (s *KanbanColumnService) Insert(userID, boardID uuid.UUID, colReq *model.KanbanColumnRequest) (*model.KanbanColumn, error) {
	projectID, err := s.kanbanBoardService.GetProjectIDByBoardID(boardID)
	if err != nil {
		return nil, err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}

	col := &model.KanbanColumn{
		BoardID:  boardID,
		Name:     colReq.Name,
		Type:     colReq.Type,
		Position: colReq.Position,
	}

	return s.repo.Insert(col)
}

// Update column
func (s *KanbanColumnService) Update(userID, boardID, columnID uuid.UUID, colReq *model.KanbanColumnRequest) (*model.KanbanColumn, error) {
	projectID, err := s.kanbanBoardService.GetProjectIDByBoardID(boardID)
	if err != nil {
		return nil, err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}

	col, err := s.repo.FindByID(columnID)
	if err != nil {
		return nil, err
	}

	col.Name = colReq.Name
	col.Type = colReq.Type
	col.Position = colReq.Position

	return s.repo.Update(col)
}

// Delete column
func (s *KanbanColumnService) Delete(userID, boardID, columnID uuid.UUID) error {
	projectID, err := s.kanbanBoardService.GetProjectIDByBoardID(boardID)
	if err != nil {
		return err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return err
	}
	return s.repo.Delete(columnID)
}
