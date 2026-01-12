package handle

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/config"
	"github.com/lubosgarancovsky/go-kit"
)

type AuthContext struct {
	UserID        string  `header:"X-User-ID" binding:"required"`
	UserRole      string  `header:"X-User-Role" binding:"required,oneof=ADMIN USER"`
	Authorization *string `header:"Authorization"`
}

func Error(c *gin.Context, err error) {
	requestID := c.GetHeader("X-Request-ID")
	if requestID == "" {
		requestID = uuid.New().String()
	}

	var apiErr *go_kit.ApiError
	if errors.As(err, &apiErr) {
		apiErr.Log()
		c.JSON(apiErr.HTTPStatus, apiErr.ToJSON(config.GlobalConfig.ServiceName, requestID))
		c.Abort()
		return
	}

	internalErr := go_kit.Unknown(err)
	internalErr.Log()
	c.JSON(500, internalErr.ToJSON(config.GlobalConfig.ServiceName, requestID))
	c.Abort()
}

func ListingQuery(c *gin.Context, parser *go_kit.Parser, listingAttr interface{}) *go_kit.ListingQuery {
	qp := &go_kit.QueryParams{}
	if err := c.ShouldBindQuery(&qp); err != nil {
		Error(c, go_kit.Wrap(go_kit.ErrBadRequest, err).WithMessage("Invalid query parameters"))
		return nil
	}

	lq, err := go_kit.NewListingQuery(qp, parser, listingAttr)
	if err != nil {
		Error(c, err)
		return nil
	}

	if lq == nil {
		Error(c, go_kit.ErrBadRequest.WithMessage("Invalid listing query"))
		return nil
	}

	return lq
}

func Page[T any](c *gin.Context, lq *go_kit.ListingQuery, totalCount int64, items []T) {
	c.JSON(http.StatusOK, &go_kit.Page[T]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	})
}

func UserID(c *gin.Context) uuid.UUID {
	user, ok := c.Get("user")
	if !ok {
		Error(c, go_kit.ErrUnauthorized)
		return uuid.Nil
	}

	userCtx, ok := user.(*AuthContext)
	if !ok {
		Error(c, go_kit.ErrUnauthorized)
		return uuid.Nil
	}

	userID, err := uuid.Parse(userCtx.UserID)
	if err != nil {
		Error(c, go_kit.ErrUnauthorized)
	}

	return userID
}
