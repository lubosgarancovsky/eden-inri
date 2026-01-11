package project_document

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type UpdateProjectDocumentService struct {
	repo ports.PersistProjectDocumentPort
}

func NewUpdateProjectDocumentService(repo ports.PersistProjectDocumentPort) *UpdateProjectDocumentService {
	return &UpdateProjectDocumentService{repo: repo}
}

func (s *UpdateProjectDocumentService) Execute(ctx context.Context, cmd *command.UpdateProjectDocumentCommand) (*entity.ProjectDocument, error) {
	doc, err := s.repo.FindByID(ctx, cmd.ProjectID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(doc)

	if err := s.repo.Update(ctx, doc); err != nil {
		return nil, err
	}

	return doc, nil
}
