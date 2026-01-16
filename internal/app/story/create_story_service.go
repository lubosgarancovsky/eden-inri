package story

import (
	"context"
	"fmt"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type CreateStoryService struct {
	storyRepo   ports.PersistStoryPort
	projectRepo ports.PersistProjectPort
	memberRepo  ports.MemberHasRolePort
	txManager   ports.TransactionManager
}

func NewCreateStoryService(
	storyRepo ports.PersistStoryPort,
	projectRepo ports.PersistProjectPort,
	memberRepo ports.MemberHasRolePort,
	txManager ports.TransactionManager,
) *CreateStoryService {
	return &CreateStoryService{
		storyRepo:   storyRepo,
		projectRepo: projectRepo,
		memberRepo:  memberRepo,
		txManager:   txManager,
	}
}

func (s *CreateStoryService) Execute(ctx context.Context, cmd *command.CreateStoryCommand) (*entity.Story, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin, entity.ProjectRoleDeveloper}
	hasRole, err := s.memberRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return nil, err
	}
	if !hasRole {
		return nil, app_error.ErrInsufficientProjectRole
	}

	story := cmd.ToDomain()

	err = s.txManager.WithTransaction(ctx, func(tx context.Context) error {
		project, err := s.projectRepo.FindByID(tx, cmd.UserID, cmd.ProjectID)
		if err != nil {
			return err
		}

		project.StorySequence++
		if err := s.projectRepo.Update(tx, project); err != nil {
			return err
		}

		story.Slug = fmt.Sprintf("%s-%d", project.Slug, project.StorySequence)

		return s.storyRepo.Create(tx, story)
	})

	if err != nil {
		return nil, err
	}

	return story, nil
}
