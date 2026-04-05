package converter

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ToTaxStatsResponse(input *entity.TaxStats) *dto.TaxStatsRes {
	return &dto.TaxStatsRes{
		Count: input.Count,
		Total: input.Total,
	}
}
