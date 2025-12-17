package listing

var StoryActivityFilter = map[string]string{
	"id":        "id",
	"storyId":   "story_id",
	"actorId":   "actor_id",
	"type":      "type",
	"createdAt": "created_at",
	"actorName": "users.name", // joined
}

var StoryActivitySort = map[string]string{
	"createdAt": "created_at",
}
