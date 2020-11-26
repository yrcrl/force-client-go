package entities

import (
	"encoding/json"
	"errors"
)

const ProductObjectName = "Product2"

type ProductFamilyType string

const (
	ProductFamilyAbo       ProductFamilyType = "Abo"
	ProductFamilyUtilities                   = "Utilities"
)

func (f *ProductFamilyType) UnmarshalJSON(b []byte) error {
	var s string
	json.Unmarshal(b, &s)
	familyType := ProductFamilyType(s)
	switch familyType {
	case ProductFamilyAbo, ProductFamilyUtilities:
		*f = familyType
		return nil
	}
	return errors.New("Invalid Product Family")
}

type ProductSet struct {
	BaseRecordSet
	Records []Product `json:"records"`
}

type Product struct {
	ID                          string            `json:"Id,omitempty"`
	AboClass                    string            `json:"Abo_class__c,omitempty"`
	IsActive                    bool              `json:"IsActive,omitempty"`
	AdditionalMileageRate       float64           `json:"additional_mileage_rate__c,omitempty"`
	AgeClass                    string            `json:"Age_class__c,omitempty"`
	AllDescription              string            `json:"allDescription__c,omitempty"`
	AllmonthlyrecurringrevenueC float64           `json:"allMonthlyRecurringRevenue__c,omitempty"`
	AllTimePeriod               float64           `json:"allTimePeriod__c,omitempty"`
	QuoteText                   string            `json:"Angebotstext__c,omitempty"`
	AutoRenewal                 int               `json:"Auto_Verlaengerung__c,omitempty"`
	ClassCode                   string            `json:"Class_Code__c,omitempty"`
	DisplayMileageLimit         float64           `json:"Display_Milage_Limit__c,omitempty"`
	DisplayName                 string            `json:"display_name__c,omitempty"`
	DisplayUrl                  string            `json:"DisplayUrl,omitempty"`
	EinkaufspreisMonatlich      float64           `json:"Einkaufspreis_monatlich__c,omitempty"`
	EinkaufspreisNotiz          string            `json:"Einkaufspreis_Notiz__c,omitempty"`
	ExternalDataSource          string            `json:"ExternalDataSourceId,omitempty"`
	ExternalID                  string            `json:"ExternalId,omitempty"`
	Fee                         bool              `json:"Fee__c,omitempty"`
	IncomeTaxEffectiveness      float64           `json:"income_tax_effectiveness__c,omitempty"`
	InsuranceDeductible         float64           `json:"insurance_deductible__c,omitempty"`
	Kündigungsfrist             int               `json:"Kuendigungsfrist__c,omitempty"`
	MilageLimit                 float64           `json:"Milage_Limit__c,omitempty"`
	MinimumTerm                 int               `json:"Mindestlaufzeit__c,omitempty"`
	NachfolgeProduktId          string            `json:"Nachfolge_Produkt__c,omitempty"`
	ProductCode                 string            `json:"ProductCode,omitempty"`
	Currency                    string            `json:"CurrencyIsoCode,omitempty"`
	ProductDescription          string            `json:"Description,omitempty"`
	ProductFamily               ProductFamilyType `json:"Family,omitempty"`
	ProductName                 string            `json:"Name,omitempty"`
	ProductSKU                  string            `json:"StockKeepingUnit,omitempty"`
	PurchasingPrice             float64           `json:"Einkaufspreis_einmalig__c,omitempty"`
	QuantityUnitOfMeasure       string            `json:"QuantityUnitOfMeasure,omitempty"`
	VATDeductible               float64           `json:"vat_deductible__c,omitempty"`
	YearlyMilageLimit           float64           `json:"Milage_Limit__c,omitempty"`
}
