package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
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

func (r *ProjectDocumentRepository) FindAll(projectID uuid.UUID, lq *list.ListingQuery) ([]model.ProjectDocument, int64, error) {
	query := r.db.Model(&model.ProjectDocument{}).Where("project_id = ?", projectID)
	if lq.Filter != nil {
		query = query.Where(lq.Filter.Query, lq.Filter.Args...)
	}

	items, total, err := helpers.List[model.ProjectDocument](query, lq)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

func (r *ProjectDocumentRepository) FindByID(projectID uuid.UUID, ID uuid.UUID) (*model.ProjectDocument, error) {
	var result model.ProjectDocument
	if err := r.db.Model(&model.ProjectDocument{}).Where("id = ? AND project_id = ?", ID, projectID).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ProjectDocumentRepository) Insert(doc *model.ProjectDocument) (*model.ProjectDocument, error) {
	if err := r.db.Clauses(clause.Returning{}).Create(doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

func (r *ProjectDocumentRepository) Update(doc *model.ProjectDocument) (*model.ProjectDocument, error) {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", doc.ID).Updates(&doc)
	if result.Error != nil {
		return nil, api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, api_err.Wrap(api_err.ErrNotFound, result.Error).WithMessage(fmt.Sprintf("Project with id %s does not exist", doc.ID))
	}
	return doc, nil
}

func (r *ProjectDocumentRepository) Delete(id uuid.UUID) error {
	result := r.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(model.ProjectDocument{})
	if result.Error != nil {
		return api_err.Wrap(api_err.ErrInternalServer, result.Error)
	}
	if result.RowsAffected == 0 {
		return api_err.ErrNotFound.WithMessage(fmt.Sprintf("Project document with id %s does not exist", id))
	}
	return nil
}
