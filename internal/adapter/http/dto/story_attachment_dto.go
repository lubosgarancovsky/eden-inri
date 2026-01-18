package dto

type ListStoryAttachmentsReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	StoryId   string `uri:"storyId"`
}

type StoryAttachmentByIDReq struct {
	UserID       string `header:"X-User-ID"`
	ProjectID    string `uri:"projectId"`
	StoryID      string `uri:"storyId"`
	AttachmentID string `uri:"attachmentId"`
}
