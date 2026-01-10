package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/internal/repositories"
	"github.com/lubosgarancovsky/go-kit/list"
)

type AttachmentService struct {
	cfg *config.Config
	r   *repositories.AttachmentRepository
}

func NewAttachmentService(cfg *config.Config, r *repositories.AttachmentRepository) *AttachmentService {
	return &AttachmentService{cfg, r}
}

func (s *AttachmentService) FindAll(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]models.Attachment, int64, error) {
	return s.r.FindAll(ctx, userID, lq)
}

func (s *AttachmentService) FindByID(ctx context.Context, userID uuid.UUID, attachmentID uuid.UUID) (*models.Attachment, error) {
	result, err := s.r.FindByID(ctx, userID, attachmentID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *AttachmentService) FindByModelID(ctx context.Context, userID, modelID uuid.UUID, modelName string) ([]*models.Attachment, error) {
	return s.r.FindByModelID(ctx, userID, modelID, modelName)
}

func (s *AttachmentService) Insert(ctx context.Context, attachment *models.Attachment) (*models.Attachment, error) {
	return s.r.Insert(ctx, attachment)
}

func (s *AttachmentService) Delete(ctx context.Context, userID, attachmentID uuid.UUID) error {
	attachment, err := s.FindByID(ctx, userID, attachmentID)
	if err != nil {
		return err
	}

	if err = s.r.Delete(ctx, userID, attachmentID); err != nil {
		return err
	}

	go s.DeleteFromDisk(context.Background(), attachment)
	return nil
}

func (s *AttachmentService) GetFilePath(attachment *models.Attachment) string {
	return filepath.Join(s.cfg.UploadsFolder, attachment.UserID.String(), attachment.Model, attachment.ModelID, attachment.OriginalName)
}

func (s *AttachmentService) GetFileName(originalName string) string {
	return strings.ToLower(strings.TrimSpace(originalName))
}

func (s *AttachmentService) GetMimeType(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	buf := make([]byte, 512)
	n, _ := file.Read(buf)
	return http.DetectContentType(buf[:n]), nil
}

func (s *AttachmentService) SaveAttachment(c *gin.Context, userID uuid.UUID, modelName string, modelID string, file multipart.FileHeader) (*models.Attachment, error) {
	mime, err := s.GetMimeType(&file)
	if err != nil {
		return nil, err
	}

	attachment := &models.Attachment{
		Model:        modelName,
		ModelID:      modelID,
		UserID:       userID,
		Size:         file.Size,
		OriginalName: file.Filename,
		MimeType:     mime,
		ServerName:   s.GetFileName(file.Filename),
	}

	if err := s.SaveFile(c, attachment, file); err != nil {
		return nil, err
	}

	return s.Insert(c.Request.Context(), attachment)
}

func (s *AttachmentService) SaveFile(c *gin.Context, attachment *models.Attachment, file multipart.FileHeader) error {
	path := s.GetFilePath(attachment)
	return c.SaveUploadedFile(&file, path)
}

func (s *AttachmentService) DeleteFromDisk(ctx context.Context, attachment *models.Attachment) {
	go func() {
		_, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		path := s.GetFilePath(attachment)
		if err := os.Remove(path); err != nil {
			fmt.Println(err)
		}
	}()
}
