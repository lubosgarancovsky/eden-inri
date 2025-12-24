package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
)

type KanbanBoardService struct {
	repo               *repository.KanbanBoardRepository
	projectUserService *ProjectUserService
}

func NewKanbanBoardService(repo *repository.KanbanBoardRepository, projectUserService *ProjectUserService) *KanbanBoardService {
	return &KanbanBoardService{repo: repo, projectUserService: projectUserService}
}

// List all boards in project (caller must be member)
func (s *KanbanBoardService) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.KanbanBoard], error) {
	items, totalCount, err := s.repo.FindAll(ctx, projectID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.KanbanBoard]{
		Items:      *items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

// Get board detail
func (s *KanbanBoardService) FindByID(ctx context.Context, projectID, boardID uuid.UUID) (*model.KanbanBoard, error) {
	return s.repo.FindByID(ctx, boardID, projectID)
}

// Create board (requires user to be member of project)
func (s *KanbanBoardService) Insert(ctx context.Context, projectID uuid.UUID, input *model.KanbanBoardRequest) (*model.KanbanBoard, error) {
	board := &model.KanbanBoard{
		ProjectID: projectID,
		Name:      input.Name,
		Status:    input.Status,
	}

	return s.repo.Create(ctx, board)
}

// Create board (requires user to be member of project)
func (s *KanbanBoardService) Update(ctx context.Context, projectID uuid.UUID, kanbanID uuid.UUID, input *model.KanbanBoardRequest) (*model.KanbanBoard, error) {
	board := &model.KanbanBoard{
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
