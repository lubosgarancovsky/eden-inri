package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type ContactPersonRepository struct {
	db *gorm.DB
}

func NewContactPersonRepository(db *gorm.DB) *ContactPersonRepository {
	return &ContactPersonRepository{db: db}
}

func (r *ContactPersonRepository) List(ctx context.Context, userID, clientID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ContactPerson, int64, error) {
	db := GetDB(ctx, r.db)

	query := db.Model(&model.ContactPerson{}).Where("user_id = ? AND client_id = ?", userID, clientID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.ContactPerson, entity.ContactPerson](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *ContactPersonRepository) FindByID(ctx context.Context, userID, clientID, contactPersonID uuid.UUID) (*entity.ContactPerson, error) {
	db := GetDB(ctx, r.db)

	var result model.ContactPerson
	if err := db.Where("user_id = ? AND client_id = ? AND id = ?", userID, clientID, contactPersonID).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("Contact person with id %s not found", contactPersonID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return result.ToDomain(), nil
}

func (r *ContactPersonRepository) Create(ctx context.Context, cp *entity.ContactPerson) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.ContactPersonFromDomain(cp)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *ContactPersonRepository) Update(ctx context.Context, cp *entity.ContactPerson) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("user_id = ? AND client_id = ? AND id = ?", cp.UserID, cp.ClientID, cp.ID).
		Updates(mapper.ContactPersonFromDomain(cp))

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("Contact person with id %s not found", cp.ID))
	}

	return nil
}

func (r *ContactPersonRepository) Delete(ctx context.Context, userID, clientID, contactPersonID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("user_id = ? AND client_id = ? AND id = ?", userID, clientID, contactPersonID).
		Delete(&model.ContactPerson{})

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("Contact person with id %s not found", contactPersonID))
	}

	return nil
}
