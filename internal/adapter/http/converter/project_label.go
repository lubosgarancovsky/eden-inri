package converter

import (
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
)

func ToLabelPortFromCreate(req *dto.CreateLabelReq) *ports.Label {
	return &ports.Label{ID: uuid.New(), ProjectID: req.ProjectID, Name: req.Name, Description: req.Description, Color: req.Color}
}

func ToLabelPortFromUpdate(req *dto.UpdateLabelReq) *ports.Label {
	return &ports.Label{ID: req.LabelID, ProjectID: req.ProjectID, Name: req.Name, Description: req.Description, Color: req.Color}
}

func ToLabelResponse(l *ports.Label) *dto.LabelRes {
	return &dto.LabelRes{ID: l.ID, ProjectID: l.ProjectID, Name: l.Name, Description: l.Description, Color: l.Color}
}

func ToLabelListResponse(list *[]ports.Label) []*dto.LabelRes {
	if list == nil {
		return nil
	}
	arr := *list
	res := make([]*dto.LabelRes, len(arr))
	for i := range arr {
		res[i] = ToLabelResponse(&arr[i])
	}
	return res
}
