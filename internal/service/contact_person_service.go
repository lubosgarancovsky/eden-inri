package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ContactPersonService struct {
	r *repository.ContactPersonRepository
}

func NewContactPersonService(r *repository.ContactPersonRepository) *ContactPersonService {
	return &ContactPersonService{r}
}

func (s *ContactPersonService) FindAll(userID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.ContactPerson], error) {
	items, totalCount, err := s.r.FindAll(userID, lq)
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

func (s *ContactPersonService) FindByID(userID uuid.UUID, id uuid.UUID) (*model.ContactPerson, error) {
	cp, err := s.r.FindByID(id)
	if err != nil {
		return nil, err
	}
	if cp.UserID != userID {
		return nil, api_err.ErrForbidden
	}
	return cp, nil
}

func (s *ContactPersonService) Create(userID uuid.UUID, input *model.ContactPersonRequest) (*model.ContactPerson, error) {
	cp := &model.ContactPerson{
		UserID:   userID,
		ClientID: input.ClientID,
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
	}
	return s.r.Insert(cp)
}

func (s *ContactPersonService) Update(userID uuid.UUID, id uuid.UUID, input *model.ContactPersonRequest) (*model.ContactPerson, error) {
	_, err := s.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	cp := &model.ContactPerson{
		ID:       id,
		UserID:   userID,
		ClientID: input.ClientID,
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
	}
	return s.r.Update(cp)
}

func (s *ContactPersonService) Delete(userID uuid.UUID, id uuid.UUID) (*model.ContactPerson, error) {
	cp, err := s.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	return cp, s.r.Delete(id)
}
