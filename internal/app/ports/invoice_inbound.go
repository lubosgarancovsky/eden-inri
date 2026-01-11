package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateInvoiceUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateInvoiceCommand) (*entity.Invoice, error)
}

type UpdateInvoiceUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateInvoiceCommand) (*entity.Invoice, error)
}

type DeleteInvoiceUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteCommand) error
}

type FindInvoiceByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindByIDQuery) (*entity.Invoice, error)
}

type ListInvoicesUseCase interface {
	Execute(ctx context.Context, query *query.ListQuery) (*[]entity.Invoice, int64, error)
}
