package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type TaxRepository struct {
	db *gorm.DB
}

func NewTaxRepository(db *gorm.DB) *TaxRepository {
	return &TaxRepository{db: db}
}

func (r *TaxRepository) List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Tax, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.Tax{}).Where("user_id = ?", userID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}
	items, total, err := ListToDomain[model.Tax, entity.Tax](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *TaxRepository) FindByID(ctx context.Context, userID, taxID uuid.UUID) (*entity.Tax, error) {
	db := GetDB(ctx, r.db)
	var m model.Tax
	if err := db.Model(&model.Tax{}).
		Where("id = ? AND user_id = ?", taxID, userID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("tax with ID: %s was not found", taxID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToDomain(), nil
}

func (r *TaxRepository) Create(ctx context.Context, userID uuid.UUID, tax *entity.Tax) (*entity.Tax, error) {
	db := GetDB(ctx, r.db)
	m := mapper.TaxFromDomain(tax)
	m.UserID = userID
	if err := db.Create(m).Error; err != nil {
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return r.FindByID(ctx, userID, tax.ID)
}

func (r *TaxRepository) Update(ctx context.Context, userID, taxID uuid.UUID, tax *entity.Tax) (*entity.Tax, error) {
	db := GetDB(ctx, r.db)
	result := db.Model(&model.Tax{}).
		Where("id = ? AND user_id = ?", taxID, userID).
		Updates(mapper.TaxFromDomain(tax))

	if result.Error != nil {
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("tax with id %s not found", taxID))
	}

	return r.FindByID(ctx, userID, taxID)
}

func (r *TaxRepository) Delete(ctx context.Context, userID, taxID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("id = ? AND user_id = ?", taxID, userID).
		Delete(&model.Tax{})

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("tax with id %s not found", taxID))
	}

	return nil
}
