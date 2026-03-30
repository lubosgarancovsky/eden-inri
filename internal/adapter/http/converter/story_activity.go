package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	"github.com/lubosgarancovsky/go-kit"
)

func ToCreateStoryActivityCommand(input *dto.CreateStoryActivityReq) (*command.CreateStoryActivityCommand, error) {
	storyID, err := uuid.Parse(input.StoryID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	actorID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.CreateStoryActivityCommand{
		StoryID: storyID,
		ActorID: actorID,
		Type:    input.Type,
		Payload: input.Payload,
	}, nil
}

func ToUpdateStoryActivityCommand(input *dto.UpdateStoryActivityReq) (*command.UpdateStoryActivityCommand, error) {
	id, err := uuid.Parse(input.ActivityID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	actorID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.UpdateStoryActivityCommand{
		ID:      id,
		ActorID: actorID,
		Payload: input.Payload,
	}, nil
}

func ToDeleteStoryActivityCommand(input *dto.DeleteStoryActivityReq) (*command.Command, error) {
	id, err := uuid.Parse(input.ActivityID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &command.Command{
		ID:     id,
		UserID: userID,
	}, nil
}

func ToListStoryActivitiesQuery(input *dto.ListStoryActivitiesReq, lq *go_kit.ListingQuery) (*query.ListStoryActivitiesQuery, error) {
	storyID, err := uuid.Parse(input.StoryID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}
	return &query.ListStoryActivitiesQuery{
		StoryID:      storyID,
		UserID:       userID,
		ListingQuery: lq,
	}, nil
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
