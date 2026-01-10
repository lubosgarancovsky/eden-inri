package ports

import (
	"context"
	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
	"time"
)

type Invoice struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	ClientID    uuid.UUID
	Number      string
	IssuedAt    time.Time
	DueAt       time.Time
	TotalAmount int64
	Currency    string
	Status      string
	Note        string
}

type PersistInvoicePort interface {
	Create(ctx context.Context, inv *Invoice) error
	Update(ctx context.Context, inv *Invoice) error
	Delete(ctx context.Context, userID, invoiceID uuid.UUID) error
}

type QueryInvoicePort interface {
	FindByID(ctx context.Context, userID, invoiceID uuid.UUID) (*Invoice, error)
	List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]Invoice, int64, error)
}
