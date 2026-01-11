package project

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type FavouriteProjectService struct {
	repo ports.PersistProjectPort
}

func NewFavouriteProjectService(repo ports.PersistProjectPort) *FavouriteProjectService {
	return &FavouriteProjectService{
		repo: repo,
	}
}

func (s *FavouriteProjectService) Execute(ctx context.Context, cmd *command.DeleteCommand) (*entity.Project, error) {
	_, err := s.repo.Favourite(ctx, cmd.UserID, cmd.ID)
	if err != nil {
		return nil, err
	}

	return s.repo.FindByID(ctx, cmd.UserID, cmd.ID)
}
