package converter

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	"github.com/lubosgarancovsky/go-kit"
)

// ToCommand Requires DTO in shape { ID, UserID }
func ToCommand(input interface{}) (*command.Command, error) {
	userID := extractUserID(input)
	if userID == uuid.Nil {
		return nil, go_kit.ErrUnauthorized
	}

	ID := extractID(input, "ID")
	if ID == uuid.Nil {
		return nil, go_kit.ErrBadRequest
	}

	return &command.Command{
		ID:     ID,
		UserID: userID,
	}, nil
}

// ToScopedCommand Requires DTO in shape { ID, ScopeID, UserID }
func ToScopedCommand(input interface{}) (*command.ScopedCommand, error) {
	userID := extractUserID(input)
	if userID == uuid.Nil {
		return nil, go_kit.ErrUnauthorized
	}

	ID := extractID(input, "ID")
	if ID == uuid.Nil {
		return nil, go_kit.ErrBadRequest
	}

	ScopeID := extractID(input, "ScopeID")
	if ScopeID == uuid.Nil {
		return nil, go_kit.ErrBadRequest
	}

	return &command.ScopedCommand{
		ID:      ID,
		ScopeID: ScopeID,
		UserID:  userID,
	}, nil
}

// ToQuery Requires DTO in shape { ID, UserID }
func ToQuery(input interface{}) (*query.Query, error) {
	userID := extractUserID(input)
	if userID == uuid.Nil {
		return nil, go_kit.ErrUnauthorized
	}

	ID := extractID(input, "ID")
	if ID == uuid.Nil {
		return nil, go_kit.ErrBadRequest
	}

	return &query.Query{
		ID:     ID,
		UserID: userID,
	}, nil
}

// ToScopedQuery Requires DTO in shape { ID, ScopeID, UserID }
func ToScopedQuery(input interface{}) (*query.ScopedQuery, error) {
	userID := extractUserID(input)
	if userID == uuid.Nil {
		return nil, go_kit.ErrUnauthorized
	}

	ID := extractID(input, "ID")
	if ID == uuid.Nil {
		return nil, go_kit.ErrBadRequest
	}

	ScopeID := extractID(input, "ScopeID")
	if ScopeID == uuid.Nil {
		return nil, go_kit.ErrBadRequest
	}

	return &query.ScopedQuery{
		ID:      ID,
		ScopeID: ScopeID,
		UserID:  userID,
	}, nil
}

// ToListQuery Requires DTO in shape { UserID }
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

// ToScopedListQuery Requires DTO in shape { UserID, ScopeID }
func ToScopedListQuery(input interface{}, lq *go_kit.ListingQuery) (*query.ScopedListQuery, error) {
	userID := extractUserID(input)
	if userID == uuid.Nil {
		return nil, go_kit.ErrUnauthorized
	}

	ScopeID := extractID(input, "ScopeID")
	if ScopeID == uuid.Nil {
		return nil, go_kit.ErrBadRequest
	}

	return &query.ScopedListQuery{
		UserID:       userID,
		ScopeID:      ScopeID,
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

	id, err := uuid.Parse(field.String())
	if err != nil {
		return uuid.Nil
	}

	return id
}

func extractID(v interface{}, name string) uuid.UUID {
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

	field := rv.FieldByName(name)
	if !field.IsValid() || field.Kind() != reflect.String {
		return uuid.Nil
	}

	id, err := uuid.Parse(field.String())
	if err != nil {
		return uuid.Nil
	}

	return id
}
