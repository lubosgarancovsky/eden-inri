package command

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type CreateStoryActivityCommand struct {
	StoryID uuid.UUID
	ActorID uuid.UUID
	Type    string
	Payload json.RawMessage
}

func (c *CreateStoryActivityCommand) ToDomain() *entity.StoryActivity {
	return &entity.StoryActivity{
		ID:        uuid.New(),
		StoryID:   c.StoryID,
		ActorID:   c.ActorID,
		Type:      entity.ActivityType(c.Type),
		Payload:   c.Payload,
		CreatedAt: time.Now(),
	}
}

type UpdateStoryActivityCommand struct {
	ID      uuid.UUID
	ActorID uuid.UUID
	Payload json.RawMessage
}

func (c *UpdateStoryActivityCommand) Apply(e *entity.StoryActivity) {
	e.Payload = c.Payload
}
