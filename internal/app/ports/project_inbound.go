package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateProjectUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateProjectCommand) (*entity.Project, error)
}

type UpdateProjectUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateProjectCommand) (*entity.Project, error)
}

type DeleteProjectUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteCommand) error
}

type FindProjectByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindByIDQuery) (*entity.Project, error)
}

type ListProjectsUseCase interface {
	Execute(ctx context.Context, query *query.ListQuery) (*[]entity.Project, int64, error)
}

type FavouriteProjectUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteCommand) (*entity.Project, error)
}
