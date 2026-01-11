package attachment

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type AttachmentUtilityService struct {
	repo ports.PersistAttachmentPort
	cfg  *config.Config
}

func NewAttachmentUtilityService(repo ports.PersistAttachmentPort, cfg *config.Config) *AttachmentUtilityService {
	return &AttachmentUtilityService{repo: repo, cfg: cfg}
}

func (s *AttachmentUtilityService) GetFilePath(attachment *entity.Attachment) string {
	return filepath.Join(s.cfg.UploadsFolder, attachment.UserID.String(), attachment.Model, attachment.ModelID, attachment.OriginalName)
}

func (s *AttachmentUtilityService) GetFileName(originalName string) string {
	return strings.ToLower(strings.TrimSpace(originalName))
}

func (s *AttachmentUtilityService) GetMimeType(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := make([]byte, 512)
	n, _ := file.Read(buf)
	return http.DetectContentType(buf[:n]), nil
}

func (s *AttachmentUtilityService) SaveAttachment(c *gin.Context, userID uuid.UUID, modelName string, modelID string, file multipart.FileHeader) (*entity.Attachment, error) {
	mime, err := s.GetMimeType(&file)
	if err != nil {
		return nil, err
	}

	attachment := &entity.Attachment{
		ID:           uuid.New(),
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

	if err := s.repo.Create(c.Request.Context(), attachment); err != nil {
		return nil, err
	}

	return attachment, nil
}

func (s *AttachmentUtilityService) SaveFile(c *gin.Context, attachment *entity.Attachment, file multipart.FileHeader) error {
	path := s.GetFilePath(attachment)
	return c.SaveUploadedFile(&file, path)
}
