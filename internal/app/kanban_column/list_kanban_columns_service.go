package kanban_column

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListKanbanColumnsService struct {
	repo         ports.PersistKanbanColumnPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewListKanbanColumnsService(repo ports.PersistKanbanColumnPort, isMemberRepo ports.IsProjectMemberPort) *ListKanbanColumnsService {
	return &ListKanbanColumnsService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *ListKanbanColumnsService) Execute(ctx context.Context, query *query.ListKanbanColumnsQuery) ([]entity.KanbanColumn, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, query.UserID, query.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return []entity.KanbanColumn{}, app_err.ErrNotAMember
	}

	return s.repo.List(ctx, query.BoardID)
}
