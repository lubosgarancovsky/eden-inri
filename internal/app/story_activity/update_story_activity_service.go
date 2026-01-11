package story_activity

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
)

type UpdateStoryActivityService struct {
	repo ports.PersistStoryActivityPort
}

func NewUpdateStoryActivityService(repo ports.PersistStoryActivityPort) *UpdateStoryActivityService {
	return &UpdateStoryActivityService{repo: repo}
}

func (s *UpdateStoryActivityService) Execute(ctx context.Context, cmd *command.UpdateStoryActivityCommand) (*entity.StoryActivity, error) {
	activity, err := s.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	// Only actor can update their own activity (e.g. comment)
	if activity.ActorID != cmd.ActorID {
		return nil, go_kit.ErrForbidden
	}

	cmd.Apply(activity)

	if err := s.repo.Update(ctx, activity); err != nil {
		return nil, err
	}

	return activity, nil
}
