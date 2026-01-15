package project_document

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type CreateProjectDocumentService struct {
	repo         ports.PersistProjectDocumentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewCreateProjectDocumentService(repo ports.PersistProjectDocumentPort, isMemberRepo ports.IsProjectMemberPort) *CreateProjectDocumentService {
	return &CreateProjectDocumentService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *CreateProjectDocumentService) Execute(ctx context.Context, cmd *command.CreateProjectDocumentCommand) (*entity.ProjectDocument, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, cmd.UserID, cmd.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_error.ErrNotAMember
	}

	doc := cmd.ToDomain()
	doc.CreatedBy = &entity.User{ID: cmd.UserID}
	if err := s.repo.Create(ctx, doc); err != nil {
		return nil, err
	}

	return s.repo.FindByID(ctx, cmd.ProjectID, doc.ID)
}
