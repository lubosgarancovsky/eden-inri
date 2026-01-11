package story

import (
	"context"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/attachment"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type SaveStoryAttachmentsService struct {
	utilityService *attachment.AttachmentUtilityService
	hasRoleRepo    ports.MemberHasRolePort
	ownerRepo      ports.GetProjectOwnerPort
}

func NewSaveStoryAttachmentsService(utilityService *attachment.AttachmentUtilityService, hasRoleRepo ports.MemberHasRolePort, ownerRepo ports.GetProjectOwnerPort) *SaveStoryAttachmentsService {
	return &SaveStoryAttachmentsService{utilityService: utilityService, hasRoleRepo: hasRoleRepo, ownerRepo: ownerRepo}
}

func (s *SaveStoryAttachmentsService) Execute(ctx context.Context, projectIDStr, storyIDStr, userIDStr string, files []interface{}) error {
	projectID := uuid.MustParse(projectIDStr)
	storyID := uuid.MustParse(storyIDStr)
	userID := uuid.MustParse(userIDStr)

	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin, entity.ProjectRoleDeveloper}
	hasRole, err := s.hasRoleRepo.HasRole(ctx, userID, projectID, roles)
	if err != nil {
		return err
	}

	if !hasRole {
		return app_err.ErrNotAMember
	}

	ownerID, err := s.ownerRepo.GetOwner(ctx, projectID)
	if err != nil {
		return err
	}

	c, ok := ctx.Value("ginContext").(*gin.Context)
	if !ok {
		return nil
	}

	for _, f := range files {
		file, ok := f.(*multipart.FileHeader)
		if !ok {
			continue
		}
		if _, err := s.utilityService.SaveAttachment(c, ownerID, "story", storyID.String(), *file); err != nil {
			return err
		}
	}

	return nil
}
