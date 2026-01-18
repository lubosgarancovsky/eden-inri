package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateContactPersonCommand(input *dto.CreateContactPersonReq) *command.CreateContactPersonCommand {
	return &command.CreateContactPersonCommand{
		UserID:   uuid.MustParse(input.UserID),
		ClientID: uuid.MustParse(input.ClientID),
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
	}
}

func ToUpdateContactPersonCommand(input *dto.UpdateContactPersonReq) *command.UpdateContactPersonCommand {
	return &command.UpdateContactPersonCommand{
		ID: uuid.MustParse(input.ContactPersonID),
		CreateContactPersonCommand: command.CreateContactPersonCommand{
			UserID:   uuid.MustParse(input.UserID),
			ClientID: uuid.MustParse(input.ClientID),
			Name:     input.Name,
			Email:    input.Email,
			Phone:    input.Phone,
		},
	}
}

func ToDeleteContactPersonCommand(input *dto.DeleteContactPersonReq) *command.DeleteContactPersonCommand {
	return &command.DeleteContactPersonCommand{
		ID:       uuid.MustParse(input.ContactPersonID),
		UserID:   uuid.MustParse(input.UserID),
		ClientID: uuid.MustParse(input.ClientID),
	}
}

func ToFindContactPersonByIDQuery(input *dto.FindContactPersonByIDReq) *query.FindContactPersonByIDQuery {
	return &query.FindContactPersonByIDQuery{
		ID:       uuid.MustParse(input.ContactPersonID),
		UserID:   uuid.MustParse(input.UserID),
		ClientID: uuid.MustParse(input.ClientID),
	}
}

func ToListContactPersonsQuery(input *dto.ListContactPersonsReq, lq *go_kit.ListingQuery) *query.ListContactPersonsQuery {
	return &query.ListContactPersonsQuery{
		UserID:       uuid.MustParse(input.UserID),
		ClientID:     uuid.MustParse(input.ClientID),
		ListingQuery: lq,
	}
}

func ToContactPersonResponse(cp *entity.ContactPerson) *dto.ContactPersonRes {
	return &dto.ContactPersonRes{
		ID:        cp.ID.String(),
		ClientID:  cp.ClientID.String(),
		Name:      cp.Name,
		Email:     cp.Email,
		Phone:     cp.Phone,
		CreatedAt: cp.CreatedAt,
		UpdatedAt: cp.UpdatedAt,
	}
}
