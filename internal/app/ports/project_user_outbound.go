package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistProjectUserPort interface {
	List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ProjectUser, int64, error)
	FindByID(ctx context.Context, projectID, memberID uuid.UUID) (*entity.ProjectUser, error)
	Delete(ctx context.Context, projectID, memberID uuid.UUID) error
	SetRole(ctx context.Context, projectID, memberID uuid.UUID, role entity.ProjectRole) error
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
