package project_document

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

type DeleteProjectDocumentService struct {
	repo ports.PersistProjectDocumentPort
}

func NewDeleteProjectDocumentService(repo ports.PersistProjectDocumentPort) *DeleteProjectDocumentService {
	return &DeleteProjectDocumentService{repo: repo}
}

func (s *DeleteProjectDocumentService) Execute(ctx context.Context, cmd *command.DeleteProjectScopedCommand) error {
	return s.repo.Delete(ctx, cmd.ProjectID, cmd.ID)
}
