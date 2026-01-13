package invoice

import (
	"context"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
)

var _ ports.DeleteInvoiceUseCase = (*DeleteInvoiceService)(nil)

type DeleteInvoiceService struct {
	repo ports.PersistInvoicePort
	tm   ports.TransactionManager
}

func NewDeleteInvoiceService(repo ports.PersistInvoicePort) *DeleteInvoiceService {
	return &DeleteInvoiceService{repo: repo}
}

func (s *DeleteInvoiceService) Execute(ctx context.Context, cmd *command.Command) error {
	return s.repo.Delete(ctx, cmd.UserID, cmd.ID)
}
