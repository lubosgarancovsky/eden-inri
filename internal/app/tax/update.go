package tax

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.UpdateTaxUseCase = (*UpdateTaxService)(nil)

type UpdateTaxService struct {
	repo ports.TaxRepository
}

func NewUpdateTaxService(repo ports.TaxRepository) *UpdateTaxService {
	return &UpdateTaxService{repo: repo}
}

func (s *UpdateTaxService) Execute(ctx context.Context, cmd *command.UpdateTaxCommand) (*entity.Tax, error) {
	tax, err := s.repo.FindByID(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(tax)

	return s.repo.Update(ctx, cmd.UserID, cmd.ID, tax)
}
