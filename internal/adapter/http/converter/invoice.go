package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

func ToInvoicePortFromCreate(req *dto.CreateInvoiceReq) *ports.Invoice {
	return &ports.Invoice{
		ID:          uuid.New(),
		UserID:      req.UserID,
		ClientID:    req.ClientID,
		Number:      req.Name, // legacy field mapping (name used as invoice number)
		IssuedAt:    req.IssuedAt,
		DueAt:       req.DueAt,
		TotalAmount: int64(req.Total),
		Currency:    "", // not present in legacy DTO
		Status:      "", // from legacy Update endpoint status
		Note:        req.Note,
	}
}

func ToInvoicePortFromUpdate(req *dto.UpdateInvoiceReq) *ports.Invoice {
	inv := ToInvoicePortFromCreate(&req.CreateInvoiceReq)
	inv.ID = req.InvoiceID
	inv.UserID = req.UserID
	return inv
}

func ToInvoiceResponse(inv *ports.Invoice) *dto.InvoiceRes {
	// Map to legacy InvoiceRes fields for backward compatibility in response shape
	return &dto.InvoiceRes{
		ID:       inv.ID,
		Name:     inv.Number,
		Note:     inv.Note,
		Total:    float64(inv.TotalAmount),
		IssuedAt: inv.IssuedAt,
		DueAt:    inv.DueAt,
	}
}

func ToInvoiceListResponse(items *[]ports.Invoice) []*dto.InvoiceRes {
	if items == nil {
		return nil
	}
	arr := *items
	res := make([]*dto.InvoiceRes, len(arr))
	for i := range arr {
		res[i] = ToInvoiceResponse(&arr[i])
	}
	return res
}
