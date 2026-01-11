package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ClientFromDomain(entity *entity.Client) *model.Client {
	client := &model.Client{}
	client.ID = entity.ID
	client.UserID = entity.UserID
	client.ClientType = string(entity.ClientType)
	client.ContractType = string(entity.ContractType)
	client.Name = entity.Name
	client.Description = entity.Description
	client.TaxNumber = entity.TaxNumber
	client.Address = entity.Address
	client.Tags = entity.Tags
	client.HourRate = entity.HourRate
	client.StartedAt = entity.StartedAt
	client.FinishedAt = entity.FinishedAt
	client.CreatedAt = entity.CreatedAt
	client.UpdatedAt = entity.UpdatedAt
	return client
}
