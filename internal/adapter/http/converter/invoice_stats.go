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

func ToMonthlyRevenueResponse(input *[]entity.InvoiceMonthlyRevenue) []dto.MonthlyRevenueRes {
	if input == nil {
		return []dto.MonthlyRevenueRes{}
	}

	results := make([]dto.MonthlyRevenueRes, len(*input))
	for i, domain := range *input {
		results[i] = dto.MonthlyRevenueRes{Month: domain.Month, Revenue: domain.Revenue}
	}

	return results
}
