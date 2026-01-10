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

type AttachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) *AttachmentRepository {
	return &AttachmentRepository{db: db}
}

func (r *AttachmentRepository) FindAll(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]models.Attachment, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(&models.Attachment{}).
		Where("user_id = ?", userID)

	if lq != nil && lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[models.Attachment](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *AttachmentRepository) FindByID(ctx context.Context, userID, attachmentID uuid.UUID) (*models.Attachment, error) {
	var result models.Attachment
	if err := r.db.
		WithContext(ctx).
		Model(&models.Attachment{}).
		Where("user_id = ? AND id = ?", userID, attachmentID).
		First(&result).
		Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *AttachmentRepository) FindByModelID(ctx context.Context, userID uuid.UUID, modelID uuid.UUID, modelName string) ([]*models.Attachment, error) {
	items := make([]*models.Attachment, 0)

	if err := r.db.
		WithContext(ctx).
		Model(models.Attachment{}).
		Where("user_id = ? AND model_id = ? AND model = ?", userID, modelID, modelName).
		Find(&items).
		Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *AttachmentRepository) Insert(ctx context.Context, payload *models.Attachment) (*models.Attachment, error) {
	if err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(payload).
		Error; err != nil {
		return nil, err
	}
	return payload, nil
}

func (r *AttachmentRepository) Update(ctx context.Context, payload *models.Attachment) (*models.Attachment, error) {
	result := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("id = ?", payload.ID).
		Updates(&payload)

	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Attachment with id %s does not exist", payload.ID))
	}
	return payload, nil
}

func (r *AttachmentRepository) Delete(ctx context.Context, userID, attachmentID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("user_id = ? AND id = ?", userID, attachmentID).
		Delete(models.Attachment{})

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Attachment with id %s does not exist", attachmentID))
	}
	return nil
}
