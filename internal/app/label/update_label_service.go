package label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.UpdateLabelUseCase = (*UpdateLabelService)(nil)

type UpdateLabelService struct {
	repo ports.PersistLabelPort
	tm   ports.TransactionManager
}

func NewUpdateLabelService(repo ports.PersistLabelPort, tm ports.TransactionManager) *UpdateLabelService {
	return &UpdateLabelService{repo: repo, tm: tm}
}

func (s *UpdateLabelService) Execute(ctx context.Context, label *ports.Label) (*ports.Label, error) {
	if err := s.tm.WithTransaction(ctx, func(ctx context.Context) error {
		return s.repo.Update(ctx, label)
	}); err != nil {
		return nil, err
	}
	return label, nil
}
