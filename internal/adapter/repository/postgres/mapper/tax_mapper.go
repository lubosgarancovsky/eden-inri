package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func TaxFromDomain(tax *entity.Tax) *model.Tax {
	return &model.Tax{
		ID:          tax.ID,
		UserID:      tax.UserID,
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
