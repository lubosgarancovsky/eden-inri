package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateClientReq struct {
	UserID       uuid.UUID  `header:"X-User-ID" binding:"required"`
	ClientType   string     `json:"clientType" binding:"required"`
	ContractType string     `json:"contractType" binding:"required"`
	Name         string     `json:"name" binding:"required,min=3,max=100"`
	Description  *string    `json:"description"`
	TaxNumber    *string    `json:"taxNumber"`
	Address      *string    `json:"address" binding:"required"`
	Tags         []string   `json:"tags"`
	HourRate     float64    `json:"hourRate" binding:"required,gte=0"`
	StartedAt    time.Time  `json:"startedAt" binding:"required"`
	FinishedAt   *time.Time `json:"finishedAt"`
}

type UpdateClientReq struct {
	UserID   uuid.UUID `header:"X-User-ID"`
	ClientID uuid.UUID `uri:"clientId" binding:"required"`
	CreateClientReq
}

type DeleteClientReq struct {
	UserID   uuid.UUID `header:"X-User-ID"`
	ClientID uuid.UUID `uri:"clientId" binding:"required"`
}

type FindClientByIDReq struct {
	UserID uuid.UUID `header:"X-User-ID"`
	ID     uuid.UUID `uri:"clientId" binding:"required"`
}

type ListClientsReq struct {
	UserID uuid.UUID `header:"X-User-ID"`
}

type ClientRes struct {
	ID           uuid.UUID  `json:"id"`
	ClientType   string     `json:"clientType"`
	ContractType string     `json:"contractType"`
	Name         string     `json:"name"`
	Description  *string    `json:"description"`
	TaxNumber    *string    `json:"taxNumber"`
	Address      *string    `json:"address"`
	Tags         []string   `json:"tags"`
	HourRate     float64    `json:"hourRate"`
	StartedAt    time.Time  `json:"startedAt"`
	FinishedAt   *time.Time `json:"finishedAt"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}
