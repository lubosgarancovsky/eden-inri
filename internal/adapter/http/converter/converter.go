package converter

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	"github.com/lubosgarancovsky/go-kit"
)

func ToDeleteCommand(input interface{}) (*command.DeleteCommand, error) {
	userID := extractUserID(input)
	if userID == uuid.Nil {
		return nil, go_kit.ErrUnauthorized
	}

	ID := extractID(input)
	if ID == uuid.Nil {
		return nil, go_kit.ErrBadRequest
	}

	return &command.DeleteCommand{
		ID:     ID,
		UserID: userID,
	}, nil
}

func ToFindByIDQuery(input interface{}) (*query.FindByIDQuery, error) {
	userID := extractUserID(input)
	if userID == uuid.Nil {
		return nil, go_kit.ErrUnauthorized
	}

	ID := extractID(input)
	if ID == uuid.Nil {
		return nil, go_kit.ErrBadRequest
	}

	return &query.FindByIDQuery{
		ID:     ID,
		UserID: userID,
	}, nil
}

func ToListQuery(input interface{}, lq *go_kit.ListingQuery) (*query.ListQuery, error) {
	userID := extractUserID(input)
	if userID == uuid.Nil {
		return nil, go_kit.ErrUnauthorized
	}

	return &query.ListQuery{
		UserID:       userID,
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

func extractUserID(v interface{}) uuid.UUID {
	rv := reflect.ValueOf(v)

	if !rv.IsValid() {
		return uuid.Nil
	}

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return uuid.Nil
	}

	field := rv.FieldByName("UserID")
	if !field.IsValid() || field.Kind() != reflect.String {
		return uuid.Nil
	}

	return uuid.MustParse(field.String())
}

func extractID(v interface{}) uuid.UUID {
	rv := reflect.ValueOf(v)

	if !rv.IsValid() {
		return uuid.Nil
	}

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return uuid.Nil
	}

	field := rv.FieldByName("ID")
	if !field.IsValid() || field.Kind() != reflect.String {
		return uuid.Nil
	}

	return uuid.MustParse(field.String())
}
