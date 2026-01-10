package services

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ProjectService struct {
	r                  *repositories.ProjectRepository
	attachmentService  *AttachmentService
	projectUserService *ProjectUserService
	ModelName          string
}

func NewProjectService(r *repositories.ProjectRepository, attachmentService *AttachmentService, projectUserService *ProjectUserService) *ProjectService {
	return &ProjectService{r: r, attachmentService: attachmentService, projectUserService: projectUserService, ModelName: "project"}
}

func (s *ProjectService) FindAll(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]models.Project, int64, error) {
	return s.r.FindAll(ctx, userID, lq)
}

func (s *ProjectService) FindByID(ctx context.Context, userID, projectID uuid.UUID) (*models.Project, error) {
	return s.r.FindByID(ctx, userID, projectID)
}

func (s *ProjectService) Create(ctx context.Context, userID uuid.UUID, input *models.ProjectRequest) (*models.Project, error) {
	var project *models.Project
	if err := s.r.WithTx(ctx, func(txRepo *repositories.ProjectRepository) error {
		result, err := txRepo.InsertProject(ctx, s.buildPayload(input))
		if err != nil {
			return err
		}

		project = result

		pu := &models.ProjectUser{
			ProjectID: project.ID,
			UserID:    userID,
			Role:      models.Owner,
		}

		if _, err = txRepo.InsertProjectUser(ctx, pu); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) Update(ctx context.Context, projectID uuid.UUID, input *models.ProjectRequest) (*models.Project, error) {
	project := s.buildPayload(input)
	project.ID = projectID
	return s.r.Update(ctx, project)
}

func (s *ProjectService) Delete(ctx context.Context, projectID uuid.UUID) error {
	return s.r.Delete(ctx, projectID)
}

func (s *ProjectService) SaveAttachments(c *gin.Context, projectID uuid.UUID, files []multipart.FileHeader) error {
	owner, err := s.projectUserService.GetOwner(c, projectID)
	if err != nil {
		return err
	}

	for _, file := range files {
		if _, err := s.attachmentService.SaveAttachment(c, owner.UserID, s.ModelName, projectID.String(), file); err != nil {
			return err
		}
	}
	return nil
}

func (s *ProjectService) ListAttachments(ctx context.Context, projectID uuid.UUID) ([]*models.Attachment, error) {
	owner, err := s.projectUserService.GetOwner(ctx, projectID)
	if err != nil {
		return []*models.Attachment{}, err
	}

	return s.attachmentService.FindByModelID(ctx, owner.UserID, projectID, s.ModelName)
}

func (s *ProjectService) DownloadAttachment(c *gin.Context, projectID, attachmentID uuid.UUID) error {
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

func (s *ProjectService) DeleteAttachment(ctx context.Context, projectID, attachmentID uuid.UUID) error {
	owner, err := s.projectUserService.GetOwner(ctx, projectID)
	if err != nil {
		return err
	}

	return s.attachmentService.Delete(ctx, owner.UserID, attachmentID)
}

func (s *ProjectService) Favourite(ctx context.Context, userID, projectID uuid.UUID) (*models.Project, error) {
	prj, err := s.r.FindByID(ctx, userID, projectID)
	if err != nil {
		return nil, err
	}

	projectUser, err := s.projectUserService.Favourite(ctx, userID, projectID)
	if err != nil {
		return nil, err
	}

	prj.IsStarred = projectUser.IsStarred
	return prj, nil
}

func (s *ProjectService) buildPayload(input *models.ProjectRequest) *models.Project {
	return &models.Project{
		Name:           input.Name,
		Description:    input.Description,
		Status:         input.Status,
		Tags:           input.Tags,
		Slug:           input.Slug,
		UpdatedAt:      time.Now(),
		LastActivityAt: time.Now(),
	}
}
