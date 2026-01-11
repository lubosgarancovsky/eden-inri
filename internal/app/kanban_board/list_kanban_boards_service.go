package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListKanbanBoardsService struct {
	repo ports.PersistKanbanBoardPort
}

func NewListKanbanBoardsService(repo ports.PersistKanbanBoardPort) *ListKanbanBoardsService {
	return &ListKanbanBoardsService{repo: repo}
}

func (s *ListKanbanBoardsService) Execute(ctx context.Context, q *query.ListProjectScopedQuery) (*[]entity.KanbanBoard, int64, error) {
	return s.repo.List(ctx, q.ProjectID, q.ListingQuery)
}
