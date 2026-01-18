package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type FavouriteProjectService struct {
	projectUserRepo ports.PersistProjectUserPort
	projectRepo     ports.PersistProjectPort
}

func NewFavouriteProjectService(projectRepo ports.PersistProjectPort, projectUserRepo ports.PersistProjectUserPort) *FavouriteProjectService {
	return &FavouriteProjectService{
		projectRepo:     projectRepo,
		projectUserRepo: projectUserRepo,
	}
}

func (s *FavouriteProjectService) Execute(ctx context.Context, cmd *command.Command) (*entity.Project, error) {
	err := s.projectUserRepo.Favourite(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return nil, err
	}

	return s.projectRepo.FindByID(ctx, cmd.UserID, cmd.ID)
}
