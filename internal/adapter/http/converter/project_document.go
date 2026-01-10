package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

func ToProjectDocumentEntityFromCreate(req *dto.CreateProjectDocumentReq) *entity.ProjectDocument {
	return &entity.ProjectDocument{
		ID:        uuid.New(),
		ProjectID: req.ProjectID,
		Name:      req.Name,
		Content:   req.Content,
		Tags:      req.Tags,
	}
}

func ToProjectDocumentEntityFromUpdate(req *dto.UpdateProjectDocumentReq) *entity.ProjectDocument {
	return &entity.ProjectDocument{
		ID:        req.DocumentID,
		ProjectID: req.ProjectID,
		Name:      req.Name,
		Content:   req.Content,
		Tags:      req.Tags,
	}
}

func ToProjectDocumentResponse(d *entity.ProjectDocument) *dto.ProjectDocumentRes {
	return &dto.ProjectDocumentRes{
		ID:        d.ID,
		ProjectID: d.ProjectID,
		Name:      d.Name,
		Content:   d.Content,
		Tags:      d.Tags,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func ToProjectDocumentListResponse(list *[]entity.ProjectDocument) []*dto.ProjectDocumentRes {
	if list == nil {
		return nil
	}
	arr := *list
	res := make([]*dto.ProjectDocumentRes, len(arr))
	for i := range arr {
		res[i] = ToProjectDocumentResponse(&arr[i])
	}
	return res
}

// Helper to pass through listing query
func ToProjectDocumentListingQuery(lq *go_kit.ListingQuery) *go_kit.ListingQuery { return lq }
