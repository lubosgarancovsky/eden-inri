package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ProjectService struct {
	r                  *repository.ProjectRepository
	attachmentService  *AttachmentService
	projectUserService *ProjectUserService
	ModelName          string
}

func NewProjectService(r *repository.ProjectRepository, attachmentService *AttachmentService, projectUserService *ProjectUserService) *ProjectService {
	return &ProjectService{r: r, attachmentService: attachmentService, projectUserService: projectUserService, ModelName: "project"}
}

func (s *ProjectService) FindAll(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]model.Project, int64, error) {
	return s.r.FindAll(ctx, userID, lq)
}

func (s *ProjectService) FindByID(ctx context.Context, userID, projectID uuid.UUID) (*model.Project, error) {
	return s.r.FindByID(ctx, userID, projectID)
}

func (s *ProjectService) Create(ctx context.Context, userID uuid.UUID, input *model.ProjectRequest) (*model.Project, error) {
	var project *model.Project
	if err := s.r.WithTx(ctx, func(txRepo *repository.ProjectRepository) error {
		result, err := txRepo.InsertProject(ctx, s.buildPayload(input))
		if err != nil {
			return err
		}

		project = result

		pu := &model.ProjectUser{
			ProjectID: project.ID,
			UserID:    userID,
			Role:      model.Owner,
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

func (s *ProjectService) Update(ctx context.Context, projectID uuid.UUID, input *model.ProjectRequest) (*model.Project, error) {
	project := s.buildPayload(input)
	project.ID = projectID
	return s.r.Update(ctx, project)
}

func (s *ProjectService) Delete(ctx context.Context, projectID uuid.UUID) error {
	return s.r.Delete(ctx, projectID)
}

func (s *ProjectService) SaveAttachments(c *gin.Context, projectID uuid.UUID, files []multipart.FileHeader) error {
	for _, file := range files {
		// Pass in projectID instead of userID as second param, to allow other members to list the attachments
		if _, err := s.attachmentService.SaveAttachment(c, projectID, s.ModelName, projectID.String(), file); err != nil {
			return err
		}
	}
	return nil
}

func (s *ProjectService) ListAttachments(ctx context.Context, projectID uuid.UUID) ([]*model.Attachment, error) {
	// Second parameter is projectID instead of required userID
	// projectID is a common determinant so every member can see all uploaded attachments, instead of just their own
	return s.attachmentService.FindByModelID(ctx, projectID, projectID, s.ModelName)
}

func (s *ProjectService) Favourite(ctx context.Context, userID, projectID uuid.UUID) (*model.Project, error) {
	prj, err := s.r.FindByID(ctx, projectID, userID)
	if err != nil {
		return nil, err
	}

	projectUser, err := s.projectUserService.Favourite(userID, projectID)
	if err != nil {
		return nil, err
	}

	prj.IsStarred = projectUser.IsStarred
	return prj, nil
}

func (s *ProjectService) buildPayload(input *model.ProjectRequest) *model.Project {
	return &model.Project{
		Name:           input.Name,
		Description:    input.Description,
		Status:         input.Status,
		Tags:           input.Tags,
		UpdatedAt:      time.Now(),
		LastActivityAt: time.Now(),
	}
}
