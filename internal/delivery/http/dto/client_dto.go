package dto

import (
	"time"
)

type CreateClientReq struct {
	ClientType   string     `json:"clientType" binding:"required"`
	ContractType string     `json:"contractType" binding:"required"`
	Name         string     `json:"name" binding:"required,min=3,max=100"`
	Description  string     `json:"description"`
	TaxNumber    string     `json:"taxNumber"`
	Address      string     `json:"address" binding:"required"`
	Tags         []string   `json:"tags"`
	HourRate     float64    `json:"hourRate" binding:"required,gte=0"`
	StartedAt    time.Time  `json:"startedAt" binding:"required"`
	FinishedAt   *time.Time `json:"finishedAt"`
}

type UpdateClientReq struct {
	ClientID string `uri:"clientId" binding:"required"`
	CreateClientReq
}

type ClientRes struct {
	ID           string     `json:"id"`
	ClientType   string     `json:"clientType"`
	ContractType string     `json:"contractType"`
	Description  string     `json:"description"`
	Name         string     `json:"name"`
	TaxNumber    string     `json:"taxNumber"`
	Address      string     `json:"address"`
	Tags         []string   `json:"tags"`
	HourRate     float64    `json:"hourRate"`
	StartedAt    time.Time  `json:"startedAt"`
	FinishedAt   *time.Time `json:"finishedAt"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}
