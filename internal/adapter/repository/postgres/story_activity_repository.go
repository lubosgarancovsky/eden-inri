package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type StoryActivityRepository struct {
	db *gorm.DB
}

func NewStoryActivityRepository(db *gorm.DB) *StoryActivityRepository {
	return &StoryActivityRepository{db: db}
}

func (r *StoryActivityRepository) Create(ctx context.Context, activity *entity.StoryActivity) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.StoryActivityFromDomain(activity)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *StoryActivityRepository) Update(ctx context.Context, activity *entity.StoryActivity) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("id = ?", activity.ID).
		Updates(mapper.StoryActivityFromDomain(activity))

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("activity with id %s not found", activity.ID))
	}

	return nil
}

func (r *StoryActivityRepository) Delete(ctx context.Context, activityID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("id = ?", activityID).
		Delete(&model.StoryActivity{})

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("activity with id %s not found", activityID))
	}

	return nil
}

func (r *StoryActivityRepository) FindByID(ctx context.Context, activityID uuid.UUID) (*entity.StoryActivity, error) {
	db := GetDB(ctx, r.db)
	var activity model.StoryActivity
	if err := db.
		Preload("Actor").
		Where("id = ?", activityID).
		First(&activity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("activity with id %s not found", activityID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return activity.ToDomain(), nil
}

func (r *StoryActivityRepository) List(ctx context.Context, storyID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.StoryActivity, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.StoryActivity{}).
		Preload("Actor").
		Where("story_id = ?", storyID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.StoryActivity, entity.StoryActivity](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}
