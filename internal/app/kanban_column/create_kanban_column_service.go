package kanban_column

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateKanbanColumnService struct {
	repo ports.PersistKanbanColumnPort
}

func NewCreateKanbanColumnService(repo ports.PersistKanbanColumnPort) *CreateKanbanColumnService {
	return &CreateKanbanColumnService{repo: repo}
}

func (s *CreateKanbanColumnService) Execute(ctx context.Context, cmd *command.CreateKanbanColumnCommand) (*entity.KanbanColumn, error) {
	col := cmd.ToDomain()
	if err := s.repo.Create(ctx, col); err != nil {
		return nil, err
	}
	return col, nil
}
