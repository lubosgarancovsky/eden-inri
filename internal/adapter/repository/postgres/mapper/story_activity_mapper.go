package mapper

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

func StoryActivityFromDomain(e *entity.StoryActivity) *model.StoryActivity {
	return &model.StoryActivity{
		ID:        e.ID,
		StoryID:   e.StoryID,
		ActorID:   e.ActorID,
		Type:      string(e.Type),
		Payload:   e.Payload,
		CreatedAt: e.CreatedAt,
	}
}
