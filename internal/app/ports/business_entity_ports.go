package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type BusinessEntityRepository interface {
	List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.BusinessEntity, int64, error)
	FindByID(ctx context.Context, userID, businessEntityID uuid.UUID) (*entity.BusinessEntity, error)
	Create(ctx context.Context, userID uuid.UUID, be *entity.BusinessEntity) (*entity.BusinessEntity, error)
	Update(ctx context.Context, userID, businessEntityID uuid.UUID, be *entity.BusinessEntity) (*entity.BusinessEntity, error)
	Delete(ctx context.Context, userID, businessEntityID uuid.UUID) error
}

type FindBusinessEntityByIDUseCase interface {
	Execute(ctx context.Context, query *query.Query) (*entity.BusinessEntity, error)
}

type ListBusinessEntitiesUseCase interface {
	Execute(ctx context.Context, query *query.ListQuery) (*[]entity.BusinessEntity, int64, error)
}

type CreateBusinessEntityUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateBusinessEntityCommand) (*entity.BusinessEntity, error)
}

type UpdateBusinessEntityUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateBusinessEntityCommand) (*entity.BusinessEntity, error)
}

type DeleteBusinessEntityUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteBusinessEntityCommand) (*entity.BusinessEntity, error)
}
