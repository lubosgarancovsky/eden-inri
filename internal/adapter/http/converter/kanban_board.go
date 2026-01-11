package converter

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ToCreateKanbanBoardCommand(input *dto.CreateKanbanBoardReq) *command.CreateKanbanBoardCommand {
	return &command.CreateKanbanBoardCommand{
		ProjectID: input.ProjectID,
		Name:      input.Name,
		Status:    input.Status,
	}
}

func ToUpdateKanbanBoardCommand(input *dto.UpdateKanbanBoardReq) *command.UpdateKanbanBoardCommand {
	return &command.UpdateKanbanBoardCommand{
		ID:        input.BoardID,
		ProjectID: input.ProjectID,
		Name:      input.Name,
		Status:    input.Status,
	}
}

func ToKanbanBoardResponse(e *entity.KanbanBoard) *dto.KanbanBoardRes {
	return &dto.KanbanBoardRes{
		ID:             e.ID,
		ProjectID:      e.ProjectID,
		Name:           e.Name,
		Status:         e.Status,
		CreatedAt:      e.CreatedAt,
		LastActivityAt: e.LastActivityAt,
	}
}
