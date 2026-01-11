package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistInvoicePort interface {
	Create(ctx context.Context, inv *entity.Invoice) error
	Update(ctx context.Context, inv *entity.Invoice) error
	Delete(ctx context.Context, userID, invoiceID uuid.UUID) error
	FindByID(ctx context.Context, userID, invoiceID uuid.UUID) (*entity.Invoice, error)
	List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Invoice, int64, error)
}
