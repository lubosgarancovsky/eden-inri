package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
)

type ProjectUserService struct {
	r *repository.ProjectUserRepository
}

func NewProjectUserService(r *repository.ProjectUserRepository) *ProjectUserService {
	return &ProjectUserService{r: r}
}

func (s *ProjectUserService) FindAll(userID, projectID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.ProjectUser], error) {
	items, totalCount, err := s.r.FindAll(userID, projectID, lq)
	if err != nil {
		return nil, err
	}
	return &list.Page[model.ProjectUser]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *ProjectUserService) Insert(userID, projectID uuid.UUID, role model.ProjectRole) error {
	role, err := s.GetUserRole(projectID, userID)
	if err != nil {
		return err
	}

	if err := s.RequireRole(role, model.Admin, model.Owner); err != nil {
		return err
	}

	return s.Insert(userID, projectID, role)
}

func (s *ProjectUserService) Update(
	callerID, projectID, memberID uuid.UUID,
	newRole model.ProjectRole,
) (*model.ProjectUser, error) {

	// Fetch caller role
	callerPU, err := s.r.GetProjectUserIfMember(projectID, callerID)
	if err != nil {
		return nil, err
	}

	if callerPU.Role != model.Owner && callerPU.Role != model.Admin {
		return nil, api_err.ErrForbidden.WithMessage("only owner or admin can update member roles")
	}

	// Prevent self-role change
	if callerID == memberID {
		return nil, api_err.ErrForbidden.WithMessage("cannot change your own role")
	}

	// Prevent updating owner role
	memberPU, err := s.r.GetProjectUserIfMember(projectID, memberID)
	if err != nil {
		return nil, err
	}

	if memberPU.Role == model.Owner {
		return nil, api_err.ErrForbidden.WithMessage("cannot change role of owner")
	}

	return s.r.Update(projectID, memberID, newRole)
}

func (s *ProjectUserService) Delete(userID, memberID, projectID uuid.UUID) error {
	role, err := s.GetUserRole(projectID, userID)
	if err != nil {
		return err
	}

	if err := s.RequireRole(role, model.Admin, model.Owner); err != nil {
		return err
	}

	memberRole, err := s.GetUserRole(projectID, memberID)
	if err != nil {
		return err
	}

	if memberRole == model.Owner {
		return api_err.ErrForbidden
	}

	return s.r.Delete(userID, projectID)
}

func (s *ProjectUserService) Favourite(userID, projectID uuid.UUID) (*model.ProjectUser, error) {
	memberPU, err := s.r.GetProjectUserIfMember(projectID, userID)
	if err != nil {
		return nil, err
	}

	newIsStarred := !memberPU.IsStarred
	return s.r.Favourite(userID, projectID, newIsStarred)
}

func (s *ProjectUserService) GetProjectUserIfMember(projectID, userID uuid.UUID) (*model.ProjectUser, error) {
	return s.r.GetProjectUserIfMember(projectID, userID)
}

func (s *ProjectUserService) GetUserRole(projectID, userID uuid.UUID) (model.ProjectRole, error) {
	return s.r.GetUserRole(projectID, userID)
}

func (s *ProjectUserService) RequireRole(actual model.ProjectRole, allowed ...model.ProjectRole) error {
	for _, r := range allowed {
		if actual == r {
			return nil
		}
	}
	return api_err.ErrForbidden
}
