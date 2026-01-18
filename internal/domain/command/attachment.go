package command

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type UploadAttachmentCommand struct {
	UserID  uuid.UUID
	ModelID string
	Model   string
	Files   []*multipart.FileHeader
}

type UpdateAttachmentCommand struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	OriginalName string
}

func (c *UpdateAttachmentCommand) Apply(att *entity.Attachment) {
	att.OriginalName = c.OriginalName
}
