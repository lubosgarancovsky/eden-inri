package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContactPersonRepository struct {
	db *gorm.DB
}

func NewContactPersonRepository(db *gorm.DB) *ContactPersonRepository {
	return &ContactPersonRepository{db: db}
}

func (r *ContactPersonRepository) FindAll(ctx context.Context, userID, clientID uuid.UUID, lq *list.ListingQuery) ([]models.ContactPerson, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(&models.ContactPerson{}).
		Where("user_id = ? AND client_id = ?", userID, clientID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[models.ContactPerson](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *ContactPersonRepository) FindByID(ctx context.Context, userID, clientID, contactPersonID uuid.UUID) (*models.ContactPerson, error) {
	var result models.ContactPerson
	if err := r.db.
		WithContext(ctx).
		Model(&models.ContactPerson{}).
		Where("user_id = ? AND client_id = ? AND id = ?", userID, clientID, contactPersonID).
		First(&result).
		Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ContactPersonRepository) Insert(ctx context.Context, cp *models.ContactPerson) (*models.ContactPerson, error) {
	if err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(cp).
		Error; err != nil {
		return nil, err
	}
	return cp, nil
}

func (r *ContactPersonRepository) Update(ctx context.Context, cp *models.ContactPerson) (*models.ContactPerson, error) {
	result := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("user_id = ? AND client_id = ? AND id = ?", cp.UserID, cp.ClientID, cp.ID).
		Updates(&cp)

	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Contact person with id %s does not exist", cp.ID))
	}
	return cp, nil
}

func (r *ContactPersonRepository) Delete(ctx context.Context, userID, clientID, contactPersonID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("user_id = ? AND client_id = ? AND id = ?", userID, clientID, contactPersonID).
		Delete(models.ContactPerson{})

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Contact person with id %s does not exist", contactPersonID))
	}
	return nil
}
