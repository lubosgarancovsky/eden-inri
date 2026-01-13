package client

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

var _ ports.FindClientByIDUseCase = (*FindClientByIDService)(nil)

type FindClientByIDService struct {
	repository ports.PersistClientPort
}

func NewFindClientByIDService(
	repository ports.PersistClientPort,
) *FindClientByIDService {
	return &FindClientByIDService{repository}
}

func (c *FindClientByIDService) Execute(ctx context.Context, q *query.Query) (*entity.Client, error) {
	return c.repository.FindByID(ctx, q.UserID, q.ID)
}
