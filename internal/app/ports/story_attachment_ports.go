package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListStoryAttachmentsUserCase interface {
	Execute(ctx context.Context, query *query.ListStoryAttachmentsQuery) (*[]entity.Attachment, int64, error)
}

type FindStoryAttachmentByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindStoryAttachmentByIDQuery) (*entity.Attachment, error)
}

type DeleteStoryAttachmentUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteStoryAttachmentCommand) error
}
