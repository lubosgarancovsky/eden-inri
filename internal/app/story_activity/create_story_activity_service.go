package story_activity

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateStoryActivityService struct {
	repo ports.PersistStoryActivityPort
}

func NewCreateStoryActivityService(repo ports.PersistStoryActivityPort) *CreateStoryActivityService {
	return &CreateStoryActivityService{repo: repo}
}

func (s *CreateStoryActivityService) Execute(ctx context.Context, cmd *command.CreateStoryActivityCommand) (*entity.StoryActivity, error) {
	activity := cmd.ToDomain()
	if err := s.repo.Create(ctx, activity); err != nil {
		return nil, err
	}
	return activity, nil
}
