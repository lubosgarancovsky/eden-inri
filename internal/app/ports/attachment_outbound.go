package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistAttachmentPort interface {
	Create(ctx context.Context, attachment *entity.Attachment) error
	Delete(ctx context.Context, userID, attachmentID uuid.UUID) error
	FindByID(ctx context.Context, userID, attachmentID uuid.UUID) (*entity.Attachment, error)
	FindByModelID(ctx context.Context, userID, modelID uuid.UUID, modelName string) ([]entity.Attachment, error)
	List(ctx context.Context, userID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.Attachment, int64, error)
}
