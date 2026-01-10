package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateClientUseCase interface {
	Execute(ctx context.Context, client *command.CreateClientCommand) (*entity.Client, error)
}

type UpdateClientUseCase interface {
	Execute(ctx context.Context, client *entity.Client) (*entity.Client, error)
}

type DeleteClientUseCase interface {
	Execute(ctx context.Context, userID, clientID uuid.UUID) error
}
