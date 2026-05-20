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

type TaxStatsHandler struct {
	getStatsUC ports.GetTaxStatsUseCase
	parser     *go_kit.Parser
}

func NewTaxStatsHandler(
	getStatsUC ports.GetTaxStatsUseCase,
	parser *go_kit.Parser,
) *TaxStatsHandler {
	return &TaxStatsHandler{getStatsUC, parser}
}

type taxStatsAttributes struct {
	Category  string `rsql:"filter"`
	Amount    string `rsql:"filter"`
	Currency  string `rsql:"filter"`
	Period    string `rsql:"filter"`
	PaidAt    string `rsql:"filter"`
	CreatedAt string `rsql:"filter"`
}

func (h *TaxStatsHandler) GetStats(c *gin.Context) {
	lq := handle.ListingQuery(c, h.parser, &taxStatsAttributes{})

	req := &dto.TaxStatsReq{}
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

	c.JSON(http.StatusOK, converter.ToTaxStatsResponse(stats))
}
