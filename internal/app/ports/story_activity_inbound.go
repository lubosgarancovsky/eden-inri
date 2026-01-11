package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateStoryActivityUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateStoryActivityCommand) (*entity.StoryActivity, error)
}

type UpdateStoryActivityUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateStoryActivityCommand) (*entity.StoryActivity, error)
}

type DeleteStoryActivityUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteCommand) error
}

type ListStoryActivitiesUseCase interface {
	Execute(ctx context.Context, query *query.ListStoryActivitiesQuery) (*[]entity.StoryActivity, int64, error)
}
