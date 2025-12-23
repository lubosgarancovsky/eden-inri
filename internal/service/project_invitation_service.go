package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type ProjectInvitationService struct {
	cfg *config.Config
	r   *repository.ProjectInvitationRepository
	prs *ProjectService
	pus *ProjectUserService
	es  *EmailService
}

func NewProjectInvitationService(
	cfg *config.Config,
	repository *repository.ProjectInvitationRepository,
	prs *ProjectService,
	pus *ProjectUserService,
	es *EmailService,
) *ProjectInvitationService {
	return &ProjectInvitationService{cfg, repository, prs, pus, es}
}

func (s *ProjectInvitationService) FindUserByEmail(email string) (*model.User, error) {
	return s.r.FindUserByEmail(email)
}

func (s *ProjectInvitationService) Create(userID, projectID uuid.UUID, req *model.ProjectInvitationRequest) error {
	user, err := s.r.FindUserByEmail(req.Email)
	if err != nil {
		return err
	}

	if user.ID == userID {
		return api_err.ErrBadRequest.WithMessage("You cannot invite yourself")
	}

	token, err := helpers.OpaqueToken(32)
	if err != nil {
		return err
	}

	invitation := &model.ProjectInvitation{
		ProjectID: projectID,
		UserID:    user.ID,
		Role:      req.Role,
		Token:     token,
		InvitedBy: userID,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(s.cfg.InvitationTokenExp)),
	}

	project, err := s.prs.FindByID(userID, projectID)
	if err != nil {
		return err
	}

	template := &model.ProjectInvitationTemplate{
		URL:         fmt.Sprintf("%s?token=%s", s.cfg.InvitationUrl, token),
		ProjectName: project.Name,
		UserName:    user.Username,
	}

	if err = s.SendInvitationEmail(user.Email, template); err != nil {
		return err
	}

	_, err = s.r.Insert(invitation)
	return err
}

func (s *ProjectInvitationService) Accept(userID uuid.UUID, token string) (*model.ProjectInvitation, error) {
	invitation, err := s.r.FindByToken(userID, token)
	if err != nil {
		return nil, err
	}

	if err = s.pus.Insert(invitation.UserID, invitation.ProjectID, invitation.Role); err != nil {
		return nil, err
	}

	acceptedInvitation := invitation
	acceptedInvitation.AcceptedAt = time.Now()
	result, err := s.r.Update(invitation)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProjectInvitationService) SendInvitationEmail(email string, template *model.ProjectInvitationTemplate) error {
	templatePath := "templates/project-invitation.html"
	return s.es.SendTemplateEmail(email, "Eden - Project invitation", templatePath, template)
}
