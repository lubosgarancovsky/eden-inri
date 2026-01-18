package dto

type StoryLabelReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	StoryID   string `uri:"storyId"`
	LabelID   string `uri:"labelId"`
}

type ListStoryLabelReq struct {
	UserID    string `header:"X-User-ID"`
	ProjectID string `uri:"projectId"`
	StoryID   string `uri:"storyId"`
}

type StoryLabelRes struct {
	ProjectLabelRes
}
