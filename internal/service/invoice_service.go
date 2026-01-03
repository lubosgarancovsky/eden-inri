package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

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

// RevenueGraph returns monthly revenue points for the last 12 months including current (UTC)
func (s *InvoiceService) RevenueGraph(ctx context.Context, userID uuid.UUID) ([]model.RevenueGraphPoint, error) {
	now := time.Now().UTC()
	// end is first day of next month (exclusive upper bound)
	end := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0)
	// start is 12 months before end (inclusive), covering 12 months including current
	start := end.AddDate(0, -12, 0)

	rows, err := s.r.MonthlyRevenue(ctx, userID, start, end)
	if err != nil {
		return nil, err
	}

	// index by month key
	byMonth := make(map[string]model.RevenueGraphPoint, len(rows))
	for _, p := range rows {
		byMonth[p.Month] = p
	}

	// build full 12-month series from start to end-1month
	points := make([]model.RevenueGraphPoint, 0, 12)
	for i := 0; i < 12; i++ {
		m := start.AddDate(0, i, 0)
		key := fmt.Sprintf("%04d-%02d", m.Year(), int(m.Month()))
		if p, ok := byMonth[key]; ok {
			points = append(points, p)
		} else {
			points = append(points, model.RevenueGraphPoint{
				Month:     key,
				Actual:    0,
				Potential: 0,
			})
		}
	}

	return points, nil
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
