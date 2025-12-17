package listing

var StoryFilter = map[string]string{
	"id":         "id",
	"projectId":  "project_id",
	"boardId":    "board_id",
	"columnId":   "column_id",
	"slug":       "slug",
	"title":      "title",
	"kind":       "kind",
	"assigneeId": "assignee_id",
	"priority":   "priority",
	"startDate":  "start_date",
	"endDate":    "end_date",
	"createdAt":  "created_at",
}

var StorySort = map[string]string{
	"position":  "position",
	"priority":  "priority",
	"createdAt": "created_at",
	"updatedAt": "updated_at",
	"title":     "title",
	"slug":      "slug",
}
