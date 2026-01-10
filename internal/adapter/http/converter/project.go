package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/query"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToCreateProjectCommand(input *dto.CreateProjectReq, userID uuid.UUID) *command.CreateProjectCommand {
	return &command.CreateProjectCommand{
		UserID:      userID,
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status,
		Tags:        input.Tags,
		Slug:        input.Slug,
	}
}

func ToUpdateProjectCommand(input *dto.UpdateProjectReq, userID uuid.UUID) *command.UpdateProjectCommand {
	return &command.UpdateProjectCommand{
		UserID:      userID,
		ProjectID:   input.ProjectID,
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status,
		Tags:        input.Tags,
	}
}

func ToDeleteProjectCommand(input *dto.DeleteProjectReq) *command.DeleteProjectCommand {
	return &command.DeleteProjectCommand{ProjectID: input.ProjectID, UserID: input.UserID}
}

func ToFindProjectByIDQuery(req *dto.FindProjectByIDReq) *query.FindProjectByIDQuery {
	return &query.FindProjectByIDQuery{UserID: req.UserID, ProjectID: req.ProjectID}
}

func ToListProjectsQuery(req *dto.ListProjectsReq, lq *go_kit.ListingQuery) *query.ListProjectsQuery {
	return &query.ListProjectsQuery{UserID: req.UserID, ListingQuery: *lq}
}

func ToProjectResponse(p *entity.Project) *dto.ProjectRes {
	return &dto.ProjectRes{
		ID:             p.ID,
		Slug:           p.Slug,
		Name:           p.Name,
		Description:    p.Description,
		Status:         string(p.Status),
		Tags:           p.Tags,
		StorySequence:  p.StorySequence,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		LastActivityAt: p.LastActivityAt,
		IsStarred:      p.IsStarred,
	}
}

func ToProjectListResponse(projects *[]entity.Project) []*dto.ProjectRes {
	if projects == nil {
		return nil
	}
	arr := *projects
	res := make([]*dto.ProjectRes, len(arr))
	for i := range arr {
		res[i] = ToProjectResponse(&arr[i])
	}
	return res
}
