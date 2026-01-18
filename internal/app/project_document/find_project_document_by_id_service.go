package project_document

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindProjectDocumentByIDService struct {
	repo         ports.PersistProjectDocumentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewFindProjectDocumentByIDService(repo ports.PersistProjectDocumentPort, isMemberRepo ports.IsProjectMemberPort) *FindProjectDocumentByIDService {
	return &FindProjectDocumentByIDService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *FindProjectDocumentByIDService) Execute(ctx context.Context, q *query.FindProjectDocumentByIDQuery) (*entity.ProjectDocument, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, q.UserID, q.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_error.ErrNotAMember
	}

	return s.repo.FindByID(ctx, q.ProjectID, q.DocumentID)
}
