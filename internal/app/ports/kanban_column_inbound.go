package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateKanbanColumnUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateKanbanColumnCommand) (*entity.KanbanColumn, error)
}

type UpdateKanbanColumnUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateKanbanColumnCommand) (*entity.KanbanColumn, error)
}

type DeleteKanbanColumnUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteKanbanColumnCommand) error
}

type ListKanbanColumnsUseCase interface {
	Execute(ctx context.Context, query *query.ListKanbanColumnsQuery) ([]entity.KanbanColumn, error)
}
