package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) FindAll(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]model.Invoice, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(&model.Invoice{}).
		Preload("Client").
		Where("user_id = ?", userID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.Invoice](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *InvoiceRepository) FindByID(ctx context.Context, userID, invoiceID uuid.UUID) (*model.Invoice, error) {
	var result model.Invoice
	if err := r.db.
		WithContext(ctx).
		Model(model.Invoice{}).
		Preload("Client").
		Where("user_id = ? AND id = ?", userID, invoiceID).
		First(&result).
		Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *InvoiceRepository) Insert(ctx context.Context, inv *model.Invoice) (*model.Invoice, error) {
	if err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(inv).
		Error; err != nil {
		return nil, err
	}
	return inv, nil
}

func (r *InvoiceRepository) Update(ctx context.Context, inv *model.Invoice) (*model.Invoice, error) {
	result := r.db.
		WithContext(ctx).
		Model(&model.Invoice{}).
		Clauses(clause.Returning{}).
		Where("user_id = ? AND id = ?", inv.UserID, inv.ID).
		Updates(inv)

	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).
			WithMessage(fmt.Sprintf("Invoice with id %s does not exist", inv.ID))
	}

	return inv, nil
}

func (r *InvoiceRepository) Delete(ctx context.Context, userID, invoiceID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("user_id = ? AND id = ?", userID, invoiceID).
		Delete(model.Invoice{})

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Invoice with id %s does not exist", invoiceID))
	}
	return nil
}

// TotalRevenue returns the sum of totals for all invoices that are paid and not canceled for the given user
func (r *InvoiceRepository) TotalRevenue(ctx context.Context, userID uuid.UUID) (float64, error) {
	var total float64
	err := r.db.
		WithContext(ctx).
		Model(&model.Invoice{}).
		Where("user_id = ? AND paid_at IS NOT NULL AND is_canceled = false", userID).
		Select("COALESCE(SUM(total), 0)").
		Scan(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

// MonthlyRevenue aggregates potential and actual revenue per month for the given user
// - potential: sum(total) of non-canceled invoices
// - actual: sum(total) of non-canceled AND paid invoices
// Grouped by issued_at month in UTC in the range [from, to)
func (r *InvoiceRepository) MonthlyRevenue(ctx context.Context, userID uuid.UUID, from, to time.Time) ([]model.RevenueGraphPoint, error) {
	// Note: using to_char(date_trunc('month', issued_at), 'YYYY-MM') for key
	type row struct {
		Month     string  `gorm:"column:month"`
		Potential float64 `gorm:"column:potential"`
		Actual    float64 `gorm:"column:actual"`
	}

	var rows []row
	// Build the query explicitly to control select and grouping
	err := r.db.WithContext(ctx).
		Table((&model.Invoice{}).TableName()).
		Where("user_id = ? AND issued_at >= ? AND issued_at < ?", userID, from, to).
		Select("to_char(date_trunc('month', issued_at AT TIME ZONE 'UTC'), 'YYYY-MM') as month, " +
			"COALESCE(SUM(CASE WHEN is_canceled = false THEN total ELSE 0 END), 0) as potential, " +
			"COALESCE(SUM(CASE WHEN is_canceled = false AND paid_at IS NOT NULL THEN total ELSE 0 END), 0) as actual").
		Group("month").
		Order("month").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	// Map to model points
	points := make([]model.RevenueGraphPoint, 0, len(rows))
	for _, r := range rows {
		points = append(points, model.RevenueGraphPoint{
			Month:     r.Month,
			Actual:    r.Actual,
			Potential: r.Potential,
		})
	}
	return points, nil
}
