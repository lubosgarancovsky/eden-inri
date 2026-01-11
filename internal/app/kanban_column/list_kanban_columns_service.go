package kanban_column

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type ListKanbanColumnsService struct {
	repo ports.PersistKanbanColumnPort
}

func NewListKanbanColumnsService(repo ports.PersistKanbanColumnPort) *ListKanbanColumnsService {
	return &ListKanbanColumnsService{repo: repo}
}

func (s *ListKanbanColumnsService) Execute(ctx context.Context, boardID uuid.UUID) ([]entity.KanbanColumn, error) {
	return s.repo.List(ctx, boardID)
}
