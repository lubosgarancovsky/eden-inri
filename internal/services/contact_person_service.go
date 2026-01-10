package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ContactPersonService struct {
	r *repositories.ContactPersonRepository
}

func NewContactPersonService(r *repositories.ContactPersonRepository) *ContactPersonService {
	return &ContactPersonService{r}
}

func (s *ContactPersonService) FindAll(ctx context.Context, userID, clientID uuid.UUID, lq *list.ListingQuery) (*list.Page[models.ContactPerson], error) {
	items, totalCount, err := s.r.FindAll(ctx, userID, clientID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[models.ContactPerson]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ContactPersonService) FindByID(ctx context.Context, userID, clientID, contactPersonID uuid.UUID) (*models.ContactPerson, error) {
	return s.r.FindByID(ctx, userID, clientID, contactPersonID)
}

func (s *ContactPersonService) Create(ctx context.Context, userID, clientID uuid.UUID, input *models.ContactPersonRequest) (*models.ContactPerson, error) {
	cp := &models.ContactPerson{
		UserID:   userID,
		ClientID: clientID,
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
	}
	return s.r.Insert(ctx, cp)
}

func (s *ContactPersonService) Update(ctx context.Context, userID, clientID, contactPersonID uuid.UUID, input *models.ContactPersonRequest) (*models.ContactPerson, error) {
	cp := &models.ContactPerson{
		ID:       contactPersonID,
		UserID:   userID,
		ClientID: clientID,
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
	}
	return s.r.Update(ctx, cp)
}

func (s *ContactPersonService) Delete(ctx context.Context, userID, clientID uuid.UUID, contactPersonID uuid.UUID) error {
	return s.r.Delete(ctx, userID, clientID, contactPersonID)
}
