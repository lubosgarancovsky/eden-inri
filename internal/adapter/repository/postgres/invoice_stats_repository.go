package postgres

import (
	"context"

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
