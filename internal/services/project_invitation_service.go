package services

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
)

type ProjectInvitationService struct {
	cfg *config.Config
	r   *repositories.ProjectInvitationRepository
	prs *ProjectService
	pus *ProjectUserService
	es  *EmailService
}

func NewProjectInvitationService(
	cfg *config.Config,
	repository *repositories.ProjectInvitationRepository,
	prs *ProjectService,
	pus *ProjectUserService,
	es *EmailService,
) *ProjectInvitationService {
	return &ProjectInvitationService{cfg, repository, prs, pus, es}
}

func (s *ProjectInvitationService) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.r.FindUserByEmail(ctx, email)
}

func (s *ProjectInvitationService) Create(ctx context.Context, userID, projectID uuid.UUID, req *models.ProjectInvitationRequest) error {
	user, err := s.r.FindUserByEmail(ctx, req.Email)
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

	invitation := &models.ProjectInvitation{
		ProjectID: projectID,
		UserID:    user.ID,
		Role:      req.Role,
		Token:     token,
		InvitedBy: userID,
		ExpiresAt: time.Now().Add(time.Second * time.Duration(s.cfg.InvitationTokenExp)),
	}

	project, err := s.prs.FindByID(ctx, userID, projectID)
	if err != nil {
		return err
	}

	template := &models.ProjectInvitationTemplate{
		URL:         fmt.Sprintf("%s?token=%s", s.cfg.InvitationUrl, token),
		ProjectName: project.Name,
		UserName:    user.Username,
	}

	if err = s.SendInvitationEmail(user.Email, template); err != nil {
		return err
	}

	_, err = s.r.Insert(ctx, invitation)
	return err
}

func (s *ProjectInvitationService) Accept(ctx context.Context, userID uuid.UUID, token string) (*models.ProjectInvitation, error) {
	invitation, err := s.r.FindByToken(ctx, userID, token)
	if err != nil {
		return nil, err
	}

	if invitation.ExpiresAt.Before(time.Now()) {
		return nil, api_err.ErrForbidden.WithMessage("Invitation expired")
	}

	if _, err = s.pus.Insert(ctx, invitation.UserID, invitation.ProjectID, invitation.Role); err != nil {
		return nil, err
	}

	acceptedInvitation := invitation
	acceptedInvitation.AcceptedAt = time.Now()
	result, err := s.r.Update(ctx, invitation)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ProjectInvitationService) SendInvitationEmail(email string, template *models.ProjectInvitationTemplate) error {
	templatePath := path.Join(s.cfg.TemplatesFolder, "project-invitation.html")
	return s.es.SendTemplateEmail(email, "Eden - Project invitation", templatePath, template)
}
