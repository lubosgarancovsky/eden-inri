package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type PersistProjectUserPort interface {
	Insert(ctx context.Context, pu *entity.ProjectUser) error
	Favourite(ctx context.Context, userID, projectID uuid.UUID) error
}

type IsProjectMemberPort interface {
	IsMember(ctx context.Context, userID, projectID uuid.UUID) (bool, error)
}

type MemberHasRolePort interface {
	HasRole(ctx context.Context, userID, projectID uuid.UUID, roles []entity.ProjectRole) (bool, error)
}

type GetProjectOwnerPort interface {
	GetOwner(ctx context.Context, projectID uuid.UUID) (uuid.UUID, error)
}
