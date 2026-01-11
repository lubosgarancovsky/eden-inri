package contact_person

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindContactPersonByIDService struct {
	repo ports.PersistContactPersonPort
}

func NewFindContactPersonByIDService(repo ports.PersistContactPersonPort) *FindContactPersonByIDService {
	return &FindContactPersonByIDService{repo: repo}
}

func (s *FindContactPersonByIDService) Execute(ctx context.Context, query *query.FindContactPersonByIDQuery) (*entity.ContactPerson, error) {
	return s.repo.FindByID(ctx, query.UserID, query.ClientID, query.ID)
}
