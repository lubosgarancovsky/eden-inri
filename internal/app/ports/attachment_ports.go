package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

// -- Inbound
type ListAttachmentsUseCase interface {
	Execute(ctx context.Context, query *query.ListQuery) (*[]entity.Attachment, int64, error)
}

type FindAttachmentByIDUseCase interface {
	Execute(ctx context.Context, query *query.Query) (*entity.Attachment, error)
}

type UploadAttachmentUseCase interface {
	Execute(ctx context.Context, cmd *command.UploadAttachmentCommand) error
}

type UpdateAttachmentUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateAttachmentCommand) (*entity.Attachment, error)
}

type DeleteAttachmentUseCase interface {
	Execute(ctx context.Context, cmd *command.Command) error
}

type ListAttachmentsByModelUseCase interface {
	Execute(ctx context.Context, query *query.ListAttachmentsByModelQuery) (*[]entity.Attachment, error)
}

// -- Outbound
type PersistAttachmentPort interface {
	Create(ctx context.Context, attachment *entity.Attachment) error
	Update(ctx context.Context, attachment *entity.Attachment) error
	Delete(ctx context.Context, attachmentID uuid.UUID) error
	FindByID(ctx context.Context, attachmentID uuid.UUID) (*entity.Attachment, error)
	List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Attachment, int64, error)
	ListByModelID(ctx context.Context, modelID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Attachment, int64, error)
}
