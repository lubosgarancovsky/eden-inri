package command

import (
	"mime/multipart"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type FileCommand struct {
	Source entity.FileSource
}

type FilesCommand struct {
	Sources []*entity.FileSource
}

type FileHeaderCommand struct {
	Source *multipart.FileHeader
}
