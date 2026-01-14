package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type PersistStoryLabelPort interface {
	List(ctx context.Context, storyID uuid.UUID) (*[]entity.StoryLabel, error)
	Assign(ctx context.Context, storyID, labelID uuid.UUID) error
	Unassign(ctx context.Context, storyID, labelID uuid.UUID) error
}

type ListStoryLabelsUseCase interface {
	Execute(ctx context.Context, q *query.ListStoryLabelsQuery) (*[]entity.StoryLabel, error)
}

type AssignStoryLabelUseCase interface {
	Execute(ctx context.Context, cmd *command.StoryLabelCommand) error
}

type UnassignStoryLabelUseCase interface {
	Execute(ctx context.Context, cmd *command.StoryLabelCommand) error
}
