package project_document

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

var _ ports.UpdateProjectDocumentUseCase = (*UpdateProjectDocumentService)(nil)

type UpdateProjectDocumentService struct {
	repository ports.PersistProjectDocumentPort
	tm         ports.TransactionManager
}

func NewUpdateProjectDocumentService(repository ports.PersistProjectDocumentPort, tm ports.TransactionManager) *UpdateProjectDocumentService {
	return &UpdateProjectDocumentService{repository: repository, tm: tm}
}

func (s *UpdateProjectDocumentService) Execute(ctx context.Context, _projectID, _documentID uuid.UUID, doc *entity.ProjectDocument) (*entity.ProjectDocument, error) {
	if err := s.tm.WithTransaction(ctx, func(ctx context.Context) error {
		return s.repository.Update(ctx, doc)
	}); err != nil {
		return nil, err
	}
	return doc, nil
}
