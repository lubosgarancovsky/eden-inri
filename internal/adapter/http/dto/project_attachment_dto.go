package dto

type ListProjectAttachmentsReq struct {
	UserID  string `header:"X-User-ID"`
	ScopeID string `uri:"projectId"`
}

type ProjectAttachmentByIDReq struct {
	UserID  string `header:"X-User-ID"`
	ScopeID string `uri:"projectId"`
	ID      string `uri:"attachmentId"`
}

type UpdateProjectAttachmentReq struct {
	UserID       string `header:"X-User-ID"`
	ProjectID    string `uri:"projectId"`
	AttachmentID string `uri:"attachmentId"`
	Name         string `json:"name"`
}
