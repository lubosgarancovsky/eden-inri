package converter

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ToInvoiceStatsResponse(input *entity.InvoiceStats) *dto.InvoiceStatsRes {
	return &dto.InvoiceStatsRes{
		Count:         input.Count,
		Total:         input.Total,
		BillableHours: input.BillableHours,
	}
}
