package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindKanbanBoardByIDService struct {
	repo ports.PersistKanbanBoardPort
}

func NewFindKanbanBoardByIDService(repo ports.PersistKanbanBoardPort) *FindKanbanBoardByIDService {
	return &FindKanbanBoardByIDService{repo: repo}
}

func (s *FindKanbanBoardByIDService) Execute(ctx context.Context, q *query.FindByIDProjectScopedQuery) (*entity.KanbanBoard, error) {
	return s.repo.FindByID(ctx, q.ProjectID, q.ID)
}
