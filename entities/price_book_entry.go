package entities

const PriceBookEntryObjectName = "PricebookEntry"

type PriceBookEntrySet struct {
	BaseRecordSet
	Records []PriceBookEntry `json:"records"`
}
type PriceBookEntry struct {
	ID               string    `json:"Id,omitempty"`
	Active           bool      `json:"IsActive,omitempty"`
	ListPrice        float64   `json:"UnitPrice,omitempty"`
	MonthlyFee       float64   `json:"Monthly_fee__c,omitempty"`
	PriceBookId      string    `json:"Pricebook2Id,omitempty"`
	PriceBook        PriceBook `json:"Pricebook2,omitempty"`
	ProductId        string    `json:"Product2Id,omitempty"`
	Product          Product   `json:"Product2__r,omitempty"`
	ProductCode      string    `json:"ProductCode,omitempty"`
	UseStandardPrice bool      `json:"UseStandardPrice,omitempty"`
}
