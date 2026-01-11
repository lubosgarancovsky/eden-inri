package command

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type SaveAttachmentCommand struct {
	UserID       uuid.UUID
	Model        string
	ModelID      string
	OriginalName string
	ServerName   string
	MimeType     string
	Size         int64
}

func (c *SaveAttachmentCommand) ToDomain() *entity.Attachment {
	return &entity.Attachment{
		ID:           uuid.New(),
		UserID:       c.UserID,
		Model:        c.Model,
		ModelID:      c.ModelID,
		OriginalName: c.OriginalName,
		ServerName:   c.ServerName,
		MimeType:     c.MimeType,
		Size:         c.Size,
	}
}
