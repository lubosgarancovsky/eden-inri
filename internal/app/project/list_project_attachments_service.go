package project

import (
	"context"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type ListProjectAttachmentsService struct {
	repo         ports.PersistAttachmentPort
	isMemberRepo ports.IsProjectMemberPort
}

func NewListProjectAttachmentsService(repo ports.PersistAttachmentPort, isMemberRepo ports.IsProjectMemberPort) *ListProjectAttachmentsService {
	return &ListProjectAttachmentsService{repo: repo, isMemberRepo: isMemberRepo}
}

func (s *ListProjectAttachmentsService) Execute(ctx context.Context, projectIDStr, userIDStr string) ([]entity.Attachment, error) {
	projectID := uuid.MustParse(projectIDStr)
	userID := uuid.MustParse(userIDStr)

	isMember, err := s.isMemberRepo.IsMember(ctx, userID, projectID)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return nil, nil // Or appropriate error
	}

	return s.repo.FindByModelID(ctx, userID, projectID, "project")
}
