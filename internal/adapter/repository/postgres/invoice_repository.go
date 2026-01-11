package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

var _ ports.QueryInvoicePort = (*InvoiceRepository)(nil)
var _ ports.PersistInvoicePort = (*InvoiceRepository)(nil)

type InvoiceRepository struct{ db *gorm.DB }

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository { return &InvoiceRepository{db: db} }

func (r *InvoiceRepository) List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Invoice, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.Invoice{}).Where("user_id = ?", userID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}
	items, total, err := ListToDomain[model.Invoice, entity.Invoice](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *InvoiceRepository) FindByID(ctx context.Context, userID, invoiceID uuid.UUID) (*entity.Invoice, error) {
	db := GetDB(ctx, r.db)
	var m model.Invoice
	if err := db.Model(&model.Invoice{}).
		Where("id = ? AND user_id = ?", invoiceID, userID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("invoice %s not found", invoiceID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToDomain(), nil
}

func (r *InvoiceRepository) Create(ctx context.Context, inv *entity.Invoice) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.InvoiceFromDomain(inv)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *InvoiceRepository) Update(ctx context.Context, inv *entity.Invoice) error {
	db := GetDB(ctx, r.db)
	res := db.Model(&model.Invoice{}).Where("id = ? AND user_id = ?", inv.ID, inv.UserID).Updates(mapper.InvoiceFromDomain(inv))
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("invoice %s not found", inv.ID))
	}
	return nil
}

func (r *InvoiceRepository) Delete(ctx context.Context, userID, invoiceID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ? AND user_id = ?", invoiceID, userID).Delete(&model.Invoice{})
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("invoice %s not found", invoiceID))
	}
	return nil
}
