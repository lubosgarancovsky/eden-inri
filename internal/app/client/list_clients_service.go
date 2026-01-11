package client

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.ListClientsUseCase = (*ListClientsService)(nil)

type ListClientsService struct {
	repository ports.PersistClientPort
}

func NewListClientsService(
	repository ports.PersistClientPort,
) *ListClientsService {
	return &ListClientsService{repository}
}

func (c *ListClientsService) Execute(ctx context.Context, query *query.ListQuery) (*[]entity.Client, int64, error) {
	return c.repository.List(ctx, query.UserID, query.ListingQuery)
}
