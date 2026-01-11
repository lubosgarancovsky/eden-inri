package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateClientReq struct {
	UserID       string     `header:"X-User-ID"`
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
}

type UpdateClientReq struct {
	UserID   string `header:"X-User-ID"`
	ClientID string `uri:"clientId"`
	CreateClientReq
}

type DeleteClientReq struct {
	UserID string `header:"X-User-ID"`
	ID     string `uri:"clientId"`
}

type FindClientByIDReq struct {
	UserID string `header:"X-User-ID"`
	ID     string `uri:"clientId"`
}

type ListClientsReq struct {
	UserID string `header:"X-User-ID"`
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
