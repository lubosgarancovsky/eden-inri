package kanban_column

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type CreateKanbanColumnService struct {
	repo        ports.PersistKanbanColumnPort
	hasRoleRepo ports.MemberHasRolePort
}

func NewCreateKanbanColumnService(repo ports.PersistKanbanColumnPort, hasRoleRepo ports.MemberHasRolePort) *CreateKanbanColumnService {
	return &CreateKanbanColumnService{repo: repo, hasRoleRepo: hasRoleRepo}
}

func (s *CreateKanbanColumnService) Execute(ctx context.Context, cmd *command.CreateKanbanColumnCommand) (*entity.KanbanColumn, error) {
	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, cmd.UserID, cmd.ProjectID, roles)
	if err != nil {
		return nil, err
	}

	if !hasRole {
		return nil, app_err.ErrNotAMember
	}

	col := cmd.ToDomain()
	if err = s.repo.Create(ctx, col); err != nil {
		return nil, err
	}
	return col, nil
}
