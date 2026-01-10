package client

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.FindClientByIDUseCase = (*FindClientByIDService)(nil)

type FindClientByIDService struct {
	repository ports.QueryClientPort
}

func NewFindClientByIDService(
	repository ports.QueryClientPort,
) *FindClientByIDService {
	return &FindClientByIDService{repository}
}

func (c *FindClientByIDService) Execute(ctx context.Context, q *query.FindClientByIDQuery) (*entity.Client, error) {
	return c.repository.FindByID(ctx, q.UserID, q.ID)
}
