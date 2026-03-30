package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func BusinessEntityFromDomain(be *entity.BusinessEntity) *model.BusinessEntity {
	return &model.BusinessEntity{
		ID:        be.ID,
		UserID:    be.UserID,
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
