package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateStoryUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateStoryCommand) (*entity.Story, error)
}

type UpdateStoryUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateStoryCommand) (*entity.Story, error)
}

type DeleteStoryUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteStoryCommand) error
}

type FindStoryByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindStoryByIDQuery) (*entity.Story, error)
}

type ListStoriesUseCase interface {
	Execute(ctx context.Context, query *query.ListStoriesQuery) (*[]entity.Story, int64, error)
}

type ListAssignedStoriesUseCase interface {
	Execute(ctx context.Context, query *query.ListAssignedStoriesQuery) (*[]entity.Story, int64, error)
}

type ChangeStoryAssigneeUseCase interface {
	Execute(ctx context.Context, cmd *command.ChangeStoryAssigneeCommand) error
}
