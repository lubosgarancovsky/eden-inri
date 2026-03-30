package business_entity

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.CreateBusinessEntityUseCase = (*CreateBusinessEntityService)(nil)

type CreateBusinessEntityService struct {
	repo ports.BusinessEntityRepository
}

func NewCreateBusinessEntityService(repo ports.BusinessEntityRepository) *CreateBusinessEntityService {
	return &CreateBusinessEntityService{repo: repo}
}

func (s *CreateBusinessEntityService) Execute(ctx context.Context, cmd *command.CreateBusinessEntityCommand) (*entity.BusinessEntity, error) {
	newBE := cmd.ToDomain()

	return s.repo.Create(ctx, cmd.UserID, newBE)
}
