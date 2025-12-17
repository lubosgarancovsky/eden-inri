package listing

var ProjectUserFilter = map[string]string{
	"projectId":     "project_id",
	"userId":        "user_id",
	"role":          "role",
	"joinedAt":      "joined_at",
	"userFirstName": "User.first_name",
	"userLastName":  "User.first_name",
}

var ProjectUserSort = map[string]string{
	"role":          "role",
	"joinedAt":      "joined_at",
	"userFirstName": "User.first_name",
	"userLastName":  "User.first_name",
}
