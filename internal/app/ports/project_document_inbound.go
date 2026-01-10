package ports

import (
	"context"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type CreateProjectDocumentUseCase interface {
	Execute(ctx context.Context, projectID uuid.UUID, doc *entity.ProjectDocument) (*entity.ProjectDocument, error)
}

type UpdateProjectDocumentUseCase interface {
	Execute(ctx context.Context, projectID, documentID uuid.UUID, doc *entity.ProjectDocument) (*entity.ProjectDocument, error)
}

type DeleteProjectDocumentUseCase interface {
	Execute(ctx context.Context, projectID, documentID uuid.UUID) error
}

type FindProjectDocumentByIDUseCase interface {
	Execute(ctx context.Context, projectID, documentID uuid.UUID) (*entity.ProjectDocument, error)
}

type ListProjectDocumentsUseCase interface {
	Execute(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ProjectDocument, int64, error)
}
