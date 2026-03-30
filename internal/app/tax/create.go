package tax

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.CreateTaxUseCase = (*CreateTaxService)(nil)

type CreateTaxService struct {
	repo ports.TaxRepository
}

func NewCreateTaxService(repo ports.TaxRepository) *CreateTaxService {
	return &CreateTaxService{repo: repo}
}

func (s *CreateTaxService) Execute(ctx context.Context, cmd *command.CreateTaxCommand) (*entity.Tax, error) {
	newTax := cmd.ToDomain()

	return s.repo.Create(ctx, cmd.UserID, newTax)
}
