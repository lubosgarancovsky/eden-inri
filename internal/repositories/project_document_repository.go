package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/models"
	"github.com/lubosgarancovsky/eden-inri/pkg/helpers"
	"github.com/lubosgarancovsky/go-kit/api_err"
	"github.com/lubosgarancovsky/go-kit/list"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProjectDocumentRepository struct {
	db *gorm.DB
}

func NewProjectDocumentRepository(db *gorm.DB) *ProjectDocumentRepository {
	return &ProjectDocumentRepository{db: db}
}

func (r *ProjectDocumentRepository) FindAll(ctx context.Context, projectID uuid.UUID, lq *list.ListingQuery) (*[]models.ProjectDocument, int64, error) {
	query := r.db.
		WithContext(ctx).
		Model(models.ProjectDocument{}).
		Where("project_id = ?", projectID)

	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[models.ProjectDocument](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return &items, total, nil
}

func (r *ProjectDocumentRepository) FindByID(ctx context.Context, projectID uuid.UUID, documentID uuid.UUID) (*models.ProjectDocument, error) {
	var result models.ProjectDocument
	if err := r.db.
		WithContext(ctx).
		Model(models.ProjectDocument{}).
		Where("project_id = ? AND id = ?", projectID, documentID).
		First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ProjectDocumentRepository) Insert(ctx context.Context, doc *models.ProjectDocument) (*models.ProjectDocument, error) {
	err := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Create(doc).
		Error

	return doc, err
}

func (r *ProjectDocumentRepository) Update(ctx context.Context, doc *models.ProjectDocument) (*models.ProjectDocument, error) {
	result := r.db.
		WithContext(ctx).
		Clauses(clause.Returning{}).
		Where("project_id = ? AND id = ?", doc.ProjectID, doc.ID).
		Updates(&doc)

	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Project with id %s does not exist", doc.ID))
	}
	return doc, nil
}

func (r *ProjectDocumentRepository) Delete(ctx context.Context, projectID, documentID uuid.UUID) error {
	result := r.db.
		WithContext(ctx).
		Where("project_id = ? AND id = ?", projectID, documentID).
		Delete(models.ProjectDocument{})

	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Project document with id %s does not exist", documentID))
	}
	return nil
}
