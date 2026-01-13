package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToUpdateProjectAttachmentCommand(input *dto.UpdateProjectAttachmentReq) (*command.UpdateProjectAttachmentCommand, error) {
	AttachmentID, err := uuid.Parse(input.AttachmentID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	UserID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	ProjectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.UpdateProjectAttachmentCommand{
		AttachmentID: AttachmentID,
		UserID:       UserID,
		ProjectID:    ProjectID,
		Name:         input.Name,
	}, nil
}
