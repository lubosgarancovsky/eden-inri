package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"gorm.io/gorm"
)

type StoryLabelRepository struct {
	db *gorm.DB
}

func NewStoryLabelRepository(db *gorm.DB) *StoryLabelRepository {
	return &StoryLabelRepository{db: db}
}

func (r *StoryLabelRepository) AssignLabel(ctx context.Context, projectID uuid.UUID, storyLabel *model.StoryLabel) error {
	var label model.Label
	if err := r.db.
		WithContext(ctx).
		Model(&model.Label{}).
		Where("id = ? AND project_id = ?", storyLabel.LabelID, projectID).
		First(&label).Error; err != nil {
		return err
	}

	return r.db.
		WithContext(ctx).
		FirstOrCreate(&storyLabel, storyLabel).
		Error
}

func (r *StoryLabelRepository) UnassignLabel(ctx context.Context, storyID, labelID uuid.UUID) error {
	return r.db.
		WithContext(ctx).
		Where("story_id = ? AND label_id = ?", storyID, labelID).
		Delete(&model.StoryLabel{}).
		Error
}

func (r *StoryLabelRepository) ListLabels(ctx context.Context, storyID uuid.UUID) ([]model.Label, error) {
	var labels []model.Label
	err := r.db.
		WithContext(ctx).
		Model(&model.Label{}).
		Joins("JOIN inri_story_labels sl ON sl.label_id = inri_labels.id").
		Where("sl.story_id = ?", storyID).
		Find(&labels).Error
	return labels, err
}
