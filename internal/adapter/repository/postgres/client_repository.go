package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
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

func (r *ClientRepository) List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Client, int64, error) {
	db := GetDB(ctx, r.db)

	query := db.Model(&model.Client{}).Where("user_id = ?", userID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.Client, entity.Client](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *ClientRepository) FindByID(ctx context.Context, userID, clientID uuid.UUID) (*entity.Client, error) {
	db := GetDB(ctx, r.db)

	var client model.Client
	if err := db.Where("client_id = ?", clientID).Where("user_id = ?", userID).First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("client with id %s not found", clientID))
		}

		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	return client.ToDomain(), nil
}

func (r *ClientRepository) Create(ctx context.Context, client *entity.Client) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.ClientFromDomain(client)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *ClientRepository) Update(ctx context.Context, client *entity.Client) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("id", client.ID).
		Where("user_id = ?", client.UserID).
		Updates(mapper.ClientFromDomain(client))

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("client with id %s not found", client.ID))
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
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("client with id %s not found", clientID))
	}

	return nil
}
