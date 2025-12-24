package service

import (
	"context"

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

func (s *ProjectDocumentService) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.ProjectDocument], error) {
	items, totalCount, err := s.r.FindAll(ctx, projectID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.ProjectDocument]{
		Items:      *items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ProjectDocumentService) FindByID(ctx context.Context, projectID uuid.UUID, documentID uuid.UUID) (*model.ProjectDocument, error) {
	return s.r.FindByID(ctx, projectID, documentID)
}

func (s *ProjectDocumentService) Create(ctx context.Context, projectID uuid.UUID, input *model.ProjectDocumentRequest) (*model.ProjectDocument, error) {
	doc := &model.ProjectDocument{
		ProjectID: projectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
	return s.r.Insert(ctx, doc)
}

func (s *ProjectDocumentService) Update(ctx context.Context, projectID uuid.UUID, id uuid.UUID, input *model.ProjectDocumentRequest) (*model.ProjectDocument, error) {
	doc := &model.ProjectDocument{
		ID:        id,
		ProjectID: projectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
	return s.r.Update(ctx, doc)
}

func (s *ProjectDocumentService) Delete(ctx context.Context, projectID uuid.UUID, documentID uuid.UUID) error {
	return s.r.Delete(ctx, projectID, documentID)
}
