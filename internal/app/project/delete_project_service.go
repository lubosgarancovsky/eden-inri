package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

type DeleteProjectService struct {
	repo ports.PersistProjectPort
}

func NewDeleteProjectService(repo ports.PersistProjectPort) *DeleteProjectService {
	return &DeleteProjectService{
		repo: repo,
	}
}

func (s *DeleteProjectService) Execute(ctx context.Context, cmd *command.DeleteCommand) error {
	return s.repo.Delete(ctx, cmd.ID)
}
