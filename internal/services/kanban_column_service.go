package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
)

type KanbanColumnService struct {
	repo               *repositories.KanbanColumnRepository
	projectUserService *ProjectUserService
	kanbanBoardService *KanbanBoardService
}

func NewKanbanColumnService(repo *repositories.KanbanColumnRepository, pus *ProjectUserService, kbs *KanbanBoardService) *KanbanColumnService {
	return &KanbanColumnService{repo: repo, projectUserService: pus, kanbanBoardService: kbs}
}

// FindAll columns of a board (caller must be member)
func (s *KanbanColumnService) FindAll(ctx context.Context, boardID uuid.UUID) ([]models.KanbanColumn, error) {
	return s.repo.FindAll(ctx, boardID)
}

// FindByID column
func (s *KanbanColumnService) FindByID(ctx context.Context, boardID, columnID uuid.UUID) (*models.KanbanColumn, error) {
	return s.repo.FindByID(ctx, boardID, columnID)
}

// Insert column
func (s *KanbanColumnService) Insert(ctx context.Context, boardID uuid.UUID, colReq *models.KanbanColumnRequest) (*models.KanbanColumn, error) {
	col := models.KanbanColumn{
		BoardID:  boardID,
		Key:      colReq.Key,
		Name:     colReq.Name,
		Type:     colReq.Type,
		Color:    colReq.Color,
		Position: colReq.Position,
	}

	return s.repo.Insert(ctx, &col)
}

// Update column
func (s *KanbanColumnService) Update(ctx context.Context, boardID, columnID uuid.UUID, colReq *models.KanbanColumnRequest) (*models.KanbanColumn, error) {
	col := models.KanbanColumn{
		ID:       columnID,
		BoardID:  boardID,
		Key:      colReq.Key,
		Name:     colReq.Name,
		Type:     colReq.Type,
		Color:    colReq.Color,
		Position: colReq.Position,
	}

	return s.repo.Update(ctx, &col)
}

// Delete column
func (s *KanbanColumnService) Delete(ctx context.Context, boardID, columnID uuid.UUID) error {
	return s.repo.Delete(ctx, boardID, columnID)
}

func (s *KanbanColumnService) GetBoardID(ctx context.Context, columnID uuid.UUID) (*uuid.UUID, error) {
	return s.repo.GetBoardID(ctx, columnID)
}
