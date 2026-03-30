package entity

import (
	"time"

	"github.com/google/uuid"
)

type TaxCategory string

const (
	HealthInsurance         TaxCategory = "health_insurance"
	SocialInsurance         TaxCategory = "social_insurance"
	IncomeTax               TaxCategory = "income_tax"
	CorporateIncomeTax      TaxCategory = "corporate_income_tax"
	CapitalIncomeTax        TaxCategory = "capital_income_tax"
	ExciseTax               TaxCategory = "excise_tax"
	PropertyTax             TaxCategory = "property_tax"
	DogTax                  TaxCategory = "dog_tax"
	AccommodationTax        TaxCategory = "accommodation_tax"
	MotorVehicleTax         TaxCategory = "motor_vehicle_tax"
	FinancialTransactionTax TaxCategory = "financial_transaction_tax"
	TaxAdvance              TaxCategory = "tax_advance"
	Penalty                 TaxCategory = "penalty"
	VAT                     TaxCategory = "vat"
	TaxDue                  TaxCategory = "tax_due"
)

type Tax struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Category    TaxCategory
	Amount      float64
	Currency    string
	Description *string
	Reference   *string
	Period      *time.Time
	PaidAt      *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
