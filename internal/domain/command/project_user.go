package command

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type ChangeProjectUserRoleCommand struct {
	UserID    uuid.UUID
	ProjectID uuid.UUID
	MemberID  uuid.UUID
	Role      entity.ProjectRole
}

func (c *ChangeProjectUserRoleCommand) Apply(pu *entity.ProjectUser) {
	pu.Role = c.Role
}
