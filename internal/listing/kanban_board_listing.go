package listing

var KanbanBoardFilter = map[string]string{
	"id":             "id",
	"name":           "name",
	"status":         "status",
	"projectId":      "project_id",
	"createdAt":      "created_at",
	"lastActivityAt": "last_activity_at",
}

var KanbanBoardSort = map[string]string{
	"name":           "name",
	"createdAt":      "created_at",
	"lastActivityAt": "last_activity_at",
}
