package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
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
	ClientID string    `uri:"clientId" binding:"required"`
	CreateClientReq
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

func (*ClientRes) TableName() string {
	return "inri_clients"
}

func (r *CreateClientReq) ToCommand() *command.CreateClientCommand {
	cmd := &command.CreateClientCommand{}
	cmd.UserID = r.UserID
	cmd.ClientType = r.ClientType
	cmd.ContractType = r.ContractType
	cmd.Name = r.Name
	cmd.Description = r.Description
	cmd.TaxNumber = r.TaxNumber
	cmd.Address = r.Address
	cmd.Tags = r.Tags
	cmd.HourRate = r.HourRate
	cmd.StartedAt = r.StartedAt
	cmd.FinishedAt = r.FinishedAt
	return cmd
}

func ToClientResponse(client *entity.Client) *ClientRes {
	return &ClientRes{
		ID:           client.ID,
		ClientType:   string(client.ClientType),
		ContractType: string(client.ContractType),
		Name:         client.Name,
		Description:  client.Description,
		TaxNumber:    client.TaxNumber,
		Address:      client.Address,
		Tags:         client.Tags,
		HourRate:     client.HourRate,
		StartedAt:    client.StartedAt,
		FinishedAt:   client.FinishedAt,
		CreatedAt:    client.CreatedAt,
		UpdatedAt:    client.UpdatedAt,
	}
}
