package ports

import (
	"context"
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type CreateInvoiceUseCase interface {
	Execute(ctx context.Context, inv *Invoice) (*Invoice, error)
}

type UpdateInvoiceUseCase interface {
	Execute(ctx context.Context, inv *Invoice) (*Invoice, error)
}

type DeleteInvoiceUseCase interface {
	Execute(ctx context.Context, userID, invoiceID uuid.UUID) error
}

type FindInvoiceByIDUseCase interface {
	Execute(ctx context.Context, userID, invoiceID uuid.UUID) (*Invoice, error)
}

type ListInvoicesUseCase interface {
	Execute(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]Invoice, int64, error)
}
