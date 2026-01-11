package kanban_column

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindKanbanColumnByIDService struct {
	repo ports.PersistKanbanColumnPort
}

func NewFindKanbanColumnByIDService(repo ports.PersistKanbanColumnPort) *FindKanbanColumnByIDService {
	return &FindKanbanColumnByIDService{repo: repo}
}

func (s *FindKanbanColumnByIDService) Execute(ctx context.Context, q *query.FindByIDBoardScopedQuery) (*entity.KanbanColumn, error) {
	return s.repo.FindByID(ctx, q.BoardID, q.ID)
}
