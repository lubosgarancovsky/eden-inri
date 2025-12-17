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
	prj, err := s.r.FindByID(id, userID)
	if err != nil {
		return nil, err
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
	return s.r.Insert(userID, prj)
}

func (s *ProjectService) Update(userID, id uuid.UUID, input *model.ProjectRequest) (*model.Project, error) {
	role, err := s.GetUserRole(id, userID)
	if err != nil {
		return nil, err
	}

	// Only owner or admin
	if err := s.RequireRole(role, model.Owner, model.Admin); err != nil {
		return nil, err
	}

	if _, err := s.FindByID(userID, id); err != nil {
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

func (s *ProjectService) Delete(userID, id uuid.UUID) (*model.Project, error) {
	role, err := s.GetUserRole(id, userID)
	if err != nil {
		return nil, err
	}

	// Only owner
	if err := s.RequireRole(role, model.Owner); err != nil {
		return nil, err
	}

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

func (s *ProjectService) ListAttachments(userID, projectID uuid.UUID) ([]*model.Attachment, error) {
	// ensure user has access
	if _, err := s.FindByID(userID, projectID); err != nil {
		return nil, err
	}
	return s.attachmentService.FindByModelID(userID, s.ModelName, projectID)
}

func (s *ProjectService) ListProjectMembers(userID, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.ProjectUser], error) {
	items, totalCount, err := s.r.ListProjectMembers(userID, projectID, lq)
	if err != nil {
		return nil, err
	}
	return &list.Page[model.ProjectUser]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ProjectService) AddProjectMember(userID, projectID uuid.UUID, role model.ProjectRole) error {
	role, err := s.GetUserRole(projectID, userID)
	if err != nil {
		return err
	}

	if err := s.RequireRole(role, model.Admin, model.Owner); err != nil {
		return err
	}

	return s.AddProjectMember(userID, projectID, role)
}

func (s *ProjectService) RemoveProjectMember(userID, memberID, projectID uuid.UUID) error {
	role, err := s.GetUserRole(projectID, userID)
	if err != nil {
		return err
	}

	if err := s.RequireRole(role, model.Admin, model.Owner); err != nil {
		return err
	}

	memberRole, err := s.GetUserRole(projectID, memberID)
	if err != nil {
		return err
	}

	if memberRole == model.Owner {
		return api_err.ErrForbidden
	}

	return s.r.RemoveMember(userID, projectID)
}

func (s *ProjectService) GetUserRole(projectID, userID uuid.UUID) (model.ProjectRole, error) {
	return s.r.GetUserRole(projectID, userID)
}

func (s *ProjectService) RequireRole(actual model.ProjectRole, allowed ...model.ProjectRole) error {
	for _, r := range allowed {
		if actual == r {
			return nil
		}
	}
	return api_err.ErrForbidden
}
