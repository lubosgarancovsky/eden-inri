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
	repo                *repository.StoryRepository
	projectService      *ProjectService
	projectUserService  *ProjectUserService
	kanbanBoardService  *KanbanBoardService
	kanbanColumnService *KanbanColumnService
	attachmentService   *AttachmentService
	ModelName           string
}

func NewStoryService(repo *repository.StoryRepository, ps *ProjectService, pus *ProjectUserService, kbs *KanbanBoardService, kcs *KanbanColumnService, as *AttachmentService) *StoryService {
	return &StoryService{repo: repo, projectService: ps, projectUserService: pus, kanbanBoardService: kbs, kanbanColumnService: kcs, attachmentService: as, ModelName: "story"}
}

func (s *StoryService) FindAllAssigned(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.StoryListItem], error) {
	items, totalCount, err := s.repo.FindAllAssigned(ctx, userID, lq)
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
	boardID, err := s.kanbanColumnService.GetBoardID(ctx, req.ColumnID)
	if err != nil {
		return nil, err
	}

	story := &model.Story{
		StoryListItem: model.StoryListItem{
			ProjectID:  projectID,
			ColumnID:   req.ColumnID,
			BoardID:    *boardID,
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

	err = s.repo.WithTx(ctx, func(txRepo *repository.StoryRepository) error {
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

func (s *StoryService) SaveAttachments(ctx *gin.Context, projectID uuid.UUID, storyID uuid.UUID, files []multipart.FileHeader) error {
	owner, err := s.projectUserService.GetOwner(ctx, projectID)
	if err != nil {
		return err
	}

	for _, file := range files {
		if _, err := s.attachmentService.SaveAttachment(ctx, owner.UserID, s.ModelName, storyID.String(), file); err != nil {
			return err
		}
	}
	return nil
}

func (s *StoryService) ListAttachments(ctx context.Context, projectID uuid.UUID, storyID uuid.UUID) ([]*model.Attachment, error) {
	owner, err := s.projectUserService.GetOwner(ctx, projectID)
	if err != nil {
		return []*model.Attachment{}, err
	}

	return s.attachmentService.FindByModelID(ctx, owner.UserID, storyID, s.ModelName)
}

func (s *StoryService) DownloadAttachment(c *gin.Context, projectID, attachmentID uuid.UUID) error {
	ctx := c.Request.Context()

	owner, err := s.projectUserService.GetOwner(ctx, projectID)
	if err != nil {
		return err
	}

	attachment, err := s.attachmentService.FindByID(ctx, owner.UserID, attachmentID)
	if err != nil {
		return err
	}

	path := s.attachmentService.GetFilePath(attachment)
	c.FileAttachment(path, attachment.OriginalName)

	return nil
}

func (s *StoryService) DeleteAttachment(ctx context.Context, projectID, attachmentID uuid.UUID) error {
	owner, err := s.projectUserService.GetOwner(ctx, projectID)
	if err != nil {
		return err
	}

	return s.attachmentService.Delete(ctx, owner.UserID, attachmentID)
}
