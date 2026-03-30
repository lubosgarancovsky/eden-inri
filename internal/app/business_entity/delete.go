package business_entity

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.DeleteBusinessEntityUseCase = (*DeleteBusinessEntityService)(nil)

type DeleteBusinessEntityService struct {
	repo ports.BusinessEntityRepository
}

func NewDeleteBusinessEntityService(repo ports.BusinessEntityRepository) *DeleteBusinessEntityService {
	return &DeleteBusinessEntityService{repo: repo}
}

func (s *DeleteBusinessEntityService) Execute(ctx context.Context, cmd *command.DeleteBusinessEntityCommand) (*entity.BusinessEntity, error) {
	be, err := s.repo.FindByID(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Delete(ctx, cmd.UserID, cmd.ID); err != nil {
		return nil, err
	}

	return be, nil
}
