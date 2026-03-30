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

type BusinessEntityRepository struct {
	db *gorm.DB
}

func NewBusinessEntityRepository(db *gorm.DB) *BusinessEntityRepository {
	return &BusinessEntityRepository{db: db}
}

func (r *BusinessEntityRepository) List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.BusinessEntity, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.BusinessEntity{}).Where("user_id = ?", userID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}
	items, total, err := ListToDomain[model.BusinessEntity, entity.BusinessEntity](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *BusinessEntityRepository) FindByID(ctx context.Context, userID, businessEntityID uuid.UUID) (*entity.BusinessEntity, error) {
	db := GetDB(ctx, r.db)
	var m model.BusinessEntity
	if err := db.Model(&model.BusinessEntity{}).
		Where("id = ? AND user_id = ?", businessEntityID, userID).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("business entity with ID: %s was not found", businessEntityID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToDomain(), nil
}

func (r *BusinessEntityRepository) Create(ctx context.Context, userID uuid.UUID, be *entity.BusinessEntity) (*entity.BusinessEntity, error) {
	db := GetDB(ctx, r.db)
	m := mapper.BusinessEntityFromDomain(be)
	m.UserID = userID
	if err := db.Create(m).Error; err != nil {
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return r.FindByID(ctx, userID, be.ID)
}

func (r *BusinessEntityRepository) Update(ctx context.Context, userID, businessEntityID uuid.UUID, be *entity.BusinessEntity) (*entity.BusinessEntity, error) {
	db := GetDB(ctx, r.db)
	result := db.Model(&model.BusinessEntity{}).
		Where("id = ? AND user_id = ?", businessEntityID, userID).
		Updates(mapper.BusinessEntityFromDomain(be))

	if result.Error != nil {
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("business entity with id %s not found", businessEntityID))
	}

	return r.FindByID(ctx, userID, businessEntityID)
}

func (r *BusinessEntityRepository) Delete(ctx context.Context, userID, businessEntityID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("id = ? AND user_id = ?", businessEntityID, userID).
		Delete(&model.BusinessEntity{})

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("business entity with id %s not found", businessEntityID))
	}

	return nil
}
