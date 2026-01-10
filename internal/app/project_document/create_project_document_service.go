package project_document

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.CreateProjectDocumentUseCase = (*CreateProjectDocumentService)(nil)

type CreateProjectDocumentService struct {
	repository ports.PersistProjectDocumentPort
	tm         ports.TransactionManager
}

func NewCreateProjectDocumentService(repository ports.PersistProjectDocumentPort, tm ports.TransactionManager) *CreateProjectDocumentService {
	return &CreateProjectDocumentService{repository: repository, tm: tm}
}

func (s *CreateProjectDocumentService) Execute(ctx context.Context, _projectID uuid.UUID, doc *entity.ProjectDocument) (*entity.ProjectDocument, error) {
	if err := s.tm.WithTransaction(ctx, func(ctx context.Context) error {
		return s.repository.Create(ctx, doc)
	}); err != nil {
		return nil, err
	}
	return doc, nil
}
