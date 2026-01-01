package repository

import (
	"context"
	"fmt"

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
		Model(model.Invoice{}).
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
