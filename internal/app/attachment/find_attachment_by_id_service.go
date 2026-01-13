package attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindAttachmentByIDService struct {
	repo ports.PersistAttachmentPort
}

func NewFindAttachmentByIDService(repo ports.PersistAttachmentPort) *FindAttachmentByIDService {
	return &FindAttachmentByIDService{repo: repo}
}

func (s *FindAttachmentByIDService) Execute(ctx context.Context, query *query.Query) (*entity.Attachment, error) {
	attachment, err := s.repo.FindByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	if attachment.UserID != query.UserID {
		return nil, app_err.ErrInsufficientPermission
	}

	return attachment, nil
}
