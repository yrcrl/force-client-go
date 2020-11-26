package force

import (
	"fmt"
	"time"
)

const InvoiceObjectName = "Invoice__c"

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	safeString := fmt.Sprintf("\"%s\"", t.Format(time.RFC3339))
	return []byte(safeString), nil
}

type InvoiceSet struct {
	BaseRecordSet
	Records []Invoice `json:"records"`
}

type Invoice struct {
	ID                    string     `json:"Id,omitempty"`
	Name                  string     `json:"Name,omitempty"`
	InvoiceNumber         string     `json:"Invoice_Number__c,omitempty"`
	AccountId             string     `json:"Account__c,omitempty"`
	DueDate               *JSONTime  `json:"Due_Date__c,omitempty"`
	StripeInvoiceID       string     `json:"Stripe_Invoice_ID__c,omitempty"`
	StripeStatus          string     `json:"Stripe_Status__c,omitempty"`
	StripeInvoicePDF      string     `json:"Stripe_Invoice_PDF__c,omitempty"`
	CustomerCompany       string     `json:"Customer_Company__c,omitempty"`
	CustomerName          string     `json:"Customer_Name__c,omitempty"`
	BillingAddressLine1   string     `json:"Billing_Address_Line1__c,omitempty"`
	BillingAddressLine2   string     `json:"Billing_Address_Line2__c,omitempty"`
	BillingAddressCity    string     `json:"Billing_Address_City__c,omitempty"`
	BillingAddressZip     string     `json:"Billing_Address_Zip__c,omitempty"`
	BillingAddressCountry string     `json:"Billing_Address_Country__c,omitempty"`
	CustomerTaxID         string     `json:"Customer_Tax_ID__c,omitempty"`
	Date                  *time.Time `json:"Date__c,omitempty"`
	Subtotal              float64    `json:"Subtotal__c,omitempty"`
	Tax                   float64    `json:"Tax__c,omitempty"`
	TaxPercentage         float64    `json:"Tax_Percentage__c,omitempty"`
	Total                 float64    `json:"Total__c,omitempty"`
	UpdatedAt             *JSONTime  `json:"Updated_At__c,omitempty"`
}
