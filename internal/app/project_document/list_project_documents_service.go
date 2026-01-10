package project_document

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

var _ ports.ListProjectDocumentsUseCase = (*ListProjectDocumentsService)(nil)

type ListProjectDocumentsService struct {
	repo ports.QueryProjectDocumentPort
}

func NewListProjectDocumentsService(repo ports.QueryProjectDocumentPort) *ListProjectDocumentsService {
	return &ListProjectDocumentsService{repo}
}

func (s *ListProjectDocumentsService) Execute(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ProjectDocument, int64, error) {
	return s.repo.List(ctx, projectID, lq)
}
