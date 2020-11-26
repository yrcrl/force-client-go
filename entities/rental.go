package force

import "time"

const RentalObjectName = "Rental__c"

type Rental struct {
	Contract             *Contract  `json:"Contract__c,omitempty"`
	HandoverDate         *time.Time `json:"Handover_date__c,omitempty"`
	HandoverNotes        string     `json:"Handover_notes__c,omitempty"`
	Now                  *time.Time `json:"Now__c"`
	PlanedHandoverDate   *time.Time `json:"Planed_Handover_date__c,omitempty"`
	PlanedReturnDate     *time.Time `json:"Planed_Return_date__c,omitempty"`
	PossibleHandoverDate *time.Time `json:"Possible_Handover_date__c,omitempty"`
	RentalName           string     `json:"Name"`
	ReturnDate           *time.Time `json:"Return_date__c,omitempty"`
	ReturnNotes          string     `json:"Return_notes__c,omitempty"`
	Status               string     `json:"Status__c,omiempty"`
	Type                 string     `json:"Type__c,omitempty"`
	VehicleId            string     `json:"Vehicle__c,omitempty"`
	Vehicle              *Vehicle   `json:"Vehicle__r,omitempty"`
}
