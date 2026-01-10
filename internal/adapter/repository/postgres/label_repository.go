package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	go_kit "github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

var _ ports.QueryLabelPort = (*LabelRepository)(nil)
var _ ports.PersistLabelPort = (*LabelRepository)(nil)

type LabelRepository struct{ db *gorm.DB }

func NewLabelRepository(db *gorm.DB) *LabelRepository { return &LabelRepository{db: db} }

func (r *LabelRepository) List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]ports.Label, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.Label{}).Where("project_id = ?", projectID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}
	itemsM, total, err := List[model.Label](query, lq)
	if err != nil {
		return nil, 0, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	items := make([]ports.Label, len(itemsM))
	for i := range itemsM {
		items[i] = *itemsM[i].ToPort()
	}
	return &items, total, nil
}

func (r *LabelRepository) FindByID(ctx context.Context, projectID, labelID uuid.UUID) (*ports.Label, error) {
	db := GetDB(ctx, r.db)
	var m model.Label
	if err := db.Model(&model.Label{}).Where("id = ? AND project_id = ?", labelID, projectID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("label %s not found", labelID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToPort(), nil
}

func (r *LabelRepository) Create(ctx context.Context, label *ports.Label) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(model.LabelFromPort(label)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *LabelRepository) Update(ctx context.Context, label *ports.Label) error {
	db := GetDB(ctx, r.db)
	res := db.Model(&model.Label{}).Where("id = ? AND project_id = ?", label.ID, label.ProjectID).Updates(model.LabelFromPort(label))
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("label %s not found", label.ID))
	}
	return nil
}

func (r *LabelRepository) Delete(ctx context.Context, projectID, labelID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ? AND project_id = ?", labelID, projectID).Delete(&model.Label{})
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("label %s not found", labelID))
	}
	return nil
}
