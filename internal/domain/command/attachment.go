package command

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type UploadAttachmentCommand struct {
	UserID  uuid.UUID
	ModelID string
	Model   string
	Files   []*multipart.FileHeader
}
