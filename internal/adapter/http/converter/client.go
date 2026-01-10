package converter

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateClientCommand(input *dto.CreateClientReq) *command.CreateClientCommand {
	return &command.CreateClientCommand{
		UserID:       input.UserID,
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
		ID: input.ClientID,
		CreateClientCommand: command.CreateClientCommand{
			UserID:       input.UserID,
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

func ToDeleteClientCommand(input *dto.DeleteClientReq) *command.DeleteClientCommand {
	return &command.DeleteClientCommand{
		ID:     input.ClientID,
		UserID: input.UserID,
	}
}

func ToFindClientByIDQuery(req *dto.FindClientByIDReq) *query.FindClientByIDQuery {
	return &query.FindClientByIDQuery{
		ID:     req.ID,
		UserID: req.UserID,
	}
}

func ToListClientsQuery(req *dto.ListClientsReq, lq *go_kit.ListingQuery) *query.ListClientsQuery {
	return &query.ListClientsQuery{
		UserID:       req.UserID,
		ListingQuery: *lq,
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

func ToClientListResponse(clients *[]entity.Client) []*dto.ClientRes {
	if clients == nil {
		return nil
	}

	clientsArray := *clients

	res := make([]*dto.ClientRes, len(clientsArray))
	for i, client := range clientsArray {
		res[i] = ToClientResponse(&client)
	}
	return res
}
