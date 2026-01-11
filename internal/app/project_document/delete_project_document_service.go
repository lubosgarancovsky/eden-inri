package project_document

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type DeleteProjectDocumentService struct {
	repo         ports.PersistProjectDocumentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewDeleteProjectDocumentService(repo ports.PersistProjectDocumentPort, isMemberRepo ports.IsProjectMemberPort) *DeleteProjectDocumentService {
	return &DeleteProjectDocumentService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *DeleteProjectDocumentService) Execute(ctx context.Context, cmd *command.DeleteProjectDocumentCommand) error {
	isMember, err := s.isMemberRepo.IsMember(ctx, cmd.UserID, cmd.ProjectID)
	if err != nil {
		return err
	}

	if !isMember {
		return app_error.ErrNotAMember
	}

	return s.repo.Delete(ctx, cmd.ProjectID, cmd.ID)
}
