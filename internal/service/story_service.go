package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm/clause"
)

type StoryService struct {
	repo               *repository.StoryRepository
	projectService     *ProjectService
	projectUserService *ProjectUserService
	kanbanBoardService *KanbanBoardService
}

func NewStoryService(repo *repository.StoryRepository, ps *ProjectService, pus *ProjectUserService, kbs *KanbanBoardService) *StoryService {
	return &StoryService{repo: repo, projectService: ps, projectUserService: pus, kanbanBoardService: kbs}
}

// FindAll stories in a column
func (s *StoryService) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.StoryListItem], error) {
	items, totalCount, err := s.repo.FindAll(ctx, projectID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.StoryListItem]{
		Items:      *items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

// FindByID story
func (s *StoryService) FindByID(ctx context.Context, projectID, storyID uuid.UUID) (*model.Story, error) {
	return s.repo.FindByID(ctx, projectID, storyID)
}

// Insert story
func (s *StoryService) Insert(ctx context.Context, userID, projectID uuid.UUID, req *model.StoryRequest) (*model.Story, error) {
	story := &model.Story{
		StoryListItem: model.StoryListItem{
			ProjectID: projectID,
			ColumnID:  req.ColumnID,
			BoardID:   req.BoardID,
			Title:     req.Title,
			Kind:      req.Kind,
			Priority:  req.Priority,
			Size:      req.Size,
		},
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		CreatedBy:   userID,
	}

	err := s.repo.WithTx(ctx, func(txRepo *repository.StoryRepository) error {
		var project model.Project

		if err := txRepo.DB().
			WithContext(ctx).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", projectID).
			First(&project).Error; err != nil {
			return err
		}

		project.StorySequence++
		if err := txRepo.DB().Save(&project).Error; err != nil {
			return err
		}

		// Generate sequential slug
		story.Slug = fmt.Sprintf("%s-%d", project.Slug, project.StorySequence)
		story.ProjectID = projectID

		return txRepo.DB().
			WithContext(ctx).
			Clauses(clause.Returning{}).
			Create(story).Error
	})

	return story, err
}

// Update story
func (s *StoryService) Update(ctx context.Context, projectID, storyID uuid.UUID, req *model.StoryRequest) (*model.Story, error) {
	story := &model.Story{
		StoryListItem: model.StoryListItem{
			ID:        storyID,
			ProjectID: projectID,
			ColumnID:  req.ColumnID,
			Title:     req.Title,
			Kind:      req.Kind,
			Priority:  req.Priority,
			Size:      req.Size,
		},
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	return s.repo.Update(ctx, story)
}

// Delete story
func (s *StoryService) Delete(ctx context.Context, projectID, storyID uuid.UUID) error {
	return s.repo.Delete(ctx, projectID, storyID)
}
