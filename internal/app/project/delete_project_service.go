package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/go-kit"
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
	project, err := s.repo.FindByID(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return err
	}

	if !project.IsOwner() {
		return go_kit.ErrForbidden.WithMessage("only owner can delete project")
	}

	return s.repo.Delete(ctx, cmd.ID)
}
