package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ProjectDocumentService struct {
	r                  *repositories.ProjectDocumentRepository
	projectUserService *ProjectUserService
}

func NewProjectDocumentService(r *repositories.ProjectDocumentRepository, pus *ProjectUserService) *ProjectDocumentService {
	return &ProjectDocumentService{r: r, projectUserService: pus}
}

func (s *ProjectDocumentService) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[models.ProjectDocument], error) {
	items, totalCount, err := s.r.FindAll(ctx, projectID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[models.ProjectDocument]{
		Items:      *items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ProjectDocumentService) FindByID(ctx context.Context, projectID uuid.UUID, documentID uuid.UUID) (*models.ProjectDocument, error) {
	return s.r.FindByID(ctx, projectID, documentID)
}

func (s *ProjectDocumentService) Create(ctx context.Context, projectID uuid.UUID, input *models.ProjectDocumentRequest) (*models.ProjectDocument, error) {
	doc := &models.ProjectDocument{
		ProjectID: projectID,
		Name:      input.Name,
		Content:   input.Content,
		Tags:      input.Tags,
	}
	return s.r.Insert(ctx, doc)
}

func (s *ProjectDocumentService) Update(ctx context.Context, projectID uuid.UUID, id uuid.UUID, input *models.ProjectDocumentRequest) (*models.ProjectDocument, error) {
	doc := &models.ProjectDocument{
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
