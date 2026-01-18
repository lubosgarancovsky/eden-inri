package attachment

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	app_err "github.com/lubosgarancovsky/eden-inri/internal/domain/error"
)

type DeleteAttachmentService struct {
	repo ports.PersistAttachmentPort
}

func NewDeleteAttachmentService(repo ports.PersistAttachmentPort) *DeleteAttachmentService {
	return &DeleteAttachmentService{repo: repo}
}

func (s *DeleteAttachmentService) Execute(ctx context.Context, cmd *command.Command) error {

	attachment, err := s.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if attachment.UserID != cmd.UserID {
		return app_err.ErrInsufficientPermission
	}

	if err = s.repo.Delete(ctx, cmd.ID); err != nil {
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
