package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

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
func (r *LabelRepository) FindAll(projectID uuid.UUID) ([]model.Label, error) {
	var labels []model.Label
	query := r.db.Model(&model.Label{}).Where("project_id = ?", projectID)
	if err := query.Order("name ASC").Find(&labels).Error; err != nil {
		return nil, err
	}
	return labels, nil
}

// FindByID a label
func (r *LabelRepository) FindByID(labelID uuid.UUID) (*model.Label, error) {
	var label model.Label
	err := r.db.Where("id = ?", labelID).First(&label).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, api_err.ErrNotFound
	}
	return &label, err
}

// Insert a new label
func (r *LabelRepository) Insert(label *model.Label) (*model.Label, error) {
	if err := r.db.Create(label).Error; err != nil {
		return nil, err
	}
	return label, nil
}

// Update label
func (r *LabelRepository) Update(label *model.Label) (*model.Label, error) {
	if err := r.db.Save(label).Error; err != nil {
		return nil, err
	}
	return label, nil
}

// Delete label
func (r *LabelRepository) Delete(labelID uuid.UUID) error {
	result := r.db.Where("id = ?", labelID).Delete(&model.Label{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound
	}
	return nil
}
