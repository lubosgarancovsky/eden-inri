package business_entity

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.UpdateBusinessEntityUseCase = (*UpdateBusinessEntityService)(nil)

type UpdateBusinessEntityService struct {
	repo ports.BusinessEntityRepository
}

func NewUpdateBusinessEntityService(repo ports.BusinessEntityRepository) *UpdateBusinessEntityService {
	return &UpdateBusinessEntityService{repo: repo}
}

func (s *UpdateBusinessEntityService) Execute(ctx context.Context, cmd *command.UpdateBusinessEntityCommand) (*entity.BusinessEntity, error) {
	be, err := s.repo.FindByID(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(be)

	return s.repo.Update(ctx, cmd.UserID, cmd.ID, be)
}
