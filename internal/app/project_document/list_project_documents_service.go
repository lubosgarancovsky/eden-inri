package project_document

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListProjectDocumentsService struct {
	repo ports.PersistProjectDocumentPort
}

func NewListProjectDocumentsService(repo ports.PersistProjectDocumentPort) *ListProjectDocumentsService {
	return &ListProjectDocumentsService{repo: repo}
}

func (s *ListProjectDocumentsService) Execute(ctx context.Context, q *query.ListProjectScopedQuery) (*[]entity.ProjectDocument, int64, error) {
	return s.repo.List(ctx, q.ProjectID, q.ListingQuery)
}
