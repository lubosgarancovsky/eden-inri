package helpers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/filter"
	"github.com/lubosgarancovsky/go-kit/list"
	"github.com/lubosgarancovsky/go-kit/rsql"
	"github.com/lubosgarancovsky/go-kit/sort"
)

func CreateListingQuery(c *gin.Context, parser *rsql.Parser, filterMap map[string]string, sortMap map[string]string) *list.ListingQuery {
	var qp list.QueryParms
	err := c.ShouldBindQuery(&qp)
	if err != nil {
		c.Error(api_err.Wrap(api_err.ErrBadRequest, err).WithMessage("Invalid query parameters"))
		return nil
	}

	var limit = 10
	if qp.PageSize > 0 {
		limit = qp.PageSize
	}

	var page = 1
	if qp.Page > 0 {
		page = qp.Page
	}

	lq := &list.ListingQuery{
		Filter: nil,
		Sort:   []sort.Sort{},
		Limit:  limit,
		Offset: (page - 1) * limit,
		Page:   page,
	}

	if qp.Filter != "" {
		ast, err := parser.Parse(qp.Filter)
		if err != nil {
			c.Error(api_err.Wrap(api_err.ErrBadRequest, err).WithMessage("Invalid filter parameter"))
			return nil
		}

		fil, err := filter.BuildFilter(ast, filterMap)
		if err != nil {
			c.Error(api_err.Wrap(api_err.ErrBadRequest, err).WithMessage("Invalid filter parameter"))
			return nil
		}

		lq.Filter = fil
	}

	if qp.Sort != "" {
		srt, err := sort.BuildSort(qp.Sort, sortMap)
		if err != nil {
			c.Error(api_err.Wrap(api_err.ErrBadRequest, err).WithMessage("Invalid sort parameter"))
			return nil
		}

		lq.Sort = srt
	}

	return lq
}

func ExtractID(c *gin.Context, name string) uuid.UUID {
	ID, ok := c.Params.Get(name)
	if !ok {
		c.Error(api_err.ErrParameterMissing.WithMessage(fmt.Sprintf("Path parameter %s is missing", name)))
		return uuid.Nil
	}

	UID, err := uuid.Parse(ID)
	if err != nil {
		c.Error(api_err.Wrap(api_err.ErrInvalidUUID.WithMessage(fmt.Sprintf("%s is not a valid UUID", ID)), err))
		return uuid.Nil
	}

	return UID
}

func GetUserContext(c *gin.Context) *model.UserContext {
	user, _ := c.Get("user")
	return user.(*model.UserContext)
}

func HandleList[T any](c *gin.Context, parser *rsql.Parser, config types.ListConfig, fn func(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]T, int64, error)) {
	lq := CreateListingQuery(c, parser, config.Filter, config.Sort)
	user := GetUserContext(c)

	items, totalCount, err := fn(c.Request.Context(), user.ID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	result := &list.Page[T]{
		Items:      *items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}

	c.JSON(200, result)
}

func HandleFindByID[T any](c *gin.Context, paramName string, fn func(ctx context.Context, userID, resourceID uuid.UUID) (T, error)) {
	resourceID := ExtractID(c, paramName)
	user := GetUserContext(c)

	result, err := fn(c.Request.Context(), user.ID, resourceID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}

func HandleCreate[T any, R any](c *gin.Context, fn func(ctx context.Context, userID uuid.UUID, payload *T) (R, error)) {
	user := GetUserContext(c)

	var input T
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := fn(c.Request.Context(), user.ID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
}

func HandleUpdate[T any, R any](c *gin.Context, paramName string, fn func(ctx context.Context, userID, resourceID uuid.UUID, payload *T) (R, error)) {
	user := GetUserContext(c)
	resourceID := ExtractID(c, paramName)

	var input T
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := fn(c.Request.Context(), user.ID, resourceID, &input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(201, result)
}

func HandleDelete(c *gin.Context, paramName string, fn func(ctx context.Context, userID, resourceID uuid.UUID) error) {
	user := GetUserContext(c)
	resourceID := ExtractID(c, paramName)

	if err := fn(c.Request.Context(), user.ID, resourceID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(204, gin.H{})
}
