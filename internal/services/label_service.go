package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/go-kit/list"
)

type LabelService struct {
	repo               *repositories.LabelRepository
	projectUserService *ProjectUserService
}

func NewLabelService(repo *repositories.LabelRepository, pus *ProjectUserService) *LabelService {
	return &LabelService{repo: repo, projectUserService: pus}
}

// FindAll labels for project/board
func (s *LabelService) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[models.Label], error) {
	items, totalCount, err := s.repo.FindAll(ctx, projectID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[models.Label]{
		Items:      *items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

// FindByID label
func (s *LabelService) FindByID(ctx context.Context, projectID, labelID uuid.UUID) (*models.Label, error) {
	return s.repo.FindByID(ctx, projectID, labelID)
}

// Insert label
func (s *LabelService) Insert(ctx context.Context, projectID uuid.UUID, req *models.LabelRequest) (*models.Label, error) {
	label := &models.Label{
		ProjectID:   &projectID,
		Name:        req.Name,
		Description: req.Description,
		Color:       req.Color,
	}
	return s.repo.Insert(ctx, label)
}

// Update label
func (s *LabelService) Update(ctx context.Context, projectID, labelID uuid.UUID, req *models.LabelRequest) (*models.Label, error) {
	label := &models.Label{
		ID:          labelID,
		ProjectID:   &projectID,
		Name:        req.Name,
		Description: req.Description,
		Color:       req.Color,
	}
	return s.repo.Update(ctx, label)
}

// Delete label
func (s *LabelService) Delete(ctx context.Context, projectID, labelID uuid.UUID) error {
	return s.repo.Delete(ctx, projectID, labelID)
}
