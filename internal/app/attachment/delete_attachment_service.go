package attachment

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type DeleteAttachmentService struct {
	repo ports.PersistAttachmentPort
}

func NewDeleteAttachmentService(repo ports.PersistAttachmentPort) *DeleteAttachmentService {
	return &DeleteAttachmentService{repo: repo}
}

func (s *DeleteAttachmentService) Execute(ctx context.Context, userIDStr, attachmentIDStr string) error {
	userID := uuid.MustParse(userIDStr)
	attachmentID := uuid.MustParse(attachmentIDStr)

	attachment, err := s.repo.FindByID(ctx, userID, attachmentID)
	if err != nil {
		return err
	}

	if err = s.repo.Delete(ctx, userID, attachmentID); err != nil {
		return err
	}

	go s.DeleteFromDisk(context.Background(), attachment)
	return nil
}

func (s *DeleteAttachmentService) DeleteFromDisk(ctx context.Context, attachment *entity.Attachment) {
	go func() {
		_, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		if err := os.Remove(attachment.GetFilePath(config.GlobalConfig.UploadsFolder)); err != nil {
			fmt.Println(err)
		}
	}()
}
