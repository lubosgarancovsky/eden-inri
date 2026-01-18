package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func ProjectInvitationFromDomain(inv *entity.ProjectInvitation) *model.ProjectInvitation {
	return &model.ProjectInvitation{
		ID:         inv.ID,
		ProjectID:  inv.ProjectID,
		UserID:     inv.UserID,
		Role:       inv.Role,
		Token:      inv.Token,
		InvitedBy:  inv.InvitedBy,
		CreatedAt:  inv.CreatedAt,
		ExpiresAt:  inv.ExpiresAt,
		AcceptedAt: inv.AcceptedAt,
	}
}
