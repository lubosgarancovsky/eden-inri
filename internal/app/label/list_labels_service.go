package label

import (
	"context"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

var _ ports.ListLabelsUseCase = (*ListLabelsService)(nil)

type ListLabelsService struct{ repo ports.QueryLabelPort }

func NewListLabelsService(repo ports.QueryLabelPort) *ListLabelsService {
	return &ListLabelsService{repo}
}

func (s *ListLabelsService) Execute(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]ports.Label, int64, error) {
	return s.repo.List(ctx, projectID, lq)
}
