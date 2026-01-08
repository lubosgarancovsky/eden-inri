package domain

import (
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Model        string
	ModelID      string
	OriginalName string
	ServerName   string
	MimeType     string
	Size         int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewAttachmentFromFile(
	userID uuid.UUID,
	modelName string,
	modelID string,
	file *multipart.FileHeader,
) (*Attachment, error) {
	mimeType, err := getAttachmentMimeType(file)
	if err != nil {
		return nil, err
	}

	serverName := getAttachmentServerName(file.Filename)
	now := time.Now()

	return &Attachment{
		ID:           uuid.New(),
		UserID:       userID,
		Model:        modelName,
		ModelID:      modelID,
		OriginalName: file.Filename,
		ServerName:   serverName,
		MimeType:     mimeType,
		Size:         file.Size,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil

}

func getAttachmentMimeType(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func(file multipart.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	buf := make([]byte, 512)
	n, _ := file.Read(buf)
	return http.DetectContentType(buf[:n]), nil
}

func getAttachmentServerName(originalName string) string {
	return strings.ToLower(strings.TrimSpace(originalName))
}
