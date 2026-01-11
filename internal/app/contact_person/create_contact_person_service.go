package contact_person

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateContactPersonService struct {
	repo ports.PersistContactPersonPort
}

func NewCreateContactPersonService(repo ports.PersistContactPersonPort) *CreateContactPersonService {
	return &CreateContactPersonService{repo: repo}
}

func (s *CreateContactPersonService) Execute(ctx context.Context, cmd *command.CreateContactPersonCommand) (*entity.ContactPerson, error) {
	cp := cmd.ToDomain()
	if err := s.repo.Create(ctx, cp); err != nil {
		return nil, err
	}
	return cp, nil
}
