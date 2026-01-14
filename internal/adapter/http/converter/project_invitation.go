package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
)

func ToProjectInvitationCommand(input *dto.ProjectInvitationReq) (*command.InviteUserToProjectCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.InviteUserToProjectCommand{
		UserID:    userID,
		ProjectID: projectID,
		Email:     input.Email,
		Role:      entity.ProjectRole(input.Role),
	}, nil
}

func ToAcceptInvitationCommand(input *dto.AcceptInvitationReq) (*command.AcceptProjectInvitationCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.AcceptProjectInvitationCommand{
		UserID: userID,
		Token:  input.Token,
	}, nil
}

func ToProjectInvitationResponse(inv *entity.ProjectInvitation) *dto.ProjectInvitationRes {
	return &dto.ProjectInvitationRes{
		ID:         inv.ID,
		ProjectID:  inv.ProjectID,
		UserID:     inv.UserID,
		InvitedBy:  inv.InvitedBy,
		Role:       string(inv.Role),
		CreatedAt:  inv.CreatedAt,
		ExpiresAt:  inv.ExpiresAt,
		AcceptedAt: inv.AcceptedAt,
	}
}
