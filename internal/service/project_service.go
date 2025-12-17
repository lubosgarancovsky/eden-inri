package service

import (
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ProjectService struct {
	r                 *repository.ProjectRepository
	attachmentService *AttachmentService
	ModelName         string
}

func NewProjectService(r *repository.ProjectRepository, attachmentService *AttachmentService) *ProjectService {
	return &ProjectService{r: r, attachmentService: attachmentService, ModelName: "project"}
}

func (s *ProjectService) FindAll(userID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.Project], error) {
	items, totalCount, err := s.r.FindAll(userID, lq)
	if err != nil {
		return nil, err
	}
	return &list.Page[model.Project]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ProjectService) FindByID(userID uuid.UUID, id uuid.UUID) (*model.Project, error) {
	prj, err := s.r.FindByID(id)
	if err != nil {
		return nil, err
	}
	if prj.UserID != userID {
		return nil, api_err.ErrForbidden
	}
	return prj, nil
}

func (s *ProjectService) Create(userID uuid.UUID, input *model.ProjectRequest) (*model.Project, error) {
	prj := &model.Project{
		UserID:      userID,
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status,
		Tags:        input.Tags,
		Slug:        input.Slug,
		IsStarred:   false,
	}
	return s.r.Insert(prj)
}

func (s *ProjectService) Update(userID uuid.UUID, id uuid.UUID, input *model.ProjectRequest) (*model.Project, error) {
	_, err := s.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	prj := &model.Project{
		ID:             id,
		UserID:         userID,
		Name:           input.Name,
		Description:    input.Description,
		Status:         input.Status,
		Tags:           input.Tags,
		Slug:           input.Slug,
		UpdatedAt:      time.Now(),
		LastActivityAt: time.Now(),
	}
	return s.r.Update(prj)
}

func (s *ProjectService) Delete(userID uuid.UUID, id uuid.UUID) (*model.Project, error) {
	prj, err := s.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	return prj, s.r.Delete(id)
}

func (s *ProjectService) SaveAttachments(c *gin.Context, userID uuid.UUID, projectID uuid.UUID, files []multipart.FileHeader) error {
	for _, file := range files {
		if _, err := s.attachmentService.SaveAttachment(c, userID, s.ModelName, projectID.String(), file); err != nil {
			return err
		}
	}
	return nil
}

func (s *ProjectService) ListAttachments(userID uuid.UUID, projectID uuid.UUID) ([]*model.Attachment, error) {
	// ensure user has access
	if _, err := s.FindByID(userID, projectID); err != nil {
		return nil, err
	}
	return s.attachmentService.FindByModelID(userID, s.ModelName, projectID)
}
