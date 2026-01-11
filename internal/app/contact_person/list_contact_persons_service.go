package contact_person

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListContactPersonsService struct {
	repo ports.PersistContactPersonPort
}

func NewListContactPersonsService(repo ports.PersistContactPersonPort) *ListContactPersonsService {
	return &ListContactPersonsService{repo: repo}
}

func (s *ListContactPersonsService) Execute(ctx context.Context, query *query.ListContactPersonsQuery) (*[]entity.ContactPerson, int64, error) {
	return s.repo.List(ctx, query.UserID, query.ClientID, query.ListingQuery)
}
