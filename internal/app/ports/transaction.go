package ports

import "context"

type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

// TxKey is the context key for transaction
type txKey struct{}

// TransactionKey is used to store/retrieve transaction from context
var TransactionKey = txKey{}
