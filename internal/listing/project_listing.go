package listing

var ProjectFilter = map[string]string{
	"name":           "name",
	"status":         "status",
	"slug":           "slug",
	"lastActivityAt": "last_activity_at",
	"isStarred":      "is_starred",
}

var ProjectSort = map[string]string{
	"name":           "name",
	"createdAt":      "created_at",
	"updatedAt":      "updated_at",
	"lastActivityAt": "last_activity_at",
}
