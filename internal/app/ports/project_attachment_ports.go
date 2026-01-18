package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListProjectAttachmentsUseCase interface {
	Execute(ctx context.Context, query *query.ScopedListQuery) (*[]entity.Attachment, int64, error)
}

type FindProjectAttachmentByIDUseCase interface {
	Execute(ctx context.Context, query *query.ScopedQuery) (*entity.Attachment, error)
}

type DeleteProjectAttachmentUseCase interface {
	Execute(ctx context.Context, cmd *command.ScopedCommand) error
}

type UpdateProjectAttachmentUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateProjectAttachmentCommand) (*entity.Attachment, error)
}
