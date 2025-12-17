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

// AttachmentRepository provides DB access for attachments
type AttachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) *AttachmentRepository {
	return &AttachmentRepository{db: db}
}

func (r *AttachmentRepository) FindAll(userID uuid.UUID, lq *list.ListingQuery) ([]model.Attachment, int64, error) {
	query := r.db.Model(&model.Attachment{}).Where("user_id = ?", userID)
	if lq != nil && lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.Attachment](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *AttachmentRepository) FindByID(attachmentID uuid.UUID) (*model.Attachment, error) {
	var result model.Attachment
	if err := r.db.Model(&model.Attachment{}).Where("id = ?", attachmentID).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *AttachmentRepository) FindByModelID(userID uuid.UUID, modelName string, modelID uuid.UUID) ([]*model.Attachment, error) {
	query := r.db.Model(&model.Attachment{}).Where("user_id = ?", userID)
	query = query.Where("model = ? AND model_id = ?", modelName, modelID)

	items := make([]*model.Attachment, 0)
	err := query.Find(&items).Error
	if err != nil {
		return items, err
	}

	return items, nil
}

func (r *AttachmentRepository) Insert(att *model.Attachment) (*model.Attachment, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(att).Error; err != nil {
		return nil, err
	}
	return att, nil
}

func (r *AttachmentRepository) Update(att *model.Attachment) (*model.Attachment, error) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", att.ID).Updates(&att)
	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Attachment with id %s does not exist", att.ID))
	}
	return att, nil
}

func (r *AttachmentRepository) Delete(attachmentID uuid.UUID) error {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", attachmentID).Delete(&model.Attachment{})
	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Attachment with id %s does not exist", attachmentID))
	}
	return nil
}
