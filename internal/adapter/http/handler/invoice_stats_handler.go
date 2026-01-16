package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/converter"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/dto"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/handle"
	"github.com/lubosgarancovsky/eden-inri/internal/adapter/http/validator"
	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	go_kit "github.com/lubosgarancovsky/go-kit"
)

type InvoiceStatsHandler struct {
	getStatsUC ports.GetInvoiceStatsUseCase
	parser     *go_kit.Parser
}

func NewInvoiceStatsHandler(getStatsUC ports.GetInvoiceStatsUseCase, parser *go_kit.Parser) *InvoiceStatsHandler {
	return &InvoiceStatsHandler{getStatsUC, parser}
}

type invoiceStatsAttributes struct {
	ClientID      string `rsql:"filter"`
	PaidAt        string `rsql:"filter"`
	DueAt         string `rsql:"filter"`
	DeliveredAt   string `rsql:"filter"`
	IssuedAt      string `rsql:"filter"`
	Total         string `rsql:"filter"`
	BillableHours string `rsql:"filter"`
}

func (h *InvoiceStatsHandler) GetStats(c *gin.Context) {
	lq := handle.ListingQuery(c, h.parser, &invoiceStatsAttributes{})

	req := &dto.InvoiceStatsReq{}
	if err := validator.BindAndValidate(c, req); err != nil {
		handle.Error(c, err)
		return
	}

	query, err := converter.ToListQuery(req, lq)
	if err != nil {
		handle.Error(c, err)
		return
	}

	stats, err := h.getStatsUC.Execute(c.Request.Context(), query)
	if err != nil {
		handle.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, converter.ToInvoiceStatsResponse(stats))
}
