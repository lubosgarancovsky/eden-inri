package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/delivery/http/dto"
	"github.com/lubosgarancovsky/go-kit/array"
)

type ClientType string

var (
	ClientIndividual   ClientType = "INDIVIDUAL"
	ClientOrganisation ClientType = "ORGANISATION"
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
	Description  string
	Name         string
	TaxNumber    string
	Address      string
	Tags         []string
	HourRate     float64
	StartedAt    time.Time
	FinishedAt   *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewClient(userID uuid.UUID, req *dto.CreateClientReq) (*Client, error) {
	clientTypes := []any{ClientIndividual, ClientOrganisation}
	if !array.Includes(clientTypes, req.ClientType) {
		return nil, errors.New("unrecognized client type")
	}

	contractTypes := []any{FullTime, PartTime, Contract, ExternalContract}
	if !array.Includes(contractTypes, req.ContractType) {
		return nil, errors.New("unrecognized contract type")
	}

	if req.HourRate < 0 {
		return nil, errors.New("hour rate cannot be negative")
	}

	now := time.Now()

	return &Client{
		ID:           uuid.New(),
		UserID:       userID,
		ClientType:   ClientType(req.ClientType),
		ContractType: ClientContractType(req.ContractType),
		Description:  req.Description,
		Name:         req.Name,
		TaxNumber:    req.TaxNumber,
		Address:      req.Address,
		Tags:         req.Tags,
		HourRate:     req.HourRate,
		StartedAt:    req.StartedAt,
		FinishedAt:   req.FinishedAt,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}
