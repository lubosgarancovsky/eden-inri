package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ClientService struct {
	r *repository.ClientRepository
}

func NewClientService(r *repository.ClientRepository) *ClientService {
	return &ClientService{r}
}

func (s *ClientService) FindAll(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]model.ClientListItem, int64, error) {
	return s.r.FindAll(ctx, userID, lq)
}

func (s *ClientService) FindByID(ctx context.Context, userID, clientID uuid.UUID) (*model.Client, error) {
	return s.r.FindByID(ctx, userID, clientID)
}

func (s *ClientService) Create(ctx context.Context, userID uuid.UUID, payload *model.ClientRequest) (*model.Client, error) {
	return s.r.Insert(ctx, buildClientPayload(userID, payload))
}

func (s *ClientService) Update(ctx context.Context, userID, clientID uuid.UUID, payload *model.ClientRequest) (*model.Client, error) {
	client := buildClientPayload(userID, payload)
	client.ID = clientID
	return s.r.Update(ctx, client)
}

func (s *ClientService) Delete(ctx context.Context, userID, clientID uuid.UUID) error {
	return s.r.Delete(ctx, userID, clientID)
}

func buildClientPayload(userID uuid.UUID, input *model.ClientRequest) *model.Client {
	return &model.Client{
		ClientListItem: model.ClientListItem{
			UserID:       userID,
			ClientType:   input.ClientType,
			ContractType: input.ContractType,
			Name:         input.Name,
			TaxNumber:    input.TaxNumber,
			Address:      input.Address,
			Tags:         input.Tags,
			HourRate:     input.HourRate,
			StartedAt:    input.StartedAt,
			FinishedAt:   input.FinishedAt,
		},
		Description: input.Description,
	}
}
