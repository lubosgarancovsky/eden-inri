package project_document

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

var _ ports.DeleteProjectDocumentUseCase = (*DeleteProjectDocumentService)(nil)

type DeleteProjectDocumentService struct {
	repository ports.PersistProjectDocumentPort
	tm         ports.TransactionManager
}

func NewDeleteProjectDocumentService(repository ports.PersistProjectDocumentPort, tm ports.TransactionManager) *DeleteProjectDocumentService {
	return &DeleteProjectDocumentService{repository: repository, tm: tm}
}

func (s *DeleteProjectDocumentService) Execute(ctx context.Context, projectID, documentID uuid.UUID) error {
	return s.tm.WithTransaction(ctx, func(ctx context.Context) error {
		return s.repository.Delete(ctx, projectID, documentID)
	})
}
