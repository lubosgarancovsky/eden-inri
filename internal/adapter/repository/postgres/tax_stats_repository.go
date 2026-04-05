package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type TaxStatsRepository struct {
	db *gorm.DB
}

func NewTaxStatsRepository(db *gorm.DB) *TaxStatsRepository {
	return &TaxStatsRepository{db: db}
}

func (r *TaxStatsRepository) Stats(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*entity.TaxStats, error) {
	db := GetDB(ctx, r.db)

	query := db.
		Model(&model.Tax{}).
		Select("COUNT(*) as count, SUM(amount) as total").
		Where("user_id = ?", userID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	var stats entity.TaxStats
	err := query.Scan(&stats).Error
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
