package tax

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.DeleteTaxUseCase = (*DeleteTaxService)(nil)

type DeleteTaxService struct {
	repo ports.TaxRepository
}

func NewDeleteTaxService(repo ports.TaxRepository) *DeleteTaxService {
	return &DeleteTaxService{repo: repo}
}

func (s *DeleteTaxService) Execute(ctx context.Context, cmd *command.DeleteTaxCommand) (*entity.Tax, error) {
	tax, err := s.repo.FindByID(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Delete(ctx, cmd.UserID, cmd.ID); err != nil {
		return nil, err
	}

	return tax, nil
}
