package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/service"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/eden-inri/pkg/types"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/rsql"
)

var StoryActivityListConfig = &types.ListConfig{
	Filter: map[string]string{
		"id":        "id",
		"storyId":   "story_id",
		"actorId":   "actor_id",
		"type":      "type",
		"createdAt": "created_at",
		"actorName": "users.name",
	},
	Sort: map[string]string{
		"createdAt": "created_at",
	},
}

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
// @Param        projectId  path  string  true "Project ID"
// @Param        storyId    path  string  true "Story ID"
// @Param        body       body  map[string]interface{} true "Activity payload including eventType and optional data"
// @Success      201  {object}  model.StoryActivity
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/activities [post]
// @security GatewayAuth
func (h *StoryActivityHandler) InsertActivity(c *gin.Context) {
	storyID := helpers.ExtractID(c, "storyId")
	userID := helpers.GetUserContext(c).ID

	var req model.StoryActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_err.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	activity, err := h.s.InsertActivity(c.Request.Context(), userID, storyID, req.Type, req.Payload)
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
// @Param        projectId  path  string  true "Project ID"
// @Param        storyId    path  string  true "Story ID"
// @Param        page      query     int     false  "Page number"     default(1)
// @Param        pageSize  query     int     false  "Items per page"  default(10)
// @Param        filter      query     string     false  "RSQL filter query"
// @Param        sort  query     string     false  "Sort query"
// @Success      200  {object}  ActivitiesPage
// @Router       /v1/inri/projects/{projectId}/stories/{storyId}/activities [get]
// @security GatewayAuth
func (h *StoryActivityHandler) ListActivities(c *gin.Context) {
	storyID := helpers.ExtractID(c, "storyId")

	lq := helpers.CreateListingQuery(c, h.parser, StoryActivityListConfig.Filter, StoryActivityListConfig.Sort)

	result, err := h.s.ListActivities(c.Request.Context(), storyID, lq)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)
}
