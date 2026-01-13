package command

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type UpdateProjectAttachmentCommand struct {
	UserID       uuid.UUID
	ProjectID    uuid.UUID
	AttachmentID uuid.UUID
	Name         string
}

func (c *UpdateProjectAttachmentCommand) Apply(att *entity.Attachment) {
	att.OriginalName = c.Name
}
