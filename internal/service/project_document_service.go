package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ProjectDocumentService struct {
	r                  *repository.ProjectDocumentRepository
	projectUserService *ProjectUserService
}

func NewProjectDocumentService(r *repository.ProjectDocumentRepository, pus *ProjectUserService) *ProjectDocumentService {
	return &ProjectDocumentService{r: r, projectUserService: pus}
}

func (s *ProjectDocumentService) FindAll(userID uuid.UUID, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.ProjectDocument], error) {
	// ensure user has access to the project
	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}

	items, totalCount, err := s.r.FindAll(projectID, lq)
	if err != nil {
		return nil, err
	}
	return &list.Page[model.ProjectDocument]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ProjectDocumentService) FindByID(userID uuid.UUID, projectID uuid.UUID, id uuid.UUID) (*model.ProjectDocument, error) {
	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}
	return s.r.FindByID(projectID, id)
}

func (s *ProjectDocumentService) Create(userID uuid.UUID, projectID uuid.UUID, input *model.ProjectDocumentRequest) (*model.ProjectDocument, error) {
	pu, err := s.projectUserService.GetProjectUserIfMember(projectID, userID)
	if err != nil {
		return nil, err
	}

	if err := s.projectUserService.RequireRole(pu.Role, model.Admin, model.Owner, model.Developer); err != nil {
		return nil, err
	}

	doc := &model.ProjectDocument{
		ProjectID: projectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
	return s.r.Insert(doc)
}

func (s *ProjectDocumentService) Update(userID uuid.UUID, projectID uuid.UUID, id uuid.UUID, input *model.ProjectDocumentRequest) (*model.ProjectDocument, error) {
	pu, err := s.projectUserService.GetProjectUserIfMember(projectID, userID)
	if err != nil {
		return nil, err
	}

	if err := s.projectUserService.RequireRole(pu.Role, model.Admin, model.Owner, model.Developer); err != nil {
		return nil, err
	}

	doc := &model.ProjectDocument{
		ID:        id,
		ProjectID: projectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
	return s.r.Update(doc)
}

func (s *ProjectDocumentService) Delete(userID uuid.UUID, projectID uuid.UUID, id uuid.UUID) (*model.ProjectDocument, error) {
	pu, err := s.projectUserService.GetProjectUserIfMember(projectID, userID)
	if err != nil {
		return nil, err
	}

	if err := s.projectUserService.RequireRole(pu.Role, model.Admin, model.Owner, model.Developer); err != nil {
		return nil, err
	}

	doc, err := s.FindByID(userID, projectID, id)
	if err != nil {
		return nil, err
	}
	return doc, s.r.Delete(id)
}
