package kanban_column

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type UpdateKanbanColumnService struct {
	repo ports.PersistKanbanColumnPort
}

func NewUpdateKanbanColumnService(repo ports.PersistKanbanColumnPort) *UpdateKanbanColumnService {
	return &UpdateKanbanColumnService{repo: repo}
}

func (s *UpdateKanbanColumnService) Execute(ctx context.Context, cmd *command.UpdateKanbanColumnCommand) (*entity.KanbanColumn, error) {
	col, err := s.repo.FindByID(ctx, cmd.BoardID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(col)

	if err := s.repo.Update(ctx, col); err != nil {
		return nil, err
	}

	return col, nil
}
