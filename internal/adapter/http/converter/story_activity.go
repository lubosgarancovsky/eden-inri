package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateStoryActivityCommand(input *dto.CreateStoryActivityReq) *command.CreateStoryActivityCommand {
	return &command.CreateStoryActivityCommand{
		StoryID: uuid.MustParse(input.StoryID),
		ActorID: uuid.MustParse(input.UserID),
		Type:    input.Type,
		Payload: input.Payload,
	}
}

func ToUpdateStoryActivityCommand(input *dto.UpdateStoryActivityReq) *command.UpdateStoryActivityCommand {
	return &command.UpdateStoryActivityCommand{
		ID:      uuid.MustParse(input.ActivityID),
		ActorID: uuid.MustParse(input.UserID),
		Payload: input.Payload,
	}
}

func ToDeleteStoryActivityCommand(input *dto.DeleteStoryActivityReq) *command.DeleteCommand {
	return &command.DeleteCommand{
		ID:     uuid.MustParse(input.ActivityID),
		UserID: uuid.MustParse(input.UserID),
	}
}

func ToListStoryActivitiesQuery(input *dto.ListStoryActivitiesReq, lq *go_kit.ListingQuery) *query.ListStoryActivitiesQuery {
	return &query.ListStoryActivitiesQuery{
		StoryID:      uuid.MustParse(input.StoryID),
		UserID:       uuid.MustParse(input.UserID),
		ListingQuery: lq,
	}
}

func ToStoryActivityResponse(e *entity.StoryActivity) *dto.StoryActivityRes {
	res := &dto.StoryActivityRes{
		ID:        e.ID,
		StoryID:   e.StoryID,
		ActorID:   e.ActorID,
		Type:      string(e.Type),
		Payload:   e.Payload,
		CreatedAt: e.CreatedAt,
	}

	// Assuming we can convert entity.User to dto.UserRes
	res.Actor = ToUserResponse(&e.Actor)

	return res
}
