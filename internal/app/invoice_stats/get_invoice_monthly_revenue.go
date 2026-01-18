package invoice_stats

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type GetInvoiceMonthlyRevenueService struct {
	repo ports.PersistInvoiceStatsPort
}

func NewGetInvoiceMonthlyRevenueService(repo ports.PersistInvoiceStatsPort) *GetInvoiceMonthlyRevenueService {
	return &GetInvoiceMonthlyRevenueService{repo: repo}
}

func (s *GetInvoiceMonthlyRevenueService) Execute(ctx context.Context, userID uuid.UUID) (*[]entity.InvoiceMonthlyRevenue, error) {
	return s.repo.MonthlyRevenue(ctx, userID)
}
