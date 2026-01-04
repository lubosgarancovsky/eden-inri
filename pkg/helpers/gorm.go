package helpers

import (
	"fmt"

	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
)

type ListState[T any] struct {
	items      []T
	totalCount int64
	lq         *list.ListingQuery
}

func OrderBy[T any](query *gorm.DB, state *ListState[T]) *gorm.DB {
	for _, clause := range state.lq.Sort {
		query = query.Order(fmt.Sprintf("%s %s", clause.Field, clause.Direction))
	}
	return query
}

func Paginate[T any](query *gorm.DB, state *ListState[T]) error {
	q := OrderBy(query, state)
	if err := q.Limit(state.lq.Limit).Offset(state.lq.Offset).Find(&state.items).Error; err != nil {
		return err
	}

	return nil
}

func Count[T any](query *gorm.DB, state *ListState[T]) error {
	if err := query.Count(&state.totalCount).Error; err != nil {
		return err
	}

	return nil
}

func List[T any](query *gorm.DB, lq *list.ListingQuery) ([]T, int64, error) {
	state := &ListState[T]{
		items:      make([]T, 0),
		totalCount: 0,
		lq:         lq,
	}

	if err := Paginate(query.WithContext(query.Statement.Context).Session(&gorm.Session{}), state); err != nil {
		return nil, 0, err
	}

	if err := Count(query.WithContext(query.Statement.Context).Session(&gorm.Session{}), state); err != nil {
		return nil, 0, err
	}

	return state.items, state.totalCount, nil
}
