package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (r *ClientRepository) List(ctx context.Context) ([]*entity.Client, int64, error) {
	// TODO: Handle listing query
	return []*entity.Client{}, 0, nil
}

func (r *ClientRepository) FindByID(ctx context.Context, userID, clientID uuid.UUID) (*entity.Client, error) {
	db := GetDB(ctx, r.db)

	var client model.Client
	if err := db.Where("client_id = ?", clientID).Where("user_id = ?", userID).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, api_err.ErrNotFound.WithMessage(fmt.Sprintf("client with id %s not found", clientID))
		}

		return nil, api_err.Wrap(api_err.ErrInternalServer, err)
	}

	return client.ToDomain(), nil
}

func (r *ClientRepository) Create(ctx context.Context, client *entity.Client) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(model.ClientFromDomain(client)).Error; err != nil {
		return api_err.Wrap(api_err.ErrInternalServer, err)
	}
	return nil
}

func (r *ClientRepository) Update(ctx context.Context, client *entity.Client) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("id", client.ID).
		Where("user_id = ?", client.UserID).
		Updates(model.ClientFromDomain(client))

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("client with id %s not found", client.ID))
	}

	return nil
}

func (r *ClientRepository) Delete(ctx context.Context, userID, clientID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("user_id = ?", userID).
		Where("client_id = ?", clientID).
		Delete(&model.Client{})

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("client with id %s not found", clientID))
	}

	return nil
}
