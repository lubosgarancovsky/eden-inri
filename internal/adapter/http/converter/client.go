package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateClientCommand(input *dto.CreateClientReq) (*command.CreateClientCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateClientCommand{
		UserID:       userID,
		ClientType:   input.ClientType,
		ContractType: input.ContractType,
		Name:         input.Name,
		Description:  input.Description,
		TaxNumber:    input.TaxNumber,
		Address:      input.Address,
		Tags:         input.Tags,
		HourRate:     input.HourRate,
		StartedAt:    input.StartedAt,
		FinishedAt:   input.FinishedAt,
	}, nil
}

func ToUpdateClientCommand(input *dto.UpdateClientReq) (*command.UpdateClientCommand, error) {
	id, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateClientCommand{
		ID: id,
		CreateClientCommand: command.CreateClientCommand{
			UserID:       userID,
			ClientType:   input.ClientType,
			ContractType: input.ContractType,
			Name:         input.Name,
			Description:  input.Description,
			TaxNumber:    input.TaxNumber,
			Address:      input.Address,
			Tags:         input.Tags,
			HourRate:     input.HourRate,
			StartedAt:    input.StartedAt,
			FinishedAt:   input.FinishedAt,
		},
	}, nil
}

func ToClientResponse(client *entity.Client) *dto.ClientRes {
	return &dto.ClientRes{
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
