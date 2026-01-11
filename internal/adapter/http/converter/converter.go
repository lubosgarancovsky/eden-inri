package converter

import (
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToDeleteCommand(input interface{}) (*command.DeleteCommand, error) {
	baseDto, ok := input.(dto.BaseDto)
	if !ok {
		return nil, go_kit.ErrInternalServer // TODO: improve error handling
	}
	return &command.DeleteCommand{
		ID:     baseDto.ID,
		UserID: baseDto.UserID,
	}, nil
}

func ToFindByIDQuery(input interface{}) (*query.FindByIDQuery, error) {
	baseDto, ok := input.(dto.BaseDto)
	if !ok {
		return nil, go_kit.ErrInternalServer
	}
	return &query.FindByIDQuery{
		ID:     baseDto.ID,
		UserID: baseDto.UserID,
	}, nil
}

func ToListQuery(input interface{}, lq *go_kit.ListingQuery) (*query.ListQuery, error) {
	userIDDto, ok := input.(dto.UserIDDto)
	if !ok {
		return nil, go_kit.ErrInternalServer
	}
	return &query.ListQuery{
		UserID:       userIDDto.UserID,
		ListingQuery: lq,
	}, nil
}

func ToListResponse[I any, O any](
	items *[]I,
	mapper func(*I) O,
) []O {
	if items == nil {
		return nil
	}

	in := *items
	out := make([]O, len(in))

	for i := range in {
		out[i] = mapper(&in[i])
	}

	return out
}
