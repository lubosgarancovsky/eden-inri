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
func (s *StoryService) FindAll(ctx context.Context, boardID, columnID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.Story], error) {
	items, totalCount, err := s.repo.FindAll(ctx, boardID, columnID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.Story]{
		Items:      *items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

// FindByID story
func (s *StoryService) FindByID(ctx context.Context, boardID, storyID uuid.UUID) (*model.Story, error) {
	return s.repo.FindByID(ctx, boardID, storyID)
}

// Insert story
func (s *StoryService) Insert(ctx context.Context, boardID uuid.UUID, req *model.StoryRequest) (*model.Story, error) {
	story := &model.Story{
		BoardID:     &boardID,
		ColumnID:    req.ColumnID,
		Title:       req.Title,
		Description: req.Description,
		Kind:        req.Kind,
		Priority:    req.Priority,
		Size:        req.Size,
		Estimate:    req.Estimate,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	err := s.repo.WithTx(ctx, func(txRepo *repository.StoryRepository) error {
		var project model.Project
		var projectID uuid.UUID

		if err := txRepo.DB().
			WithContext(ctx).
			Model(model.KanbanBoard{}).
			Select("project_id").
			Where("id = ?", boardID).
			Scan(projectID).
			Error; err != nil {
			return err
		}

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
func (s *StoryService) Update(ctx context.Context, boardID uuid.UUID, req *model.StoryRequest) (*model.Story, error) {
	story := &model.Story{
		BoardID:     &boardID,
		ColumnID:    req.ColumnID,
		Title:       req.Title,
		Description: req.Description,
		Kind:        req.Kind,
		Priority:    req.Priority,
		Size:        req.Size,
		Estimate:    req.Estimate,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	return s.repo.Update(ctx, story)
}

// Delete story
func (s *StoryService) Delete(ctx context.Context, boardID, storyID uuid.UUID) error {
	return s.repo.Delete(ctx, boardID, storyID)
}
