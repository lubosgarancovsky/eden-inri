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
	Execute(ctx context.Context, query *query.FindByIDQuery) (*entity.Attachment, error)
}

type UploadAttachmentUseCase interface {
	Execute(ctx context.Context, cmd *command.UploadAttachmentCommand) error
}

type DeleteAttachmentUseCase interface {
	Execute(ctx context.Context, userID, attachmentID string) error
}

type FindAttachmentsByModelUseCase interface {
	Execute(ctx context.Context, query *query.FindAttachmentsByModelQuery) ([]entity.Attachment, error)
}

// -- Outbound
type PersistAttachmentPort interface {
	Create(ctx context.Context, attachment *entity.Attachment) error
	Delete(ctx context.Context, userID, attachmentID uuid.UUID) error
	FindByID(ctx context.Context, userID, attachmentID uuid.UUID) (*entity.Attachment, error)
	FindByModelID(ctx context.Context, userID, modelID uuid.UUID, modelName string) ([]entity.Attachment, error)
	List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Attachment, int64, error)
}
