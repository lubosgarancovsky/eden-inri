package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/mapper"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/repository/postgres/model"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
	"github.com/lubosgarancovsky/go-kit"
	"gorm.io/gorm"
)

type ProjectDocumentRepository struct {
	db *gorm.DB
}

func NewProjectDocumentRepository(db *gorm.DB) *ProjectDocumentRepository {
	return &ProjectDocumentRepository{db: db}
}

func (r *ProjectDocumentRepository) Create(ctx context.Context, doc *entity.ProjectDocument) error {
	db := GetDB(ctx, r.db)
	if err := db.Create(mapper.ProjectDocumentFromDomain(doc)).Error; err != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return nil
}

func (r *ProjectDocumentRepository) Update(ctx context.Context, doc *entity.ProjectDocument) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ?", doc.ID).Where("project_id = ?", doc.ProjectID).Updates(mapper.ProjectDocumentFromDomain(doc))
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("document with id %s not found", doc.ID))
	}
	return nil
}

func (r *ProjectDocumentRepository) Delete(ctx context.Context, projectID, documentID uuid.UUID) error {
	db := GetDB(ctx, r.db)
	res := db.Where("id = ?", documentID).Where("project_id = ?", projectID).Delete(&model.ProjectDocument{})
	if res.Error != nil {
		return go_kit.Wrap(go_kit.ErrInternalServer, res.Error)
	}
	if res.RowsAffected == 0 {
		return go_kit.ErrNotFound.WithMessage(fmt.Sprintf("document with id %s not found", documentID))
	}
	return nil
}

func (r *ProjectDocumentRepository) FindByID(ctx context.Context, projectID, documentID uuid.UUID) (*entity.ProjectDocument, error) {
	db := GetDB(ctx, r.db)
	var m model.ProjectDocument
	if err := db.Where("id = ?", documentID).Where("project_id = ?", projectID).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, go_kit.ErrNotFound.WithMessage(fmt.Sprintf("document with id %s not found", documentID))
		}
		return nil, go_kit.Wrap(go_kit.ErrInternalServer, err)
	}
	return m.ToDomain(), nil
}

func (r *ProjectDocumentRepository) List(ctx context.Context, projectID uuid.UUID, lq *go_kit.ListingQuery) (*[]entity.ProjectDocument, int64, error) {
	db := GetDB(ctx, r.db)
	query := db.Model(&model.ProjectDocument{}).Where("project_id = ?", projectID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := ListToDomain[model.ProjectDocument, entity.ProjectDocument](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}
