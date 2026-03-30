package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateBusinessEntityCommand(input *dto.CreateBusinessEntityReq) (*command.CreateBusinessEntityCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateBusinessEntityCommand{
		UserID:  userID,
		ICO:     input.ICO,
		DIC:     input.DIC,
		ICDPH:   input.ICDPH,
		Title:   input.Title,
		Email:   input.Email,
		Phone:   input.Phone,
		Street:  input.Street,
		ZipCode: input.ZipCode,
		City:    input.City,
		Country: input.Country,
		IBAN:    input.IBAN,
		SWIFT:   input.SWIFT,
	}, nil
}

func ToUpdateBusinessEntityCommand(input *dto.UpdateBusinessEntityReq) (*command.UpdateBusinessEntityCommand, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateBusinessEntityCommand{
		ID:      id,
		UserID:  userID,
		ICO:     input.ICO,
		DIC:     input.DIC,
		ICDPH:   input.ICDPH,
		Title:   input.Title,
		Email:   input.Email,
		Phone:   input.Phone,
		Street:  input.Street,
		ZipCode: input.ZipCode,
		City:    input.City,
		Country: input.Country,
		IBAN:    input.IBAN,
		SWIFT:   input.SWIFT,
	}, nil
}

func ToDeleteBusinessEntityCommand(input *dto.DeleteBusinessEntityReq) (*command.DeleteBusinessEntityCommand, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.DeleteBusinessEntityCommand{
		ID:     id,
		UserID: userID,
	}, nil
}

func ToBusinessEntityQuery(input *dto.FindBusinessEntityByIDReq) (*query.Query, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.Query{
		ID:     id,
		UserID: userID,
	}, nil
}

func ToListBusinessEntitiesQuery(input *dto.ListBusinessEntitiesReq, lq *go_kit.ListingQuery) (*query.ListQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.ListQuery{
		UserID:       userID,
		ListingQuery: lq,
	}, nil
}

func ToBusinessEntityResponse(be *entity.BusinessEntity) *dto.BusinessEntityRes {
	return &dto.BusinessEntityRes{
		ID:        be.ID.String(),
		ICO:       be.ICO,
		DIC:       be.DIC,
		ICDPH:     be.ICDPH,
		Title:     be.Title,
		Email:     be.Email,
		Phone:     be.Phone,
		Street:    be.Street,
		ZipCode:   be.ZipCode,
		City:      be.City,
		Country:   be.Country,
		IBAN:      be.IBAN,
		SWIFT:     be.SWIFT,
		CreatedAt: be.CreatedAt,
		UpdatedAt: be.UpdatedAt,
	}
}
