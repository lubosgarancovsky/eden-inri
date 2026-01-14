package dto

type ListStoryAttachmentsReq struct {
	UserID  string `header:"X-User-ID"`
	ScopeID string `uri:"storyId"`
}

type StoryAttachmentByIDReq struct {
	UserID  string `header:"X-User-ID"`
	ScopeID string `uri:"storyId"`
	ID      string `uri:"attachmentId"`
}
