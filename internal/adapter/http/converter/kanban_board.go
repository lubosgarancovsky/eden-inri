package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateKanbanBoardCommand(input *dto.CreateKanbanBoardReq) (*command.CreateKanbanBoardCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateKanbanBoardCommand{
		UserID:    userID,
		ProjectID: projectID,
		Name:      input.Name,
		Status:    input.Status,
	}, nil
}

func ToUpdateKanbanBoardCommand(input *dto.UpdateKanbanBoardReq) (*command.UpdateKanbanBoardCommand, error) {
	id, err := uuid.Parse(input.BoardID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateKanbanBoardCommand{
		ID:        id,
		UserID:    userID,
		ProjectID: projectID,
		Name:      input.Name,
		Status:    input.Status,
	}, nil
}

func ToDeleteKanbanBoardCommand(input *dto.DeleteKanbanBoardReq) (*command.DeleteKanbanBoardCommand, error) {
	id, err := uuid.Parse(input.BoardID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.DeleteKanbanBoardCommand{
		ID:        id,
		ProjectID: projectID,
		UserID:    userID,
	}, nil
}

func ToListKanbanBoardQuery(input *dto.ListKanbanBoardsReq, lq *go_kit.ListingQuery) (*query.ListKanbanBoardQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.ListKanbanBoardQuery{
		UserID:       userID,
		ProjectID:    projectID,
		ListingQuery: lq,
	}, nil
}

func ToFindKanbanBoardByIDQuery(input *dto.FindKanbanBoardByIDReq) (*query.FindByIDKanbanBoardQuery, error) {
	id, err := uuid.Parse(input.BoardID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.FindByIDKanbanBoardQuery{
		ID:        id,
		ProjectID: projectID,
		UserID:    userID,
	}, nil
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
