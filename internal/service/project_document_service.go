package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ProjectDocumentService struct {
	r          *repository.ProjectDocumentRepository
	projectRep *repository.ProjectRepository
}

func NewProjectDocumentService(r *repository.ProjectDocumentRepository, projectRep *repository.ProjectRepository) *ProjectDocumentService {
	return &ProjectDocumentService{r: r, projectRep: projectRep}
}

func (s *ProjectDocumentService) FindAll(userID uuid.UUID, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.ProjectDocument], error) {
	// ensure user has access to the project
	prj, err := s.projectRep.FindByID(projectID)
	if err != nil {
		return nil, err
	}
	if prj.UserID != userID {
		return nil, api_err.ErrForbidden
	}

	items, totalCount, err := s.r.FindAll(userID, projectID, lq)
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
	doc, err := s.r.FindByID(projectID, id)
	if err != nil {
		return nil, err
	}
	if doc.UserID != userID {
		return nil, api_err.ErrForbidden
	}
	return doc, nil
}

func (s *ProjectDocumentService) Create(userID uuid.UUID, projectID uuid.UUID, input *model.ProjectDocumentRequest) (*model.ProjectDocument, error) {
	// ensure project belongs to user
	prj, err := s.projectRep.FindByID(projectID)
	if err != nil {
		return nil, err
	}
	if prj.UserID != userID {
		return nil, api_err.ErrForbidden
	}

	doc := &model.ProjectDocument{
		UserID:    userID,
		ProjectID: projectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
	return s.r.Insert(doc)
}

func (s *ProjectDocumentService) Update(userID uuid.UUID, projectID uuid.UUID, id uuid.UUID, input *model.ProjectDocumentRequest) (*model.ProjectDocument, error) {
	_, err := s.FindByID(userID, projectID, id)
	if err != nil {
		return nil, err
	}
	doc := &model.ProjectDocument{
		ID:        id,
		UserID:    userID,
		ProjectID: projectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
	return s.r.Update(doc)
}

func (s *ProjectDocumentService) Delete(userID uuid.UUID, projectID uuid.UUID, id uuid.UUID) (*model.ProjectDocument, error) {
	doc, err := s.FindByID(userID, projectID, id)
	if err != nil {
		return nil, err
	}
	return doc, s.r.Delete(id)
}
