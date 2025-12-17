package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/listing"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

type StoryActivityHandler struct {
	parser *rsql.Parser
	s      *service.StoryActivityService
}

type ActivitiesPage struct {
	Items      []model.StoryActivity
	Page       int
	PageSize   int
	TotalCount int64
}

func NewStoryActivityHandler(p *rsql.Parser, s *service.StoryActivityService) *StoryActivityHandler {
	return &StoryActivityHandler{s: s, parser: p}
}

// InsertActivity @Summary      Add activity to a story
// @Description  Records a new activity for a story (e.g., label added, state changed, comment added)
// @Tags         Story Activities
// @Accept       json
// @Produce      json
// @Param        kanbanId  path  string  true "Kanban ID"
// @Param        storyId    path  string  true "Story ID"
// @Param        body       body  map[string]interface{} true "Activity payload including eventType and optional data"
// @Success      201  {object}  model.StoryActivity
// @Router       /v1/inri/kanban/{kanbanId}/stories/{storyId}/activities [post]
func (h *StoryActivityHandler) InsertActivity(c *gin.Context) {
	kanbanID, err := helpers.ExtractID(c, "kanbanId")
	if err != nil {
		c.Error(err)
		return
	}
	storyID, err := helpers.ExtractID(c, "storyId")
	if err != nil {
		c.Error(err)
		return
	}
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	var req model.StoryActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	activity, err := h.s.InsertActivity(user.ID, kanbanID, storyID, req.Type, req.Payload)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(201, activity)
}

// ListActivities @Summary      List story activities
// @Description  Returns all activity events for a story (paginated)
// @Tags         Story Activities
// @Accept       json
// @Produce      json
// @Param        kanbanId    path  string  true "Kanban ID"
// @Param        storyId    path  string  true "Story ID"
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}  ActivitiesPage
// @Router       /v1/inri/kanban/{kanbanId}/stories/{storyId}/activities [get]
func (h *StoryActivityHandler) ListActivities(c *gin.Context) {
	kanbanID, err := helpers.ExtractID(c, "kanbanId")
	if err != nil {
		c.Error(err)
		return
	}
	storyID, err := helpers.ExtractID(c, "storyId")
	if err != nil {
		c.Error(err)
		return
	}
	user, err := helpers.GetUserContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	lq, apiErr := helpers.CreateListingQuery(c, h.parser, listing.InvoiceFilter, listing.InvoiceSort)
	if apiErr != nil {
		c.Error(apiErr)
		return
	}

	result, err := h.s.ListActivities(user.ID, kanbanID, storyID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}
