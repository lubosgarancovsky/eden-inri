package client

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.DeleteClientUseCase = (*DeleteClientService)(nil)

type DeleteClientService struct {
	repository ports.PersistClientPort
}

func NewDeleteClientService(repository ports.PersistClientPort) *DeleteClientService {
	return &DeleteClientService{repository}
}

func (c *DeleteClientService) Execute(ctx context.Context, userID, projectID uuid.UUID) error {
	return c.repository.Delete(ctx, userID, projectID)
}
