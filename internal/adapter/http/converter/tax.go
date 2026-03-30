package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateTaxCommand(input *dto.CreateTaxReq) (*command.CreateTaxCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateTaxCommand{
		UserID:      userID,
		Category:    entity.TaxCategory(input.Category),
		Amount:      input.Amount,
		Currency:    input.Currency,
		Description: input.Description,
		Reference:   input.Reference,
		Period:      input.Period,
		PaidAt:      input.PaidAt,
	}, nil
}

func ToUpdateTaxCommand(input *dto.UpdateTaxReq) (*command.UpdateTaxCommand, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateTaxCommand{
		ID:          id,
		UserID:      userID,
		Description: input.Description,
		Reference:   input.Reference,
		PaidAt:      input.PaidAt,
	}, nil
}

func ToDeleteTaxCommand(input *dto.DeleteTaxReq) (*command.DeleteTaxCommand, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.DeleteTaxCommand{
		ID:     id,
		UserID: userID,
	}, nil
}

func ToTaxQuery(input *dto.FindTaxByIDReq) (*query.Query, error) {
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

func ToListTaxesQuery(input *dto.ListTaxesReq, lq *go_kit.ListingQuery) (*query.ListQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.ListQuery{
		UserID:       userID,
		ListingQuery: lq,
	}, nil
}

func ToTaxResponse(tax *entity.Tax) *dto.TaxRes {
	return &dto.TaxRes{
		ID:          tax.ID.String(),
		Category:    string(tax.Category),
		Amount:      tax.Amount,
		Currency:    tax.Currency,
		Description: tax.Description,
		Reference:   tax.Reference,
		Period:      tax.Period,
		PaidAt:      tax.PaidAt,
		CreatedAt:   tax.CreatedAt,
		UpdatedAt:   tax.UpdatedAt,
	}
}
