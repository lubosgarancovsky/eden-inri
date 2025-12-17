package service

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
)

type LabelService struct {
	repo               *repository.LabelRepository
	projectUserService *ProjectUserService
}

func NewLabelService(repo *repository.LabelRepository, pus *ProjectUserService) *LabelService {
	return &LabelService{repo: repo, projectUserService: pus}
}

// FindAll labels for project/board
func (s *LabelService) FindAll(userID, projectID uuid.UUID) ([]model.Label, error) {
	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}
	return s.repo.FindAll(projectID)
}

// FindByID label
func (s *LabelService) FindByID(userID, projectID, labelID uuid.UUID) (*model.Label, error) {
	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}
	return s.repo.FindByID(labelID)
}

// Insert label
func (s *LabelService) Insert(userID, projectID uuid.UUID, req *model.LabelRequest) (*model.Label, error) {
	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}
	label := &model.Label{
		ProjectID:   &projectID,
		Name:        req.Name,
		Description: req.Description,
		Color:       req.Color,
	}
	return s.repo.Insert(label)
}

// Update label
func (s *LabelService) Update(userID, projectID, labelID uuid.UUID, req *model.LabelRequest) (*model.Label, error) {
	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return nil, err
	}
	label, err := s.repo.FindByID(labelID)
	if err != nil {
		return nil, err
	}
	label.Name = req.Name
	label.Description = req.Description
	label.Color = req.Color
	return s.repo.Update(label)
}

// Delete label
func (s *LabelService) Delete(userID, projectID, labelID uuid.UUID) error {
	if _, err := s.projectUserService.GetProjectUserIfMember(projectID, userID); err != nil {
		return err
	}
	return s.repo.Delete(labelID)
}
