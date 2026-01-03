package service

import (
	"context"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/list"
)

type InvoiceService struct {
	r                 *repository.InvoiceRepository
	attachmentService *AttachmentService
	ModelName         string
}

func NewInvoiceService(r *repository.InvoiceRepository, attachmentService *AttachmentService) *InvoiceService {
	return &InvoiceService{r: r, attachmentService: attachmentService, ModelName: "invoice"}
}

func (s *InvoiceService) FindAll(ctx context.Context, userID uuid.UUID, lq *list.ListingQuery) (*[]model.Invoice, int64, error) {
	return s.r.FindAll(ctx, userID, lq)
}

func (s *InvoiceService) FindByID(ctx context.Context, userID uuid.UUID, invoiceID uuid.UUID) (*model.Invoice, error) {
	return s.r.FindByID(ctx, userID, invoiceID)
}

func (s *InvoiceService) Create(ctx context.Context, userID uuid.UUID, input *model.InvoiceRequest) (*model.Invoice, error) {
	return s.r.Insert(ctx, s.buildPayload(userID, input))
}

func (s *InvoiceService) Update(ctx context.Context, userID uuid.UUID, invoiceID uuid.UUID, input *model.InvoiceRequest) (*model.Invoice, error) {
	invoice := s.buildPayload(userID, input)
	invoice.ID = invoiceID
	return s.r.Update(ctx, invoice)
}

func (s *InvoiceService) Delete(ctx context.Context, userID uuid.UUID, invoiceID uuid.UUID) error {
	return s.r.Delete(ctx, userID, invoiceID)
}

func (s *InvoiceService) TotalRevenue(ctx context.Context, userID uuid.UUID) (float64, error) {
	return s.r.TotalRevenue(ctx, userID)
}

func (s *InvoiceService) SaveAttachments(c *gin.Context, userID uuid.UUID, invoiceID uuid.UUID, files []multipart.FileHeader) error {
	for _, file := range files {
		_, err := s.attachmentService.SaveAttachment(c, userID, s.ModelName, invoiceID.String(), file)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *InvoiceService) ListAttachments(ctx context.Context, userID uuid.UUID, invoiceID uuid.UUID) ([]*model.Attachment, error) {
	return s.attachmentService.FindByModelID(ctx, userID, invoiceID, s.ModelName)
}

func (s *InvoiceService) buildPayload(userID uuid.UUID, input *model.InvoiceRequest) *model.Invoice {
	return &model.Invoice{
		UserID:        userID,
		ClientID:      input.ClientID,
		Name:          input.Name,
		Note:          input.Note,
		ExternalID:    input.ExternalID,
		Total:         input.Total,
		BillableHours: input.BillableHours,
		IssuedAt:      input.IssuedAt,
		DueAt:         input.DueAt,
		DeliveredAt:   input.DeliveredAt,
		PaidAt:        input.PaidAt,
		IsCanceled:    input.IsCanceled,
		ExternalLink:  input.ExternalLink,
	}
}
