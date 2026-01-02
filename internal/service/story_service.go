package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
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
	attachmentService  *AttachmentService
	ModelName          string
}

func NewStoryService(repo *repository.StoryRepository, ps *ProjectService, pus *ProjectUserService, kbs *KanbanBoardService, as *AttachmentService) *StoryService {
	return &StoryService{repo: repo, projectService: ps, projectUserService: pus, kanbanBoardService: kbs, attachmentService: as, ModelName: "story"}
}

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

func (s *StoryService) FindByID(ctx context.Context, projectID, storyID uuid.UUID) (*model.Story, error) {
	return s.repo.FindByID(ctx, projectID, storyID)
}

func (s *StoryService) FindBySlug(ctx context.Context, projectID uuid.UUID, slug string) (*model.Story, error) {
	return s.repo.FindBySlug(ctx, projectID, slug)
}

func (s *StoryService) Insert(ctx context.Context, userID, projectID uuid.UUID, req *model.StoryRequest) (*model.Story, error) {
	story := &model.Story{
		StoryListItem: model.StoryListItem{
			ProjectID:  projectID,
			ColumnID:   req.ColumnID,
			AssigneeID: req.AssigneeID,
			Title:      req.Title,
			Kind:       req.Kind,
			Priority:   req.Priority,
			Size:       req.Size,
			Estimate:   req.Estimate,
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

func (s *StoryService) Update(ctx context.Context, projectID, storyID uuid.UUID, req *model.StoryRequest) (*model.Story, error) {
	story := &model.Story{
		StoryListItem: model.StoryListItem{
			ID:         storyID,
			ProjectID:  projectID,
			ColumnID:   req.ColumnID,
			AssigneeID: req.AssigneeID,
			Title:      req.Title,
			Kind:       req.Kind,
			Priority:   req.Priority,
			Size:       req.Size,
			Estimate:   req.Estimate,
		},
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	return s.repo.Update(ctx, story)
}

func (s *StoryService) Delete(ctx context.Context, projectID, storyID uuid.UUID) error {
	return s.repo.Delete(ctx, projectID, storyID)
}

func (s *StoryService) GetAssignee(ctx context.Context, projectID, storyID uuid.UUID) (*model.User, error) {
	return s.repo.GetAssignee(ctx, projectID, storyID)
}

func (s *StoryService) ChangeAssignee(ctx context.Context, projectID, storyID uuid.UUID, input *model.StoryAssigneeRequest) (*model.StoryAssigneeRequest, error) {
	if err := s.repo.ChangeAssignee(ctx, projectID, storyID, input.AssigneeID); err != nil {
		return nil, err
	}
	return input, nil
}

func (s *StoryService) SaveAttachments(c *gin.Context, projectID uuid.UUID, storyID uuid.UUID, files []multipart.FileHeader) error {
	for _, file := range files {
		// ! projectID is passed instead of userID, so all members can access the list of attachments
		if _, err := s.attachmentService.SaveAttachment(c, projectID, s.ModelName, storyID.String(), file); err != nil {
			return err
		}
	}
	return nil
}

func (s *StoryService) ListAttachments(ctx context.Context, projectID uuid.UUID, storyID uuid.UUID) ([]*model.Attachment, error) {
	return s.attachmentService.FindByModelID(ctx, projectID, storyID, s.ModelName)
}
