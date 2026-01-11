package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateProjectLabelUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateProjectLabelCommand) (*entity.ProjectLabel, error)
}

type UpdateProjectLabelUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateProjectLabelCommand) (*entity.ProjectLabel, error)
}

type DeleteProjectLabelUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteProjectLabelCommand) error
}

type FindProjectLabelByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindProjectLabelByIDQuery) (*entity.ProjectLabel, error)
}

type ListProjectLabelsUseCase interface {
	Execute(ctx context.Context, query *query.ListProjectLabelsQuery) (*[]entity.ProjectLabel, int64, error)
}
