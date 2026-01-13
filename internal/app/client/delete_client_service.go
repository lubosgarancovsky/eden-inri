package client

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

var _ ports.DeleteClientUseCase = (*DeleteClientService)(nil)

type DeleteClientService struct {
	repository ports.PersistClientPort
}

func NewDeleteClientService(repository ports.PersistClientPort) *DeleteClientService {
	return &DeleteClientService{repository}
}

func (c *DeleteClientService) Execute(ctx context.Context, cmd *command.Command) error {
	return c.repository.Delete(ctx, cmd.UserID, cmd.ID)
}
