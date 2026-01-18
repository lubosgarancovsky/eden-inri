package client

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.CreateClientUseCase = (*CreateClientService)(nil)

type CreateClientService struct {
	repository ports.PersistClientPort
}

func NewCreateClientService(
	repository ports.PersistClientPort,
) *CreateClientService {
	return &CreateClientService{repository}
}

func (c *CreateClientService) Execute(ctx context.Context, cmd *command.CreateClientCommand) (*entity.Client, error) {
	client := cmd.ToDomain()
	if err := c.repository.Create(ctx, client); err != nil {
		return nil, err
	}
	return client, nil
}
