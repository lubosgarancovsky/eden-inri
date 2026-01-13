package story_activity

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/go-kit"
)

type DeleteStoryActivityService struct {
	repo ports.PersistStoryActivityPort
}

func NewDeleteStoryActivityService(repo ports.PersistStoryActivityPort) *DeleteStoryActivityService {
	return &DeleteStoryActivityService{repo: repo}
}

func (s *DeleteStoryActivityService) Execute(ctx context.Context, cmd *command.Command) error {
	activity, err := s.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if activity.ActorID != cmd.UserID {
		return go_kit.ErrForbidden
	}

	return s.repo.Delete(ctx, cmd.ID)
}
