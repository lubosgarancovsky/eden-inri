package attachment

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindAttachmentsByModelService struct {
	repo ports.PersistAttachmentPort
}

func NewFindAttachmentsByModelService(repo ports.PersistAttachmentPort) *FindAttachmentsByModelService {
	return &FindAttachmentsByModelService{repo: repo}
}

func (s *FindAttachmentsByModelService) Execute(ctx context.Context, query *query.FindAttachmentsByModelQuery) ([]entity.Attachment, error) {
	return s.repo.FindByModelID(ctx, query.UserID, query.ModelID, query.ModelName)
}
