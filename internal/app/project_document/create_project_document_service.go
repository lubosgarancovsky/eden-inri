package project_document

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateProjectDocumentService struct {
	repo ports.PersistProjectDocumentPort
}

func NewCreateProjectDocumentService(repo ports.PersistProjectDocumentPort) *CreateProjectDocumentService {
	return &CreateProjectDocumentService{repo: repo}
}

func (s *CreateProjectDocumentService) Execute(ctx context.Context, cmd *command.CreateProjectDocumentCommand) (*entity.ProjectDocument, error) {
	doc := cmd.ToDomain()
	if err := s.repo.Create(ctx, doc); err != nil {
		return nil, err
	}
	return doc, nil
}
