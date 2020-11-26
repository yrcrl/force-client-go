package entities

import (
	"encoding/json"
	"errors"
)

const InvoiceItemObjectName = "Invoice_Item__c"

type InvoiceItemType string

const (
	InvoiceItemTypeNone    InvoiceItemType = "None"
	InvoiceItemTypeAbo     InvoiceItemType = "Abo"
	InvoiceItemTypeService InvoiceItemType = "Service"
)

func (i *InvoiceItemType) UnmarshalJSON(b []byte) error {
	var s string
	json.Unmarshal(b, &s)
	itemType := InvoiceItemType(s)
	switch itemType {
	case InvoiceItemTypeNone, InvoiceItemTypeAbo, InvoiceItemTypeService:
		*i = itemType
		return nil
	}
	return errors.New("Invalid Invoice Item Type")
}

type InvoiceItemSet struct {
	BaseRecordSet
	Records []InvoiceItem `json:"records"`
}

type InvoiceItem struct {
	ID                  string          `json:"Id,omitempty"`
	AccountId           string          `json:"Account__c,omitempty"`
	Account             *Account        `json:"Account__r,omitempty"`
	ContractId          string          `json:"Contract__c,omitempty"`
	Quantity            float64         `json:"Quantity__c,omitempty"`
	UnitAmount          float64         `json:"Unit_Amount__c,omitempty"`
	Amount              float64         `json:"Amount__c,omitempty"`
	Description         string          `json:"Description__c,omitempty"`
	Name                string          `json:",omitempty"`
	InvoiceId           string          `json:"Invoice__c,omitempty"`
	StripeInvoiceItemID string          `json:"Stripe_Invoice_Item_ID__c,omitempty"`
	BillingPeriodEnd    *JSONDate       `json:"Billing_Period_End__c,omitempty"`
	BillingPeriodStart  *JSONDate       `json:"Billing_Period_Start__c,omitempty"`
	VehicleId           string          `json:"Vehicle__c,omitempty"`
	InvoiceItemType     InvoiceItemType `json:"Invoice_Item_Type__c,omitempty"`
}
