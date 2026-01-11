package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type PersistContactPersonPort interface {
	Create(ctx context.Context, cp *entity.ContactPerson) error
	Update(ctx context.Context, cp *entity.ContactPerson) error
	Delete(ctx context.Context, userID, clientID, contactPersonID uuid.UUID) error
	FindByID(ctx context.Context, userID, clientID, contactPersonID uuid.UUID) (*entity.ContactPerson, error)
	List(ctx context.Context, userID, clientID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ContactPerson, int64, error)
}
