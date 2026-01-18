package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ContactPersonFromDomain(entity *entity.ContactPerson) *model.ContactPerson {
	return &model.ContactPerson{
		ID:        entity.ID,
		UserID:    entity.UserID,
		ClientID:  entity.ClientID,
		Name:      entity.Name,
		Email:     entity.Email,
		Phone:     entity.Phone,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
