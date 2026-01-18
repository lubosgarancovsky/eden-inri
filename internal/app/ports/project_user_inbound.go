package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListProjectUsersUseCase interface {
	Execute(ctx context.Context, query *query.ScopedListQuery) (*[]entity.ProjectUser, int64, error)
}

type DeleteProjectUserUseCase interface {
	Execute(ctx context.Context, cmd *command.ScopedCommand) error
}

type ChangeProjectUserRoleUseCase interface {
	Execute(ctx context.Context, cmd *command.ChangeProjectUserRoleCommand) (*entity.ProjectUser, error)
}
