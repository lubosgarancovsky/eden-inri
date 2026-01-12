package attachment

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type UploadAttachmentService struct {
	repo      ports.PersistAttachmentPort
	txManager ports.TransactionManager
}

func NewUploadAttachmentService(repo ports.PersistAttachmentPort, txManager ports.TransactionManager) *UploadAttachmentService {
	return &UploadAttachmentService{repo, txManager}
}

func (s *UploadAttachmentService) Execute(ctx context.Context, cmd *command.UploadAttachmentCommand) error {
	attachments, err := instantiateAttachments(cmd.Model, cmd.ModelID, cmd.UserID, cmd.Files)
	if err != nil {
		return err
	}

	tempId := uuid.New().String()
	tempDir := filepath.Join(config.GlobalConfig.UploadsFolder, "tmp", tempId)
	if err = os.MkdirAll(tempDir, os.ModePerm); err != nil {
		return err
	}

	// Save files into a temporary directory before database mutation
	for _, attachment := range attachments {
		if err := saveFileIntoTemp(tempDir, attachment); err != nil {
			return err
		}
	}

	// Save attachments into DB
	err = s.txManager.WithTransaction(ctx, func(txCtx context.Context) error {
		for _, attachment := range attachments {
			if err := s.repo.Create(txCtx, attachment); err != nil {
				return err
			}
		}
		return nil
	})

	// If transaction succeeded, move files from the temp folder
	if err == nil {
		for _, attachment := range attachments {
			if err := moveFileFromTemp(attachment); err != nil {
				return err
			}
		}
	}

	// Remove temp folder
	os.RemoveAll(tempDir)
	return nil
}

func saveFileIntoTemp(tempDir string, attachment *entity.Attachment) error {
	tempPath := filepath.Join(tempDir, attachment.ServerName)
	attachment.TempFilePath = tempPath

	if err := saveFileToDisk(attachment.File, tempPath); err != nil {
		return err
	}

	return nil
}

func moveFileFromTemp(attachment *entity.Attachment) error {
	finalPath := attachment.GetFilePath(config.GlobalConfig.UploadsFolder)
	if err := os.MkdirAll(filepath.Dir(finalPath), os.ModePerm); err != nil {
		return err
	}
	if err := os.Rename(attachment.TempFilePath, finalPath); err != nil {
		return err
	}
	return nil
}

func instantiateAttachments(model, modelID string, userID uuid.UUID, files []*multipart.FileHeader) ([]*entity.Attachment, error) {
	attachments := make([]*entity.Attachment, len(files))
	errCh := make(chan error, len(files)) // buffered to avoid goroutine leak
	var wg sync.WaitGroup

	for i, file := range files {
		wg.Add(1)
		go func(i int, file *multipart.FileHeader) {
			defer wg.Done()
			attachment, err := entity.NewAttachment(userID, model, modelID, file)
			if err != nil {
				errCh <- err
				return
			}
			attachments[i] = attachment
		}(i, file)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return attachments, nil
}

func saveFileToDisk(file *multipart.FileHeader, destDir string) error {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Ensure destination directory exists
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}

	// Full path to save
	destPath := filepath.Join(destDir, file.Filename)

	// Create destination file
	dst, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy file content
	_, err = io.Copy(dst, src)
	return err
}
