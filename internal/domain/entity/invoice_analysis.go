package entity

type InvoiceAnalysis struct {
	InvoiceNumber string         `json:"invoiceNumber"`
	OrderNumber   string         `json:"orderNumber"`
	Supplier      CompanyDetails `json:"supplier"`
	Client        CompanyDetails `json:"client"`
	Items         []InvoiceItem  `json:"items"`

	DeliveredAt string `json:"deliveredAt"`
	IssuedAt    string `json:"issuedAt"`
	DueAt       string `json:"dueAt"`

	VariableSymbol string  `json:"variableSymbol"`
	IBAN           string  `json:"iban"`
	SWIFT          string  `json:"swift"`
	Currency       string  `json:"currency"`
	Total          float64 `json:"total"`
}

type CompanyDetails struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	ICO     string `json:"ICO"`
	DIC     string `json:"DIC"`
	ICDPH   string `json:"IC_DPH"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type InvoiceItem struct {
	Name      string  `json:"name"`
	Count     float64 `json:"count"`
	Unit      string  `json:"unit"`
	UnitPrice float64 `json:"unitPrice"`
	Sum       float64 `json:"sum"`
}
