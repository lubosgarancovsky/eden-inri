package converter

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ToAttachmentResponse(e *entity.Attachment) *dto.AttachmentRes {
	return &dto.AttachmentRes{
		ID:           e.ID.String(),
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
