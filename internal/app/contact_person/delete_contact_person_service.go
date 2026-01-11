package contact_person

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

type DeleteContactPersonService struct {
	repo ports.PersistContactPersonPort
}

func NewDeleteContactPersonService(repo ports.PersistContactPersonPort) *DeleteContactPersonService {
	return &DeleteContactPersonService{repo: repo}
}

func (s *DeleteContactPersonService) Execute(ctx context.Context, cmd *command.DeleteContactPersonCommand) error {
	return s.repo.Delete(ctx, cmd.UserID, cmd.ClientID, cmd.ID)
}
