package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistInvoiceStatsPort interface {
	Stats(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*entity.InvoiceStats, error)
	MonthlyRevenue(ctx context.Context, userID uuid.UUID) (*[]entity.InvoiceMonthlyRevenue, error)
}

type GetInvoiceStatsUseCase interface {
	Execute(ctx context.Context, query *query.ListQuery) (*entity.InvoiceStats, error)
}

type GetMonthlyRevenueUseCase interface {
	Execute(ctx context.Context, userID uuid.UUID) (*[]entity.InvoiceMonthlyRevenue, error)
}
