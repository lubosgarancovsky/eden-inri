package dto

import "time"

type AttachmentRes struct {
	ID           string    `json:"id"`
	Model        string    `json:"model"`
	ModelID      string    `json:"modelId"`
	OriginalName string    `json:"originalName"`
	MimeType     string    `json:"mimeType"`
	Size         int64     `json:"size"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ListAttachmentsReq struct {
	UserID string `header:"X-User-ID"`
}

type AttachmentByIDReq struct {
	UserID string `header:"X-User-ID"`
	ID     string `uri:"attachmentId"`
}

type UpdateAttachmentReq struct {
	UserID string `header:"X-User-ID"`
	ID     string `uri:"attachmentId"`
	Name   string `json:"name"`
}

type UploadAttachmentReq struct {
	UserID  string
	ModelID string
	Model   string
}
