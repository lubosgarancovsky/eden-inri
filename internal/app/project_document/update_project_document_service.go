package project_document

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type UpdateProjectDocumentService struct {
	repo         ports.PersistProjectDocumentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewUpdateProjectDocumentService(repo ports.PersistProjectDocumentPort, isMemberRepo ports.IsProjectMemberPort) *UpdateProjectDocumentService {
	return &UpdateProjectDocumentService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *UpdateProjectDocumentService) Execute(ctx context.Context, cmd *command.UpdateProjectDocumentCommand) (*entity.ProjectDocument, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, cmd.UserID, cmd.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_error.ErrNotAMember
	}

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
