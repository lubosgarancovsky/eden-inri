package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	go_kit "github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

var _ ports.QueryProjectDocumentPort = (*ProjectDocumentRepository)(nil)
var _ ports.PersistProjectDocumentPort = (*ProjectDocumentRepository)(nil)

type ProjectDocumentRepository struct{ db *gorm.DB }

func NewProjectDocumentRepository(db *gorm.DB) *ProjectDocumentRepository {
	return &ProjectDocumentRepository{db: db}
}

func (r *ProjectDocumentRepository) List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ProjectDocument, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.ProjectDocument{}).Where("project_id = ?", projectID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}
	itemsM, total, err := List[model.ProjectDocument](query, lq)
	if err != nil {
		return nil, 0, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	items := make([]entity.ProjectDocument, len(itemsM))
	for i := range itemsM {
		items[i] = *itemsM[i].ToDomain()
	}
	return &items, total, nil
}

func (r *ProjectDocumentRepository) FindByID(ctx context.Context, projectID, documentID uuid.UUID) (*entity.ProjectDocument, error) {
	db := GetDB(ctx, r.db)
	var m model.ProjectDocument
	if err := db.Model(&model.ProjectDocument{}).Where("id = ? AND project_id = ?", documentID, projectID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("document %s not found", documentID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToDomain(), nil
}

func (r *ProjectDocumentRepository) Create(ctx context.Context, doc *entity.ProjectDocument) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(model.ProjectDocumentFromDomain(doc)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *ProjectDocumentRepository) Update(ctx context.Context, doc *entity.ProjectDocument) error {
	db := GetDB(ctx, r.db)
	res := db.Model(&model.ProjectDocument{}).Where("id = ? AND project_id = ?", doc.ID, doc.ProjectID).Updates(model.ProjectDocumentFromDomain(doc))
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("document %s not found", doc.ID))
	}
	return nil
}

func (r *ProjectDocumentRepository) Delete(ctx context.Context, projectID, documentID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ? AND project_id = ?", documentID, projectID).Delete(&model.ProjectDocument{})
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("document %s not found", documentID))
	}
	return nil
}
