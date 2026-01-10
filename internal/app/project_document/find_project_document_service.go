package project_document

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.FindProjectDocumentByIDUseCase = (*FindProjectDocumentService)(nil)

type FindProjectDocumentService struct {
	repo ports.QueryProjectDocumentPort
}

func NewFindProjectDocumentService(repo ports.QueryProjectDocumentPort) *FindProjectDocumentService {
	return &FindProjectDocumentService{repo}
}

func (s *FindProjectDocumentService) Execute(ctx context.Context, projectID, documentID uuid.UUID) (*entity.ProjectDocument, error) {
	return s.repo.FindByID(ctx, projectID, documentID)
}
