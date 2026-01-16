package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToListStoryAttachmentsQuery(input *dto.ListStoryAttachmentsReq, lq *go_kit.ListingQuery) (*query.ListStoryAttachmentsQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectId, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	storyID, err := uuid.Parse(input.StoryId)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &query.ListStoryAttachmentsQuery{
		UserID:       userID,
		ProjectID:    projectId,
		StoryID:      storyID,
		ListingQuery: lq,
	}, nil
}

func ToFindStoryAttachmentByIDQuery(input *dto.StoryAttachmentByIDReq) (*query.FindStoryAttachmentByIDQuery, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectId, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	storyID, err := uuid.Parse(input.StoryID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	attachmentID, err := uuid.Parse(input.AttachmentID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &query.FindStoryAttachmentByIDQuery{
		UserID:       userID,
		ProjectID:    projectId,
		StoryID:      storyID,
		AttachmentID: attachmentID,
	}, nil
}

func ToDeleteStoryAttachmentCommand(input *dto.StoryAttachmentByIDReq) (*command.DeleteStoryAttachmentCommand, error) {
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	projectId, err := uuid.Parse(input.ProjectID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	storyID, err := uuid.Parse(input.StoryID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	attachmentID, err := uuid.Parse(input.AttachmentID)
	if err != nil {
		return nil, go_kit.ErrInvalidUUID
	}

	return &command.DeleteStoryAttachmentCommand{
		UserID:       userID,
		ProjectID:    projectId,
		StoryID:      storyID,
		AttachmentID: attachmentID,
	}, nil
}
