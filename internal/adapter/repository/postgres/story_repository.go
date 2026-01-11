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

type StoryRepository struct {
	db *gorm.DB
}

func NewStoryRepository(db *gorm.DB) *StoryRepository {
	return &StoryRepository{db: db}
}

func (r *StoryRepository) Create(ctx context.Context, story *entity.Story) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.StoryFromDomain(story)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *StoryRepository) Update(ctx context.Context, story *entity.Story) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("id = ? AND project_id = ?", story.ID, story.ProjectID).
		Updates(mapper.StoryFromDomain(story))

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("story with id %s not found in project %s", story.ID, story.ProjectID))
	}

	return nil
}

func (r *StoryRepository) Delete(ctx context.Context, projectID, storyID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.
		Where("id = ? AND project_id = ?", storyID, projectID).
		Delete(&model.Story{})

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("story with id %s not found in project %s", storyID, projectID))
	}

	return nil
}

func (r *StoryRepository) FindByID(ctx context.Context, projectID, storyID uuid.UUID) (*entity.Story, error) {
	db := GetDB(ctx, r.db)
	var story model.Story
	if err := db.
		Preload("Assignee").
		Preload("Creator").
		Where("id = ? AND project_id = ?", storyID, projectID).
		First(&story).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("story with id %s not found in project %s", storyID, projectID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return story.ToDomain(), nil
}

func (r *StoryRepository) FindBySlug(ctx context.Context, projectID uuid.UUID, slug string) (*entity.Story, error) {
	db := GetDB(ctx, r.db)
	var story model.Story
	if err := db.
		Preload("Assignee").
		Preload("Creator").
		Where("slug = ? AND project_id = ?", slug, projectID).
		First(&story).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("story with slug %s not found in project %s", slug, projectID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return story.ToDomain(), nil
}

func (r *StoryRepository) List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Story, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.Story{}).
		Preload("Assignee").
		Preload("Creator").
		Where("project_id = ?", projectID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.Story, entity.Story](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *StoryRepository) ListAssigned(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Story, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.Story{}).
		Preload("Assignee").
		Preload("Creator").
		Where("assignee_id = ?", userID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.Story, entity.Story](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *StoryRepository) ChangeAssignee(ctx context.Context, projectID, storyID uuid.UUID, assigneeID *uuid.UUID) error {
	db := GetDB(ctx, r.db)
	result := db.Model(&model.Story{}).
		Where("id = ? AND project_id = ?", storyID, projectID).
		Update("assignee_id", assigneeID)

	if result.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, result.Error)
	}

	if result.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("story with id %s not found in project %s", storyID, projectID))
	}

	return nil
}
