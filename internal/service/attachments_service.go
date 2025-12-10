package service

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
)

type AttachmentService struct {
	cfg *config.Config
	r   *repository.AttachmentRepository
}

func NewAttachmentService(cfg *config.Config, r *repository.AttachmentRepository) *AttachmentService {
	return &AttachmentService{cfg, r}
}

func (s *AttachmentService) SaveAttachment(c *gin.Context, userID uuid.UUID, modelName string, modelID string, file multipart.FileHeader) (*model.Attachment, error) {
	mime, err := s.GetMimeType(&file)
	if err != nil {
		return nil, err
	}

	attachment := &model.Attachment{
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

	return s.Insert(attachment)
}

func (s *AttachmentService) SaveFile(c *gin.Context, attachment *model.Attachment, file multipart.FileHeader) error {
	path := s.GetFilePath(attachment)
	return c.SaveUploadedFile(&file, path)
}

func (s *AttachmentService) FindAll(userID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.Attachment], error) {
	items, totalCount, err := s.r.FindAll(userID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.Attachment]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *AttachmentService) FindByID(userID uuid.UUID, attachmentID uuid.UUID) (*model.Attachment, error) {
	result, err := s.r.FindByID(attachmentID)
	if err != nil {
		return nil, err
	}

	if result.UserID != userID {
		return nil, api_err.ErrForbidden
	}

	return result, nil
}

func (s *AttachmentService) FindByModelID(userID uuid.UUID, modelName string, modelID uuid.UUID) ([]*model.Attachment, error) {
	return s.r.FindByModelID(userID, modelName, modelID)
}

func (s *AttachmentService) Insert(attachment *model.Attachment) (*model.Attachment, error) {
	return s.r.Insert(attachment)
}

func (s *AttachmentService) Delete(userID uuid.UUID, attachmentID uuid.UUID) (*model.Attachment, error) {
	att, err := s.FindByID(userID, attachmentID)
	if err != nil {
		return nil, err
	}

	if err := s.r.Delete(attachmentID); err != nil {
		return nil, err
	}

	return att, s.DeleteFromDisk(att)
}

func (s *AttachmentService) GetFilePath(attachment *model.Attachment) string {
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

func (s *AttachmentService) DeleteFromDisk(attachment *model.Attachment) error {
	path := s.GetFilePath(attachment)
	return os.Remove(path)
}
