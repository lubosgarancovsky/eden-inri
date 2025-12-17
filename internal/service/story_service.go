package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
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
func (s *StoryService) FindAll(userID, boardID, columnID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.Story], error) {
	projectID, err := s.kanbanBoardService.GetProjectIDByBoardID(boardID)
	if err != nil {
		return nil, err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}

	items, totalCount, err := s.repo.FindAll(columnID, lq)
	return &list.Page[model.Story]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

// FindByID story
func (s *StoryService) FindByID(userID, storyID, projectID uuid.UUID) (*model.Story, error) {
	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}
	return s.repo.FindByID(storyID)
}

// Insert story
func (s *StoryService) Insert(userID, boardID uuid.UUID, req *model.StoryRequest) (*model.Story, error) {
	projectID, err := s.kanbanBoardService.GetProjectIDByBoardID(boardID)
	if err != nil {
		return nil, err
	}

	var story *model.Story
	err = s.repo.DB().Transaction(func(tx *gorm.DB) error {
		var project model.Project
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", projectID).First(&project).Error; err != nil {
			return err
		}

		if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
			return err
		}

		// Increment last_story_number
		project.StorySequence++
		if err := tx.Save(&project).Error; err != nil {
			return err
		}

		// Generate sequential slug
		storySlug := fmt.Sprintf("%s-%d", project.Slug, project.StorySequence)

		story = &model.Story{
			ProjectID:   projectID,
			BoardID:     &boardID,
			ColumnID:    req.ColumnID,
			Slug:        storySlug,
			Title:       req.Title,
			Description: req.Description,
			Kind:        req.Kind,
			Priority:    req.Priority,
			Size:        req.Size,
			Estimate:    req.Estimate,
			StartDate:   req.StartDate,
			EndDate:     req.EndDate,
		}

		return tx.Create(story).Error
	})

	if err != nil {
		return nil, err
	}

	return story, nil
}

// Update story
func (s *StoryService) Update(userID, storyID, boardID uuid.UUID, req *model.StoryRequest) (*model.Story, error) {
	projectID, err := s.kanbanBoardService.GetProjectIDByBoardID(boardID)
	if err != nil {
		return nil, err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}

	story, err := s.repo.FindByID(storyID)
	if err != nil {
		return nil, err
	}

	// Only editable fields
	story.Title = req.Title
	story.Description = req.Description
	story.Estimate = req.Estimate
	story.Size = req.Size
	story.Priority = req.Priority
	story.StartDate = req.StartDate
	story.EndDate = req.EndDate
	story.ColumnID = req.ColumnID

	return s.repo.Update(story)
}

// Delete story
func (s *StoryService) Delete(userID, storyID, boardID uuid.UUID) error {
	projectID, err := s.kanbanBoardService.GetProjectIDByBoardID(boardID)
	if err != nil {
		return err
	}

	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return err
	}

	return s.repo.Delete(storyID)
}
