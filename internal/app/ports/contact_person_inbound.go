package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type CreateContactPersonUseCase interface {
	Execute(ctx context.Context, cmd *command.CreateContactPersonCommand) (*entity.ContactPerson, error)
}

type UpdateContactPersonUseCase interface {
	Execute(ctx context.Context, cmd *command.UpdateContactPersonCommand) (*entity.ContactPerson, error)
}

type DeleteContactPersonUseCase interface {
	Execute(ctx context.Context, cmd *command.DeleteContactPersonCommand) error
}

type FindContactPersonByIDUseCase interface {
	Execute(ctx context.Context, query *query.FindContactPersonByIDQuery) (*entity.ContactPerson, error)
}

type ListContactPersonsUseCase interface {
	Execute(ctx context.Context, query *query.ListContactPersonsQuery) (*[]entity.ContactPerson, int64, error)
}
