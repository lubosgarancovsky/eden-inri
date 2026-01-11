package dto

import "time"

type AttachmentRes struct {
	ID           string    `json:"id"`
	Model        string    `json:"model"`
	ModelID      string    `json:"modelId"`
	OriginalName string    `json:"originalName"`
	ServerName   string    `json:"serverName"`
	MimeType     string    `json:"mimeType"`
	Size         int64     `json:"size"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ListAttachmentsReq struct {
	UserID string `header:"X-User-ID"`
}

type FindAttachmentByIDReq struct {
	UserID       string `header:"X-User-ID"`
	AttachmentID string `uri:"attachmentId"`
}

type DeleteAttachmentReq struct {
	UserID       string `header:"X-User-ID"`
	AttachmentID string `uri:"attachmentId"`
}
