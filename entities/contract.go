package entities

import (
	"fmt"
	"strings"
	"time"
)

type JSONDate struct {
	time.Time
}

func (t *JSONDate) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), `"`)
	v, err := time.Parse("2006-01-02", s)
	*t = JSONDate{v}
	return
}

func (t JSONDate) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *JSONDate) String() string {
	return fmt.Sprintf("%q", t.Format("2006-01-02"))
}

const ContractObjectName = "Contract"

type ContractSet struct {
	BaseRecordSet
	Records []Contract `json:"records"`
}

type Contract struct {
	ID                  string   `json:"Id,omitempty"`
	ContractNumber      string   `json:"ContractNumber,omitempty"`
	Number              string   `json:"Name,omitempty"`
	AccountID           string   `json:"AccountId,omitempty"`
	Account             *Account `json:"Account,omitempty"`
	Status              string   `json:"Status,omitempty"`
	MonthlyFee          float64  `json:"Monthly_Fee__c,omitempty"`
	CustomLineItemDescr string   `json:"custom_line_item_description__c,omitempty"`

	Abo                       string     `json:"Abo__c,omitempty"`
	AboClass                  string     `json:"Abo_class__c,omitempty"`
	ActivatedById             string     `json:",omitempty"`
	ActivatedBy               string     `json:"ActivatedById,omitempty"`
	ActivatedDate             *time.Time `json:",omitempty"`
	ActiveRentalId            string     `json:"Active_Rental__c,omitempty,omitempty"`
	ActiveRental              *Rental    `json:"Active_Rental__r,omitempty,omitempty"`
	ActiveVehicle             string     `json:"Active_Vehicle__c,omitempty"`
	ActiveVehicleLicensePlate string     `json:"Active_Vehicle_License_Plate__c,omitempty"`
	ActiveVehicleName         string     `json:"Active_Vehicle_Name__c,omitempty"`
	AgeClass                  string     `json:"Age_class__c,omitempty"`
	BillingAddress            string     `json:"BillingAddress,omitempty"`
	ContractFee               float64    `json:"Contract_Fee__c,omitempty"`
	BillingActive             bool       `json:"Billing_Active__c,omitempty"`
	CustomerCancellationDate  *JSONDate  `json:"Customer_Cancellation_Date__c,omitempty"`
	ContractEndDate           *JSONDate  `json:"Contract_End_Date__c,omitempty"`

	// CompanySigned By    CompanySignedId    Lookup(User)        False
	// Company Signed Date    CompanySignedDate    Date        False
	// Contact    Contact__c    Lookup(Contact)        True
	// Contract Currency    CurrencyIsoCode    Picklist        False
	// Contract End Date    EndDate    Date        True
	// Contract Fee Reported    Contract_Fee_Reported__c    Checkbox        False
	// Contract Owner    OwnerId    Lookup(User)        True
	// Contract Start Date    StartDate    Date        False
	// Contract Term (months)    ContractTerm    Number(4, 0)        False
	// Contract Terminated Date    Contract_Terminated_Date__c    Date        False
	// Created By    CreatedById    Lookup(User)        False
	// Created Date    Created_Date__c    Formula (Date)        False
	// Customer Signed By    CustomerSignedId    Lookup(Contact)        True
	// Customer Signed Date    CustomerSignedDate    Date        False
	// Customer Signed Title    CustomerSignedTitle    Text(40)        False
	// Deposit Received At    Deposit_Received_At__c    Roll-Up Summary (MAX Deposit)        False
	// Description    Description    Long Text Area(32000)        False
	// Initial fee    Initial_fee__c    Currency(16, 2)        False
	// Is Carl Contract    Is_Carl_Contract__c    Checkbox        False
	// Last Modified By    LastModifiedById    Lookup(User)        False
	// Legal Contract Fee    legal_contract_fee__c    Number(16, 2)        False
	// Opportunity    Opportunity__c    Lookup(Opportunity)        True
	// Owner Expiration Notice    OwnerExpirationNotice    Picklist        False
	// Shipping Address    ShippingAddress    Address        False
	// Special Terms    SpecialTerms    Text Area(4000)        False
	// Termination Date    Termination_Date__c    Date        False
	// Termineted By Customer    Termineted_By_Customer__c    Checkbox        False
	// Total Monthly Fee    TotalMonthlyFee__c    Formula (Currency)
}

const ContractLineItemName = "ContractLineItem__c"

type ContractLineItemSet struct {
	BaseRecordSet
	Records []ContractLineItem `json:"records"`
}

type ContractLineItem struct {
	ID                        string    `json:"Id,omitempty"`
	Name                      string    `json:"Name,omitempty"`
	ContractID                string    `json:"Contract__c,omitempty"`
	Contract                  *Contract `json:"Contract__r,omitempty"`
	IsActive                  bool      `json:"IsActive__c,omitemptry"`
	MonthlyFee                float64   `json:"Monthly_Fee__c,omitempty"`
	ProductID                 string    `json:"Product__c,omitempty"`
	Product                   *Product  `json:"Product__r,omitempty"`
	MainOpportunityLineItemID string    `json:"Main_Opportunity_Line_Item_ID__c,omitempty"`
}
