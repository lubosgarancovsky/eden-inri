package repository

import (
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

// AssignLabel adds a label to a story
func (r *StoryLabelRepository) AssignLabel(storyID, labelID uuid.UUID) error {
	storyLabel := model.StoryLabel{
		StoryID: storyID,
		LabelID: labelID,
	}
	return r.db.FirstOrCreate(&storyLabel, storyLabel).Error
}

// UnassignLabel removes a label from a story
func (r *StoryLabelRepository) UnassignLabel(storyID, labelID uuid.UUID) error {
	return r.db.Where("story_id = ? AND label_id = ?", storyID, labelID).Delete(&model.StoryLabel{}).Error
}

// ListLabels returns all labels assigned to a story
func (r *StoryLabelRepository) ListLabels(storyID uuid.UUID) ([]model.Label, error) {
	var labels []model.Label
	err := r.db.Joins("JOIN inri_kanban_story_labels sl ON sl.label_id = inri_kanban_label.id").
		Where("sl.story_id = ?", storyID).
		Find(&labels).Error
	return labels, err
}
