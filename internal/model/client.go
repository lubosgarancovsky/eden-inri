package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ClientListItem struct {
	ID           uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null" json:"-"`
	ClientType   string         `json:"clientType"`
	ContractType string         `json:"contractType"`
	Name         string         `json:"name"`
	TaxNumber    string         `json:"taxNumber"`
	Address      string         `json:"address"`
	Tags         pq.StringArray `json:"tags" gorm:"type:text[]" swaggertype:"array,string"`
	HourRate     float64        `json:"hourRate"`
	StartedAt    time.Time      `json:"startedAt"`
	FinishedAt   *time.Time     `json:"finishedAt"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

type Client struct {
	ClientListItem
	Description string `json:"description"`
}

type ClientRequest struct {
	ClientType   string     `json:"clientType"`
	ContractType string     `json:"contractType"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	TaxNumber    string     `json:"taxNumber"`
	Address      string     `json:"address"`
	Tags         []string   `json:"tags"`
	HourRate     float64    `json:"hourRate"`
	StartedAt    time.Time  `json:"startedAt"`
	FinishedAt   *time.Time `json:"finishedAt"`
}

func (cl *ClientListItem) TableName() string {
	return "inri_clients"
}

func (c *Client) TableName() string {
	return "inri_clients"
}
