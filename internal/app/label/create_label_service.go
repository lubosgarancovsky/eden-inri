package label

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.CreateLabelUseCase = (*CreateLabelService)(nil)

type CreateLabelService struct {
	repo ports.PersistLabelPort
	tm   ports.TransactionManager
}

func NewCreateLabelService(repo ports.PersistLabelPort, tm ports.TransactionManager) *CreateLabelService {
	return &CreateLabelService{repo: repo, tm: tm}
}

func (s *CreateLabelService) Execute(ctx context.Context, label *ports.Label) (*ports.Label, error) {
	if err := s.tm.WithTransaction(ctx, func(ctx context.Context) error {
		return s.repo.Create(ctx, label)
	}); err != nil {
		return nil, err
	}
	return label, nil
}
