package kanban_board

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_error "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
)

type FindKanbanBoardByIDService struct {
	repo         ports.PersistKanbanBoardPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewFindKanbanBoardByIDService(repo ports.PersistKanbanBoardPort, isMemberRepo ports.IsProjectMemberPort) *FindKanbanBoardByIDService {
	return &FindKanbanBoardByIDService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *FindKanbanBoardByIDService) Execute(ctx context.Context, q *query.FindByIDKanbanBoardQuery) (*entity.KanbanBoard, error) {
	isMember, err := s.isMemberRepo.IsMember(ctx, q.UserID, q.ProjectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, app_error.ErrNotAMember
	}

	return s.repo.FindByID(ctx, q.ProjectID, q.ID)
}
