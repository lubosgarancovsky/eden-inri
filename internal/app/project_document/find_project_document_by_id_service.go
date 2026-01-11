package project_document

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindProjectDocumentByIDService struct {
	repo ports.PersistProjectDocumentPort
}

func NewFindProjectDocumentByIDService(repo ports.PersistProjectDocumentPort) *FindProjectDocumentByIDService {
	return &FindProjectDocumentByIDService{repo: repo}
}

func (s *FindProjectDocumentByIDService) Execute(ctx context.Context, q *query.FindByIDProjectScopedQuery) (*entity.ProjectDocument, error) {
	return s.repo.FindByID(ctx, q.ProjectID, q.ID)
}
