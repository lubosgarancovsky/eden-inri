package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ClientService struct {
	r *repository.ClientRepository
}

func NewClientService(r *repository.ClientRepository) *ClientService {
	return &ClientService{r}
}

func (s *ClientService) FindAll(userID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.ClientListItem], error) {
	items, totalCount, err := s.r.FindAll(userID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.ClientListItem]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ClientService) FindByID(userID, clientID uuid.UUID) (*model.Client, error) {
	client, err := s.r.FindByID(clientID)
	if err != nil {
		return nil, err
	}

	if client.UserID != userID {
		return nil, api_err.ErrForbidden
	}

	return client, nil
}

func (s *ClientService) Create(userID uuid.UUID, input *model.ClientRequest) (*model.Client, error) {
	client := fromRequest(userID, input)
	return s.r.Insert(client)
}

func (s *ClientService) Update(userID, clientID uuid.UUID, input *model.ClientRequest) (*model.Client, error) {
	_, err := s.FindByID(userID, clientID)
	if err != nil {
		return nil, err
	}

	client := fromRequest(userID, input)
	client.ID = clientID
	return s.r.Update(client)
}

func (s *ClientService) Delete(userID, clientID uuid.UUID) (*model.Client, error) {
	client, err := s.FindByID(userID, clientID)
	if err != nil {
		return nil, err
	}

	return client, s.r.Delete(clientID)
}

func fromRequest(userID uuid.UUID, input *model.ClientRequest) *model.Client {
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
