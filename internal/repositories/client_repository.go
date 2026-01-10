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

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) FindAll(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]models.ClientListItem, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(&models.ClientListItem{}).
		Select("*").
		Where("user_id = ?", userID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[models.ClientListItem](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *ClientRepository) FindByID(ctx context.Context, userID, clientID uuid.UUID) (*models.Client, error) {
	var result models.Client
	if err := r.db.
		WithContext(ctx).
		Model(&models.Client{}).
		Select("*").
		Where("user_id = ? AND id = ?", userID, clientID).
		First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ClientRepository) Insert(ctx context.Context, client *models.Client) (*models.Client, error) {
	if err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Select("*").
		Create(client).
		Error; err != nil {
		return nil, err
	}

	return client, nil
}

func (r *ClientRepository) Update(ctx context.Context, client *models.Client) (*models.Client, error) {
	result := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("user_id = ? AND id = ?", client.UserID, client.ID).
		Updates(&client)

	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Client with id %s does not exist", client.ID))
	}
	return client, nil
}

func (r *ClientRepository) Delete(ctx context.Context, userID, clientID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("user_id = ? AND id = ?", userID, clientID).
		Delete(models.Client{})

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Client with id %s does not exist", clientID))
	}
	return nil
}
