package invoice

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/lubosgarancovsky/eden-inri/internal/app/ports"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/command"
	"github.com/lubosgarancovsky/eden-inri/internal/domain/entity"
)

const promptTemplate = `
You are reading an invoice. Return the following fields exactly as printed on the invoice.

Rules:
- Copy text and numbers exactly as shown, but convert the DD.MM.YYYY dates to ISO strings
- Invoice has dates in DD.MM.YYYY format, in JSON they must be in iso-string format
- If a field is not visible or present use empty value or null for numbers
- Do not invent any values
- Respond with ONLY VALID JSON OBJECT, no extra text
- Do not invent any values
- Correctly distinguish between supplier and client, they are usually in two columns next to each other

Important:
- Supplier and Client information may appear in two side-by-side columns.
- Text on the same line does not necessarily belong to the same entity.
- The left column belongs to Supplier.
- The right column belongs to Client.
- Do not merge values from the left and right columns.
- If a line contains both supplier and client data, split it according to the column structure.
- Supplier fields must contain only supplier information.
- Client fields must contain only client information.

{
	"invoiceNumber": "string"
	"orderNumber": "string",
	"supplier": {
		"name": "string",
		"address": "string",
		"ICO": "string",
		"DIC": "string",
		"IC_DPH": "string",
		"phone": "string",
		"email": "string",
	},
	"client": {
		"name": "string",
		"address": "string",
		"ICO": "string",
		"DIC": "string",
		"IC_DPH": "string",
		"phone": "string",
		"email": "string",
	},
	"items": [{
		"name": "string",
		"count": "number",
		"unit": "h | pc",
		"unitPrice": "number",
		"sum": "number"
	}],
	"deliveredAt": "iso string",
	"issuedAt": "iso string",
	"dueAt": "iso string",
	"variableSymbol": "string",
	"iban": "string",
	"swift": "string",
	"currency": "string",
	"total": "number"
}

INVOICE TEXT:
%s
`

type AnalyzeInvoicePdfService struct {
	ollama ports.OllamaPort
}

func NewAnalyzeInvoicePdfService(ollama ports.OllamaPort) *AnalyzeInvoicePdfService {
	return &AnalyzeInvoicePdfService{ollama}
}

func (s *AnalyzeInvoicePdfService) Execute(ctx context.Context, cmd *command.FilesCommand) (*entity.InvoiceAnalysis, error) {
	var allText strings.Builder

	for _, source := range cmd.Sources {

		path, cleanup, err := readerToTempFile(source.Reader)
		if err != nil {
			return nil, err
		}

		defer cleanup()

		text, err := extractPDFText(ctx, path)
		if err != nil {
			return nil, err
		}

		allText.WriteString("\n--- DOCUMENT ---\n")
		allText.WriteString(text)
	}

	prompt := fmt.Sprintf(promptTemplate, allText.String(), allText.String())

	response, err := s.ollama.Analyze(prompt, []string{})
	if err != nil {
		return nil, err
	}

	var invoice entity.InvoiceAnalysis

	if err := json.Unmarshal([]byte(response), &invoice); err != nil {
		return nil, err
	}

	return &invoice, nil
}

func extractPDFText(ctx context.Context, path string) (string, error) {
	cmd := exec.CommandContext(
		ctx,
		"pdftotext",
		"-layout", // preserves spacing
		path,
		"-", // output to stdout
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("pdftotext failed: %w, output=%s", err, string(out))
	}

	return string(out), nil
}

func readerToTempFile(r io.Reader) (string, func() error, error) {
	tmp, err := os.CreateTemp("", "invoice-*.pdf")
	if err != nil {
		return "", nil, err
	}

	_, err = io.Copy(tmp, r)
	if err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return "", nil, err
	}

	tmp.Close()

	cleanup := func() error {
		return os.Remove(tmp.Name())
	}

	return tmp.Name(), cleanup, nil
}
