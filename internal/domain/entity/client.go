package entity

import (
	"time"

	"github.com/google/uuid"
)

type ClientType string

var (
	ClientIndividual   ClientType = "INDIVIDUAL"
	ClientOrganisation ClientType = "ORGANIZATION"
)

type ClientContractType string

var (
	FullTime         ClientContractType = "FULL_TIME"
	PartTime         ClientContractType = "PART_TIME"
	Contract         ClientContractType = "CONTRACT"
	ExternalContract ClientContractType = "EXTERNAL_CONTRACT"
)

type Client struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ClientType   ClientType
	ContractType ClientContractType
	Name         string
	Description  *string
	TaxNumber    *string
	Address      *string
	Tags         []string
	HourRate     float64
	StartedAt    time.Time
	FinishedAt   *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
