package contact_person

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type UpdateContactPersonService struct {
	repo ports.PersistContactPersonPort
}

func NewUpdateContactPersonService(repo ports.PersistContactPersonPort) *UpdateContactPersonService {
	return &UpdateContactPersonService{repo: repo}
}

func (s *UpdateContactPersonService) Execute(ctx context.Context, cmd *command.UpdateContactPersonCommand) (*entity.ContactPerson, error) {
	cp, err := s.repo.FindByID(ctx, cmd.UserID, cmd.ClientID, cmd.ID)
	if err != nil {
		return nil, err
	}

	cmd.Apply(cp)

	if err := s.repo.Update(ctx, cp); err != nil {
		return nil, err
	}

	return cp, nil
}
