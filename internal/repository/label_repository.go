package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type LabelRepository struct {
	db *gorm.DB
}

func NewLabelRepository(db *gorm.DB) *LabelRepository {
	return &LabelRepository{db: db}
}

// FindAll labels for project or board
func (r *LabelRepository) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*[]model.Label, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(model.Label{}).
		Where("project_id = ?", projectID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.Label](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

// FindByID a label
func (r *LabelRepository) FindByID(ctx context.Context, projectID, labelID uuid.UUID) (*model.Label, error) {
	var label model.Label
	err := r.db.
		WithContext(ctx).
		Where("project_id = ? AND id = ?", projectID, labelID).
		First(&label).
		Error

	return &label, err
}

// Insert a new label
func (r *LabelRepository) Insert(ctx context.Context, label *model.Label) (*model.Label, error) {
	err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(label).
		Error

	return label, err
}

// Update label
func (r *LabelRepository) Update(ctx context.Context, label *model.Label) (*model.Label, error) {
	err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("project_id = ? AND id = ?", label.ProjectID, label.ID).
		Updates(label).
		Error

	return label, err
}

// Delete label
func (r *LabelRepository) Delete(ctx context.Context, projectID, labelID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("project_id = ? AND id = ?", projectID, labelID).
		Delete(&model.Label{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}
