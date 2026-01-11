package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListAttachmentsUseCase interface {
	Execute(ctx context.Context, query *query.ListQuery) (*[]entity.Attachment, int64, error)
}

type FindAttachmentByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindByIDQuery) (*entity.Attachment, error)
}

type DeleteAttachmentUseCase interface {
	Execute(ctx context.Context, userID, attachmentID string) error
}

type FindAttachmentsByModelUseCase interface {
	Execute(ctx context.Context, query *query.FindAttachmentsByModelQuery) ([]entity.Attachment, error)
}
