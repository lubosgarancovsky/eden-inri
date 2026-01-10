package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ProjectUserService struct {
	r *repositories.ProjectUserRepository
}

func NewProjectUserService(r *repositories.ProjectUserRepository) *ProjectUserService {
	return &ProjectUserService{r: r}
}

func (s *ProjectUserService) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[models.ProjectUser], error) {
	items, totalCount, err := s.r.FindAll(ctx, projectID, lq)
	if err != nil {
		return nil, err
	}
	return &list.Page[models.ProjectUser]{
		Items:      *items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ProjectUserService) FindByID(ctx context.Context, userID, projectID uuid.UUID) (*models.ProjectUser, error) {
	return s.r.FindByID(ctx, userID, projectID)
}

func (s *ProjectUserService) Insert(ctx context.Context, userID, projectID uuid.UUID, role models.ProjectRole) (*models.ProjectUser, error) {
	projectUser := &models.ProjectUser{
		UserID:    userID,
		ProjectID: projectID,
		Role:      role,
	}

	return s.r.Insert(ctx, projectUser)
}

func (s *ProjectUserService) Update(
	ctx context.Context,
	callerID, projectID, memberID uuid.UUID,
	newRole models.ProjectRole,
) (*models.ProjectUser, error) {
	// Prevent self-role change
	if callerID == memberID {
		return nil, api_err.ErrForbidden.WithMessage("cannot change your own role")
	}

	projectUser := &models.ProjectUser{
		UserID:    memberID,
		ProjectID: projectID,
		Role:      newRole,
	}

	return s.r.Update(ctx, projectUser)
}

func (s *ProjectUserService) Delete(ctx context.Context, memberID, projectID uuid.UUID) error {
	return s.r.Delete(ctx, memberID, projectID)
}

func (s *ProjectUserService) Favourite(ctx context.Context, userID, projectID uuid.UUID) (*models.ProjectUser, error) {
	projectUser, err := s.r.FindByID(ctx, userID, projectID)
	if err != nil {
		return nil, err
	}

	newIsStarred := !projectUser.IsStarred
	return s.r.Favourite(ctx, userID, projectID, newIsStarred)
}

func (s *ProjectUserService) GetUserRole(ctx context.Context, projectID, userID uuid.UUID) (models.ProjectRole, error) {
	return s.r.GetUserRole(ctx, projectID, userID)
}

func (s *ProjectUserService) RequireRole(actual models.ProjectRole, allowed ...models.ProjectRole) error {
	for _, r := range allowed {
		if actual == r {
			return nil
		}
	}
	return api_err.ErrForbidden
}

func (s *ProjectUserService) GetOwner(ctx context.Context, projectID uuid.UUID) (*models.ProjectUser, error) {
	return s.r.GetOwner(ctx, projectID)
}
