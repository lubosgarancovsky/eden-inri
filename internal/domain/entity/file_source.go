package entity

import "io"

type FileSource struct {
	Name     string
	MimeType string
	Reader   io.Reader
	Size     int64
}
