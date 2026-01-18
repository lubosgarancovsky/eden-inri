package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func AttachmentFromDomain(e *entity.Attachment) *model.Attachment {
	return &model.Attachment{
		ID:           e.ID,
		UserID:       e.UserID,
		Model:        e.Model,
		ModelID:      e.ModelID,
		OriginalName: e.OriginalName,
		ServerName:   e.ServerName,
		MimeType:     e.MimeType,
		Size:         e.Size,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}
