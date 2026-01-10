package label

import (
	"context"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.DeleteLabelUseCase = (*DeleteLabelService)(nil)

type DeleteLabelService struct {
	repo ports.PersistLabelPort
	tm   ports.TransactionManager
}

func NewDeleteLabelService(repo ports.PersistLabelPort, tm ports.TransactionManager) *DeleteLabelService {
	return &DeleteLabelService{repo: repo, tm: tm}
}

func (s *DeleteLabelService) Execute(ctx context.Context, projectID, labelID uuid.UUID) error {
	return s.tm.WithTransaction(ctx, func(ctx context.Context) error {
		return s.repo.Delete(ctx, projectID, labelID)
	})
}
