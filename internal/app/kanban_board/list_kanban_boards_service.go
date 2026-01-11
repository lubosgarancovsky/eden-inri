package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type ListKanbanBoardsService struct {
	repo         ports.PersistKanbanBoardPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewListKanbanBoardsService(repo ports.PersistKanbanBoardPort, isMemberRepo ports.IsProjectMemberPort) *ListKanbanBoardsService {
	return &ListKanbanBoardsService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *ListKanbanBoardsService) Execute(ctx context.Context, q *query.ListKanbanBoardQuery) (*[]entity.KanbanBoard, int64, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, q.UserID, q.ProjectID)
	if err != nil {
		return nil, 0, err
	}

	if !isMember {
		return nil, 0, app_err.ErrNotAMember
	}

	return s.repo.List(ctx, q.ProjectID, q.ListingQuery)
}
