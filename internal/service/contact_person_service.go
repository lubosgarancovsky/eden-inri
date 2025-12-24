package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ContactPersonService struct {
	r *repository.ContactPersonRepository
}

func NewContactPersonService(r *repository.ContactPersonRepository) *ContactPersonService {
	return &ContactPersonService{r}
}

func (s *ContactPersonService) FindAll(ctx context.Context, userID, clientID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.ContactPerson], error) {
	items, totalCount, err := s.r.FindAll(ctx, userID, clientID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.ContactPerson]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ContactPersonService) FindByID(ctx context.Context, userID, clientID, contactPersonID uuid.UUID) (*model.ContactPerson, error) {
	return s.r.FindByID(ctx, userID, clientID, contactPersonID)
}

func (s *ContactPersonService) Create(ctx context.Context, userID, clientID uuid.UUID, input *model.ContactPersonRequest) (*model.ContactPerson, error) {
	cp := &model.ContactPerson{
		UserID:   userID,
		ClientID: clientID,
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
	}
	return s.r.Insert(ctx, cp)
}

func (s *ContactPersonService) Update(ctx context.Context, userID, clientID, contactPersonID uuid.UUID, input *model.ContactPersonRequest) (*model.ContactPerson, error) {
	cp := &model.ContactPerson{
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
