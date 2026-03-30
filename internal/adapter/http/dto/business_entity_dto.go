package dto

import "time"

type CreateBusinessEntityReq struct {
	UserID  string  `header:"X-User-ID"`
	ICO     string  `json:"ico"`
	DIC     string  `json:"dic"`
	ICDPH   *string `json:"icDph"`
	Title   string  `json:"title"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Street  string  `json:"street"`
	ZipCode string  `json:"zipCode"`
	City    string  `json:"city"`
	Country string  `json:"country"`
	IBAN    string  `json:"iban"`
	SWIFT   string  `json:"swift"`
}

type UpdateBusinessEntityReq struct {
	ID      string  `uri:"businessEntityId"`
	UserID  string  `header:"X-User-ID"`
	ICO     string  `json:"ico"`
	DIC     string  `json:"dic"`
	ICDPH   *string `json:"icDph"`
	Title   string  `json:"title"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Street  string  `json:"street"`
	ZipCode string  `json:"zipCode"`
	City    string  `json:"city"`
	Country string  `json:"country"`
	IBAN    string  `json:"iban"`
	SWIFT   string  `json:"swift"`
}

type DeleteBusinessEntityReq struct {
	ID     string `uri:"businessEntityId"`
	UserID string `header:"X-User-ID"`
}

type FindBusinessEntityByIDReq struct {
	ID     string `uri:"businessEntityId"`
	UserID string `header:"X-User-ID"`
}

type ListBusinessEntitiesReq struct {
	UserID string `header:"X-User-ID"`
}

type BusinessEntityRes struct {
	ID        string    `json:"id"`
	ICO       string    `json:"ico"`
	DIC       string    `json:"dic"`
	ICDPH     *string   `json:"icDph"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Street    string    `json:"street"`
	ZipCode   string    `json:"zipCode"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	IBAN      string    `json:"iban"`
	SWIFT     string    `json:"swift"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
