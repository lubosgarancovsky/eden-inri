package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type InvoiceStatsRepository struct {
	db *gorm.DB
}

func NewInvoiceStatsRepository(db *gorm.DB) *InvoiceStatsRepository {
	return &InvoiceStatsRepository{db: db}
}

func (r *InvoiceStatsRepository) Stats(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*entity.InvoiceStats, error) {
	db := GetDB(ctx, r.db)

	query := db.
		Model(&model.Invoice{}).
		Select("COUNT(*) as count, SUM(total) as total, SUM(billable_hours) as billable_hours").
		Where("user_id = ?", userID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	var stats entity.InvoiceStats
	err := query.Scan(&stats).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *InvoiceStatsRepository) MonthlyRevenue(ctx context.Context, userID uuid.UUID) (*[]entity.InvoiceMonthlyRevenue, error) {
	type MonthlyRevenue struct {
		Month   time.Time
		Revenue float64
	}

	db := GetDB(ctx, r.db)

	var items []MonthlyRevenue

	err := db.
		Table("inri_invoice").
		Select("DATE_TRUNC('month', paid_at) AS month, SUM(total) AS revenue").
		Where("is_canceled = false").
		Where("paid_at IS NOT NULL").
		Where("paid_at >= NOW() - INTERVAL '12 months'").
		Where("user_id = ?", userID).
		Group("month").
		Order("month").
		Scan(&items).Error

	if err != nil {
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	result := make([]entity.InvoiceMonthlyRevenue, len(items))
	for i := range items {
		result[i] = entity.InvoiceMonthlyRevenue{Month: items[i].Month, Revenue: items[i].Revenue}
	}

	return &result, nil
}
