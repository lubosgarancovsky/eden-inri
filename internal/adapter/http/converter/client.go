package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ToCreateClientCommand(input *dto.CreateClientReq) *command.CreateClientCommand {
	return &command.CreateClientCommand{
		UserID:       uuid.MustParse(input.UserID),
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
	}
}

func ToUpdateClientCommand(input *dto.UpdateClientReq) *command.UpdateClientCommand {
	return &command.UpdateClientCommand{
		ID: uuid.MustParse(input.ClientID),
		CreateClientCommand: command.CreateClientCommand{
			UserID:       uuid.MustParse(input.UserID),
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
	}
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
