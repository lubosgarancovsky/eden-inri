package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateContactPersonCommand(input *dto.CreateContactPersonReq) (*command.CreateContactPersonCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	clientID, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateContactPersonCommand{
		UserID:   userID,
		ClientID: clientID,
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
	}, nil
}

func ToUpdateContactPersonCommand(input *dto.UpdateContactPersonReq) (*command.UpdateContactPersonCommand, error) {
	id, err := uuid.Parse(input.ContactPersonID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	clientID, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateContactPersonCommand{
		ID: id,
		CreateContactPersonCommand: command.CreateContactPersonCommand{
			UserID:   userID,
			ClientID: clientID,
			Name:     input.Name,
			Email:    input.Email,
			Phone:    input.Phone,
		},
	}, nil
}

func ToDeleteContactPersonCommand(input *dto.DeleteContactPersonReq) (*command.DeleteContactPersonCommand, error) {
	id, err := uuid.Parse(input.ContactPersonID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	clientID, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.DeleteContactPersonCommand{
		ID:       id,
		UserID:   userID,
		ClientID: clientID,
	}, nil
}

func ToFindContactPersonByIDQuery(input *dto.FindContactPersonByIDReq) (*query.FindContactPersonByIDQuery, error) {
	id, err := uuid.Parse(input.ContactPersonID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	clientID, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.FindContactPersonByIDQuery{
		ID:       id,
		UserID:   userID,
		ClientID: clientID,
	}, nil
}

func ToListContactPersonsQuery(input *dto.ListContactPersonsReq, lq *go_kit.ListingQuery) (*query.ListContactPersonsQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	clientID, err := uuid.Parse(input.ClientID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.ListContactPersonsQuery{
		UserID:       userID,
		ClientID:     clientID,
		ListingQuery: lq,
	}, nil
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
