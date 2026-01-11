package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateKanbanBoardReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Status    string    `json:"status"`
}

type UpdateKanbanBoardReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
	BoardID   uuid.UUID `uri:"boardId" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Status    string    `json:"status"`
}

type DeleteKanbanBoardReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
	BoardID   uuid.UUID `uri:"boardId" binding:"required"`
}

type FindKanbanBoardByIDReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
	BoardID   uuid.UUID `uri:"boardId" binding:"required"`
}

type ListKanbanBoardsReq struct {
	ProjectID uuid.UUID `uri:"projectId" binding:"required"`
}

type KanbanBoardRes struct {
	ID             uuid.UUID `json:"id"`
	ProjectID      uuid.UUID `json:"projectId"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"createdAt"`
	LastActivityAt time.Time `json:"lastActivityAt"`
}

func (r CreateKanbanBoardReq) GetID() uuid.UUID        { return uuid.Nil }
func (r CreateKanbanBoardReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r CreateKanbanBoardReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r UpdateKanbanBoardReq) GetID() uuid.UUID        { return r.BoardID }
func (r UpdateKanbanBoardReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r UpdateKanbanBoardReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r DeleteKanbanBoardReq) GetID() uuid.UUID        { return r.BoardID }
func (r DeleteKanbanBoardReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r DeleteKanbanBoardReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r FindKanbanBoardByIDReq) GetID() uuid.UUID        { return r.BoardID }
func (r FindKanbanBoardByIDReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r FindKanbanBoardByIDReq) GetProjectID() uuid.UUID { return r.ProjectID }

func (r ListKanbanBoardsReq) GetUserID() uuid.UUID    { return uuid.Nil }
func (r ListKanbanBoardsReq) GetProjectID() uuid.UUID { return r.ProjectID }
