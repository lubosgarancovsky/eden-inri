package postgres

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"gorm.io/gorm"
)

type TransactionManager struct {
	db *gorm.DB
}

var _ ports.TransactionManager = (*TransactionManager)(nil)

func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (tm *TransactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	// Start transaction
	tx := tm.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Store transaction in context
	ctx = context.WithValue(ctx, ports.TransactionKey, tx)

	// Execute function
	if err := fn(ctx); err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit().Error
}

// GetDB returns the transaction if it exists in context, otherwise returns the main DB
func GetDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(ports.TransactionKey).(*gorm.DB); ok {
		return tx
	}
	return db
}
