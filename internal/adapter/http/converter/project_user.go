package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
)

func ToChangeProjectUserRoleCommand(input *dto.ChangeProjectUserRoleReq) (*command.ChangeProjectUserRoleCommand, error) {
	UserID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	ProjectID, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	MemberID, err := uuid.Parse(input.MemberID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.ChangeProjectUserRoleCommand{
		UserID:    UserID,
		ProjectID: ProjectID,
		MemberID:  MemberID,
		Role:      entity.ProjectRole(input.Role),
	}, nil
}

func ToProjectUserResponse(pu *entity.ProjectUser) *dto.ProjectUserRes {
	return &dto.ProjectUserRes{
		Role:      string(pu.Role),
		IsStarred: pu.IsStarred,
		JoinedAt:  pu.JoinedAt,
		User:      ToUserResponse(pu.User),
	}
}
