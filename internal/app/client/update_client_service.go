package client

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.UpdateClientUseCase = (*UpdateClientService)(nil)

type UpdateClientService struct {
	repository ports.PersistClientPort
}

func NewUpdateClientService(repository ports.PersistClientPort) *UpdateClientService {
	return &UpdateClientService{repository}
}

func (c *UpdateClientService) Execute(ctx context.Context, cmd *command.UpdateClientCommand) (*entity.Client, error) {
	client := cmd.ToDomain()
	if err := c.repository.Update(ctx, client); err != nil {
		return nil, err
	}
	return client, nil
}
