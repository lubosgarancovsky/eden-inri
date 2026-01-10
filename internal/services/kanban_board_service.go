package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/go-kit/list"
)

type KanbanBoardService struct {
	repo               *repositories.KanbanBoardRepository
	projectUserService *ProjectUserService
}

func NewKanbanBoardService(repo *repositories.KanbanBoardRepository, projectUserService *ProjectUserService) *KanbanBoardService {
	return &KanbanBoardService{repo: repo, projectUserService: projectUserService}
}

// List all boards in project (caller must be member)
func (s *KanbanBoardService) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[models.KanbanBoard], error) {
	items, totalCount, err := s.repo.FindAll(ctx, projectID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[models.KanbanBoard]{
		Items:      *items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

// Get board detail
func (s *KanbanBoardService) FindByID(ctx context.Context, projectID, boardID uuid.UUID) (*models.KanbanBoard, error) {
	return s.repo.FindByID(ctx, boardID, projectID)
}

// Create board (requires user to be member of project)
func (s *KanbanBoardService) Insert(ctx context.Context, projectID uuid.UUID, input *models.KanbanBoardRequest) (*models.KanbanBoard, error) {
	board := &models.KanbanBoard{
		ProjectID: projectID,
		Name:      input.Name,
		Status:    input.Status,
	}

	return s.repo.Create(ctx, board)
}

// Create board (requires user to be member of project)
func (s *KanbanBoardService) Update(ctx context.Context, projectID uuid.UUID, kanbanID uuid.UUID, input *models.KanbanBoardRequest) (*models.KanbanBoard, error) {
	board := &models.KanbanBoard{
		ID:        kanbanID,
		ProjectID: projectID,
		Name:      input.Name,
		Status:    input.Status,
	}

	return s.repo.Update(ctx, board)
}

// Delete board (caller must be member)
func (s *KanbanBoardService) Delete(ctx context.Context, projectID, boardID uuid.UUID) error {
	return s.repo.Delete(ctx, boardID, projectID)
}

func (s *KanbanBoardService) GetProjectIDByBoardID(ctx context.Context, boardID uuid.UUID) (uuid.UUID, error) {
	return s.repo.GetProjectIDByBoardID(ctx, boardID)
}
