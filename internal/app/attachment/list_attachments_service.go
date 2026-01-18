package attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListAttachmentsService struct {
	repo ports.PersistAttachmentPort
}

func NewListAttachmentsService(repo ports.PersistAttachmentPort) *ListAttachmentsService {
	return &ListAttachmentsService{repo: repo}
}

func (s *ListAttachmentsService) Execute(ctx context.Context, query *query.ListQuery) (*[]entity.Attachment, int64, error) {
	return s.repo.List(ctx, query.UserID, query.ListingQuery)
}
