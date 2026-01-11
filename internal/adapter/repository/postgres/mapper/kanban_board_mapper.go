package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func KanbanBoardFromDomain(e *entity.KanbanBoard) *model.KanbanBoard {
	return &model.KanbanBoard{
		ID:             e.ID,
		Name:           e.Name,
		Status:         e.Status,
		ProjectID:      e.ProjectID,
		CreatedAt:      e.CreatedAt,
		LastActivityAt: e.LastActivityAt,
	}
}
