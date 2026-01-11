package attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindAttachmentByIDService struct {
	repo ports.PersistAttachmentPort
}

func NewFindAttachmentByIDService(repo ports.PersistAttachmentPort) *FindAttachmentByIDService {
	return &FindAttachmentByIDService{repo: repo}
}

func (s *FindAttachmentByIDService) Execute(ctx context.Context, query *query.FindByIDQuery) (*entity.Attachment, error) {
	return s.repo.FindByID(ctx, query.UserID, query.ID)
}
