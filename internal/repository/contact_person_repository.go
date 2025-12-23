package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
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

func (r *ContactPersonRepository) FindAll(userID, clientId uuid.UUID, lq *list.ListingQuery) ([]model.ContactPerson, int64, error) {
	query := r.db.Model(&model.ContactPerson{}).Where("user_id = ? AND client_id", userID, clientId)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.ContactPerson](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *ContactPersonRepository) FindByID(clientId, id uuid.UUID) (*model.ContactPerson, error) {
	var result model.ContactPerson
	if err := r.db.Model(&model.ContactPerson{}).Where("id = ? AND client_id", id, clientId).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ContactPersonRepository) Insert(cp *model.ContactPerson) (*model.ContactPerson, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(cp).Error; err != nil {
		return nil, err
	}
	return cp, nil
}

func (r *ContactPersonRepository) Update(cp *model.ContactPerson) (*model.ContactPerson, error) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ? AND client_id", cp.ID, cp.ClientID).Updates(&cp)
	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Contact person with id %s does not exist", cp.ID))
	}
	return cp, nil
}

func (r *ContactPersonRepository) Delete(clientId, id uuid.UUID) error {
	result := r.db.Clauses(clause.Returning{}).Where("id = ? AND client_id", id, clientId).Delete(model.ContactPerson{})
	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Contact person with id %s does not exist", id))
	}
	return nil
}
