package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateProjectDocumentUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateProjectDocumentCommand) (*entity.ProjectDocument, error)
}

type UpdateProjectDocumentUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateProjectDocumentCommand) (*entity.ProjectDocument, error)
}

type DeleteProjectDocumentUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteProjectDocumentCommand) error
}

type FindProjectDocumentByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindProjectDocumentByIDQuery) (*entity.ProjectDocument, error)
}

type ListProjectDocumentsUseCase interface {
	Execute(ctx context.Context, query *query.ListProjectDocumentsQuery) (*[]entity.ProjectDocument, int64, error)
}
