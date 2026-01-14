package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"gorm.io/gorm"
)

type StoryLabelRepository struct {
	db *gorm.DB
}

func NewStoryLabelRepository(db *gorm.DB) *StoryLabelRepository {
	return &StoryLabelRepository{db: db}
}

func (r *StoryLabelRepository) List(ctx context.Context, storyID uuid.UUID) (*[]entity.StoryLabel, error) {
	db := GetDB(ctx, r.db)

	var labels []model.StoryLabel
	if err := db.Where("story_id = ?", storyID).Find(&labels).Error; err != nil {
		return nil, err
	}

	entities := make([]entity.StoryLabel, len(labels))
	for i, label := range labels {
		entities[i] = *label.ToDomain()
	}
	return &entities, nil
}

func (r *StoryLabelRepository) Assign(ctx context.Context, storyID, labelID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	storyLabel := model.StoryLabel{StoryID: storyID, LabelID: labelID}
	return db.Model(&storyLabel).Create(&storyLabel).Error
}

func (r *StoryLabelRepository) Unassign(ctx context.Context, storyID, labelID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	return db.Where("story_id = ? AND label_id = ?", storyID, labelID).Delete(&model.StoryLabel{}).Error
}
