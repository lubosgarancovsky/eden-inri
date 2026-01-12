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

type AttachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) *AttachmentRepository {
	return &AttachmentRepository{db: db}
}

func (r *AttachmentRepository) List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Attachment, int64, error) {
	db := GetDB(ctx, r.db)

	query := db.Model(&model.Attachment{}).Where("user_id = ?", userID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.Attachment, entity.Attachment](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *AttachmentRepository) FindByID(ctx context.Context, userID, attachmentID uuid.UUID) (*entity.Attachment, error) {
	db := GetDB(ctx, r.db)

	var att model.Attachment
	if err := db.Where("id = ?", attachmentID).Where("user_id = ?", userID).First(&att).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("attachment with id %s not found", attachmentID))
		}

		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	domainAtt := att.ToDomain()
	return domainAtt, nil
}

func (r *AttachmentRepository) FindByModelID(ctx context.Context, userID, modelID uuid.UUID, modelName string) ([]entity.Attachment, error) {
	db := GetDB(ctx, r.db)

	var items []model.Attachment
	if err := db.Where("user_id = ? AND model_id = ? AND model = ?", userID, modelID, modelName).Find(&items).Error; err != nil {
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}

	return MapToDomain[model.Attachment, entity.Attachment](items), nil
}

func (r *AttachmentRepository) Create(ctx context.Context, attachment *entity.Attachment) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.AttachmentFromDomain(attachment)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *AttachmentRepository) Delete(ctx context.Context, userID, attachmentID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("user_id = ?", userID).
		Where("id = ?", attachmentID).
		Delete(&model.Attachment{})

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("attachment with id %s not found", attachmentID))
	}

	return nil
}
