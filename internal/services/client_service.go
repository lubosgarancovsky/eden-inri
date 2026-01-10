package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ClientService struct {
	r *repositories.ClientRepository
}

func NewClientService(r *repositories.ClientRepository) *ClientService {
	return &ClientService{r}
}

func (s *ClientService) FindAll(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]models.ClientListItem, int64, error) {
	return s.r.FindAll(ctx, userID, lq)
}

func (s *ClientService) FindByID(ctx context.Context, userID, clientID uuid.UUID) (*models.Client, error) {
	return s.r.FindByID(ctx, userID, clientID)
}

func (s *ClientService) Create(ctx context.Context, userID uuid.UUID, payload *models.ClientRequest) (*models.Client, error) {
	return s.r.Insert(ctx, buildClientPayload(userID, payload))
}

func (s *ClientService) Update(ctx context.Context, userID, clientID uuid.UUID, payload *models.ClientRequest) (*models.Client, error) {
	client := buildClientPayload(userID, payload)
	client.ID = clientID
	return s.r.Update(ctx, client)
}

func (s *ClientService) Delete(ctx context.Context, userID, clientID uuid.UUID) error {
	return s.r.Delete(ctx, userID, clientID)
}

func buildClientPayload(userID uuid.UUID, input *models.ClientRequest) *models.Client {
	return &models.Client{
		ClientListItem: models.ClientListItem{
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
