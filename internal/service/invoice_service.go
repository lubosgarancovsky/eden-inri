package service

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lubosgarancovsky/eden-inri/internal/model"
	"github.com/lubosgarancovsky/eden-inri/internal/repository"
	"github.com/lubosgarancovsky/go-kit/api_err"
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

func (s *InvoiceService) FindAll(userID uuid.UUID, lq *list.ListingQuery) (*list.Page[model.Invoice], error) {
	items, totalCount, err := s.r.FindAll(userID, lq)
	if err != nil {
		return nil, err
	}

	return &list.Page[model.Invoice]{
		Items:      items,
		Page:       lq.Page,
		PageSize:   lq.Limit,
		TotalCount: totalCount,
	}, nil
}

func (s *InvoiceService) FindByID(userID uuid.UUID, id uuid.UUID) (*model.Invoice, error) {
	inv, err := s.r.FindByID(id)
	if err != nil {
		return nil, err
	}
	if inv.UserID != userID {
		return nil, api_err.ErrForbidden
	}
	return inv, nil
}

func (s *InvoiceService) Create(userID uuid.UUID, input *model.InvoiceRequest) (*model.Invoice, error) {
	inv := &model.Invoice{
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
	return s.r.Insert(inv)
}

func (s *InvoiceService) Update(userID uuid.UUID, id uuid.UUID, input *model.InvoiceRequest) (*model.Invoice, error) {
	_, err := s.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	inv := &model.Invoice{
		ID:            id,
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
	return s.r.Update(inv)
}

func (s *InvoiceService) Delete(userID uuid.UUID, id uuid.UUID) (*model.Invoice, error) {
	inv, err := s.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	return inv, s.r.Delete(id)
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
