package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateClientUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateClientCommand) (*entity.Client, error)
}

type UpdateClientUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateClientCommand) (*entity.Client, error)
}

type DeleteClientUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteClientCommand) error
}

type FindClientByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindClientByIDQuery) (*entity.Client, error)
}

type ListClientsUseCase interface {
	Execute(ctx context.Context, query *query.ListClientsQuery) (*[]entity.Client, int64, error)
}
