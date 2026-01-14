package ports

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

type QueryUserPort interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
