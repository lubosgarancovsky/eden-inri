package dto

import "time"

type CreateContactPersonReq struct {
	UserID   string  `header:"X-User-ID"`
	ClientID string  `uri:"clientId"`
	Name     string  `json:"name"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

type UpdateContactPersonReq struct {
	UserID          string  `header:"X-User-ID"`
	ClientID        string  `uri:"clientId"`
	ContactPersonID string  `uri:"contactPersonId"`
	Name            string  `json:"name"`
	Email           *string `json:"email"`
	Phone           *string `json:"phone"`
}

type DeleteContactPersonReq struct {
	UserID          string `header:"X-User-ID"`
	ClientID        string `uri:"clientId"`
	ContactPersonID string `uri:"contactPersonId"`
}

type FindContactPersonByIDReq struct {
	UserID          string `header:"X-User-ID"`
	ClientID        string `uri:"clientId"`
	ContactPersonID string `uri:"contactPersonId"`
}

type ListContactPersonsReq struct {
	UserID   string `header:"X-User-ID"`
	ClientID string `uri:"clientId"`
}

type ContactPersonRes struct {
	ID        string    `json:"id"`
	ClientID  string    `json:"clientId"`
	Name      string    `json:"name"`
	Email     *string   `json:"email"`
	Phone     *string   `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
