package entity

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	go_kit "github.com/lubosgarancovsky/go-kit"
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
	TempFilePath string
	File         *multipart.FileHeader
}

func NewAttachment(userID uuid.UUID, modelName string, modelID string, file *multipart.FileHeader) (*Attachment, error) {
	if file == nil {
		return nil, go_kit.ErrBadRequest.WithMessage("File is missing")
	}

	mimeType, err := getMimeType(file)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	att := &Attachment{
		ID:           uuid.New(),
		UserID:       userID,
		Model:        modelName,
		ModelID:      modelID,
		OriginalName: file.Filename,
		ServerName:   file.Filename,
		MimeType:     mimeType,
		Size:         file.Size,
		CreatedAt:    now,
		UpdatedAt:    now,
		File:         file,
	}

	att.ServerName = att.GenerateServerName()
	return att, nil
}

// GetFilePath Computes a path where a file is saved on a server disk
func (a *Attachment) GetFilePath(uploadFolder string) string {
	return filepath.Join(uploadFolder, a.UserID.String(), a.Model, a.ModelID, a.ServerName)
}

func (a *Attachment) GenerateServerName() string {
	now := time.Now().Unix()
	uid := uuid.New().String()
	return fmt.Sprintf("%d_%s", now, strings.ReplaceAll(uid, "-", ""))
}

func getMimeType(fileHeader *multipart.FileHeader) (string, error) {
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
