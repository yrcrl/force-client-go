package entities

const InvoiceLineItemObjectName = "Invoice_Line_Item__c"

type InvoiceLineItemSet struct {
	BaseRecordSet
	Records []InvoiceLineItem `json:"records"`
}

type InvoiceLineItem struct {
	ID                      string  `json:"Id,omitempty"`
	Number                  string  `json:"Name,omitempty"`
	Description             string  `json:"Description__c,omitempty"`
	Quantity                float64 `json:"Quantity__c,omitempty"`
	Amount                  float64 `json:"Amount__c,omitempty"`
	InvoiceID               string  `json:"Invoice__c,omitempty"`
	StripeInvoiceItemID     string  `json:"Stripe_Invoice_Item_ID__c,omitempty"`
	StripeInvoiceLineItemID string  `json:"Stripe_Invoice_Line_Item_ID__c,omitempty"`
}
