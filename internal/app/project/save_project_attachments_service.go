package project

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

type SaveProjectAttachmentsService struct {
	utilityService *attachment.AttachmentUtilityService
	hasRoleRepo    ports.MemberHasRolePort
	ownerRepo      ports.GetProjectOwnerPort
}

func NewSaveProjectAttachmentsService(utilityService *attachment.AttachmentUtilityService, hasRoleRepo ports.MemberHasRolePort, ownerRepo ports.GetProjectOwnerPort) *SaveProjectAttachmentsService {
	return &SaveProjectAttachmentsService{utilityService: utilityService, hasRoleRepo: hasRoleRepo, ownerRepo: ownerRepo}
}

func (s *SaveProjectAttachmentsService) Execute(ctx context.Context, projectIDStr, userIDStr string, files []interface{}) error {
	projectID := uuid.MustParse(projectIDStr)
	userID := uuid.MustParse(userIDStr)

	roles := []entity.ProjectRole{entity.ProjectRoleOwner, entity.ProjectRoleAdmin}
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
		if _, err := s.utilityService.SaveAttachment(c, ownerID, "project", projectID.String(), *file); err != nil {
			return err
		}
	}

	return nil
}
