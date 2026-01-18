package converter

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ToUploadAttachmentCommand(req *dto.UploadAttachmentReq, files []*multipart.FileHeader) *command.UploadAttachmentCommand {
	return &command.UploadAttachmentCommand{
		UserID:  uuid.MustParse(req.UserID),
		ModelID: req.ModelID,
		Model:   req.Model,
		Files:   files,
	}
}

func ToAttachmentResponse(e *entity.Attachment) *dto.AttachmentRes {
	return &dto.AttachmentRes{
		ID:           e.ID.String(),
		Model:        e.Model,
		ModelID:      e.ModelID,
		OriginalName: e.OriginalName,
		MimeType:     e.MimeType,
		Size:         e.Size,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}

func ToUpdateAttachmentCommand(req *dto.UpdateAttachmentReq) *command.UpdateAttachmentCommand {
	return &command.UpdateAttachmentCommand{
		ID:           uuid.MustParse(req.ID),
		UserID:       uuid.MustParse(req.UserID),
		OriginalName: req.Name,
	}
}
