package entities

import (
	"fmt"
	"strings"
)

const AccountObjectName = "Account"

type AccountSet struct {
	BaseRecordSet
	Records []Account `json:"records"`
}

type Account struct {
	Attr            Attr   `json:"attributes,omitempty"`
	ID              string `json:"Id,omitempty"`
	AccountCurrency string `json:"CurrencyIsoCode,omitempty"`
	AccountName     string `json:"Name,omitempty"`
	AccountNumber   string `json:"AccountNumber,omitempty"`
	//AccountOwner    User   `json:"OwnerId,omitempty"`
	AccountRecordType string       `json:"RecordTypeId,omitempty"`
	AccountSource     string       `json:"AccountSource,omitempty"`
	AnnualRevenue     Currency     `json:"AnnualRevenue,omitempty"`
	BilendoCustomerID string       `json:"Bilendo_Customer_ID__c,omitempty"`
	BillingAddress    *Address     `json:"BillingAddress,omitempty"`
	BillingEmail      string       `json:"Billing_E_Mail__c,omitempty"`
	Category          string       `json:"Category__c,omitempty"`
	CompanyName       string       `json:"Company_name__c,omitempty"`
	DataComKey        string       `json:"Jigsaw,omitempty"`
	Description       LongTextArea `json:"Description,omitempty"`
	Employees         float64      `json:"NumberOfEmployees,omitempty"`
	Fax               Fax          `json:"Fax,omitempty"`
	FleetAccount      bool         `json:"Fleet_Account__c,omitempty"`
	Industry          string       `json:"Industry,omitempty"`
	Ownership         string       `json:"Ownership,omitempty"`
	ParentAccount     string       `json:"ParentId,omitempty"`
	Phone             string       `json:"Phone,omitempty"`
	Rating            string       `json:"Rating,omitempty"`
	ShippingAddress   string       `json:"ShippingAddress,omitempty"`
	SICCode           string       `json:"Sic,omitempty"`
	SICDescription    string       `json:"SicDesc,omitempty"`
	StripeCustomerID  string       `json:"Stripe_Customer_ID__c,omitempty"`
	TaxID             string       `json:"Tax_number__c,omitempty"`
	TickerSymbol      string       `json:"TickerSymbol,omitempty"`
	Type              string       `json:"Type,omitempty"`
	Website           string       `json:"Website,omitempty"`
	CreatedBy         *User        `json:"CreatedById,omitempty"`
	LastModifiedBy    *User        `json:"LastModifiedById,omitempty"`

	/* only create an Invoice in Stripe if set */
	UseStripeInvoicing bool `json:"-"`
}

func (a *Account) BillingAddressString() string {
	result := fmt.Sprintf(
		"%s %s %s %s %s",
		a.CompanyName,
		a.BillingAddress.Street,
		a.BillingAddress.Zip,
		a.BillingAddress.City,
		a.BillingAddress.CountryGerman(),
	)
	return strings.TrimSpace(result)
}

type Address struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	//        "geocodeAccuracy": null,
	//        "latitude": null,
	//        "longitude": null,
	Zip string `json:"postalCode,omitempty"`
	//        "state": null,
	Street string `json:"street,omitempty"`
}

func (a *Address) CountryGerman() string {
	switch a.Country {
	case "Austria":
		return "Österreich"
	}
	return a.Country
}
