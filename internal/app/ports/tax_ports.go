package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type TaxRepository interface {
	List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Tax, int64, error)
	FindByID(ctx context.Context, userID, taxID uuid.UUID) (*entity.Tax, error)
	Create(ctx context.Context, userID uuid.UUID, tax *entity.Tax) (*entity.Tax, error)
	Update(ctx context.Context, userID, taxID uuid.UUID, tax *entity.Tax) (*entity.Tax, error)
	Delete(ctx context.Context, userID, taxID uuid.UUID) error
}

type FindTaxByIDUseCase interface {
	Execute(ctx context.Context, query *query.Query) (*entity.Tax, error)
}

type ListTaxesUseCase interface {
	Execute(ctx context.Context, query *query.ListQuery) (*[]entity.Tax, int64, error)
}

type CreateTaxUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateTaxCommand) (*entity.Tax, error)
}

type UpdateTaxUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateTaxCommand) (*entity.Tax, error)
}

type DeleteTaxUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteTaxCommand) (*entity.Tax, error)
}
