package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateKanbanBoardCommand(input *dto.CreateKanbanBoardReq) *command.CreateKanbanBoardCommand {
	return &command.CreateKanbanBoardCommand{
		ProjectID: uuid.MustParse(input.ProjectID),
		Name:      input.Name,
		Status:    input.Status,
	}
}

func ToUpdateKanbanBoardCommand(input *dto.UpdateKanbanBoardReq) *command.UpdateKanbanBoardCommand {
	return &command.UpdateKanbanBoardCommand{
		ID:        uuid.MustParse(input.BoardID),
		ProjectID: uuid.MustParse(input.ProjectID),
		Name:      input.Name,
		Status:    input.Status,
	}
}

func ToDeleteKanbanBoardCommand(input *dto.DeleteKanbanBoardReq) *command.DeleteKanbanBoardCommand {
	return &command.DeleteKanbanBoardCommand{
		ID:        uuid.MustParse(input.BoardID),
		ProjectID: uuid.MustParse(input.ProjectID),
		UserID:    uuid.MustParse(input.UserID),
	}
}

func ToListKanbanBoardQuery(input *dto.ListKanbanBoardsReq, lq *go_kit.ListingQuery) *query.ListKanbanBoardQuery {
	return &query.ListKanbanBoardQuery{
		UserID:       uuid.MustParse(input.UserID),
		ProjectID:    uuid.MustParse(input.ProjectID),
		ListingQuery: lq,
	}
}

func ToFindKanbanBoardByIDQuery(input *dto.FindKanbanBoardByIDReq) *query.FindByIDKanbanBoardQuery {
	return &query.FindByIDKanbanBoardQuery{
		ID:        uuid.MustParse(input.BoardID),
		ProjectID: uuid.MustParse(input.ProjectID),
		UserID:    uuid.MustParse(input.UserID),
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
