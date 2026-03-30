package converter

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
)

func ToUploadAttachmentCommand(req *dto.UploadAttachmentReq, files []*multipart.FileHeader) (*command.UploadAttachmentCommand, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UploadAttachmentCommand{
		UserID:  userID,
		ModelID: req.ModelID,
		Model:   req.Model,
		Files:   files,
	}, nil
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

func ToUpdateAttachmentCommand(req *dto.UpdateAttachmentReq) (*command.UpdateAttachmentCommand, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateAttachmentCommand{
		ID:           id,
		UserID:       userID,
		OriginalName: req.Name,
	}, nil
}
