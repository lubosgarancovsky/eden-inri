package dto

import "time"

type CreateTaxReq struct {
	UserID      string     `header:"X-User-ID"`
	Category    string     `json:"category"`
	Amount      float64    `json:"amount"`
	Currency    string     `json:"currency"`
	Description *string    `json:"description"`
	Reference   *string    `json:"reference"`
	Period      *time.Time `json:"period"`
	PaidAt      *time.Time `json:"paidAt"`
}

type UpdateTaxReq struct {
	ID          string     `uri:"taxId"`
	UserID      string     `header:"X-User-ID"`
	Description *string    `json:"description"`
	Reference   *string    `json:"reference"`
	PaidAt      *time.Time `json:"paidAt"`
}

type DeleteTaxReq struct {
	ID     string `uri:"taxId"`
	UserID string `header:"X-User-ID"`
}

type FindTaxByIDReq struct {
	ID     string `uri:"taxId"`
	UserID string `header:"X-User-ID"`
}

type ListTaxesReq struct {
	UserID string `header:"X-User-ID"`
}

type TaxRes struct {
	ID          string     `json:"id"`
	Category    string     `json:"category"`
	Amount      float64    `json:"amount"`
	Currency    string     `json:"currency"`
	Description *string    `json:"description"`
	Reference   *string    `json:"reference"`
	Period      *time.Time `json:"period"`
	PaidAt      *time.Time `json:"paidAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
