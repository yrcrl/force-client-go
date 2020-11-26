package entities

import "time"

const VehicleObjectName = "Vehicle__c"

type Vehicle struct {
	Age                 int64      `json:"Age__c,omitempty"`
	AllSeasonTires      string     `json:"all_season_tires__c,omitempty"`
	Color               string     `json:"Color__c,omitempty"`
	DeliveryDate        *time.Time `json:"Delivery_Date__c,omitempty"`
	Documents           string     `json:"documents__c,omitempty"`
	Email               string     `json:"email__c,omitempty"`
	FleetioID           string     `json:"Fleetio_ID__c,omitempty"`
	HasRental           int        `json:"Has_Rental__c,omitempty"`
	HasReservation      int        `json:"Has_Reservation__c,omitempty"`
	InitialRegistration *time.Time `json:"Initial_registration__c,omitempty"`
	LicensePlate        string     `json:"license_plate__c,omitempty"`
	MakeId              string     `json:"make__c,omitempty"`
	Make                *Make      `json:"make__r,omitempty"`
	MeterValue          int        `json:"Meter_Value__c,omitempty"`
	ModelId             string     `json:"Model__c,omitempty"`
	Model               *Model     `json:"Model__r,omitempty"`
	NewMotoID           string     `json:"NUUID__c,omitempty"`
	RentalStatus        string     `json:"Rental_Status__c,omitempty"`
	Status              string     `json:"state__c,omitempty"`
	TrimId              string     `json:"Trim__c,omitempty"`
	Trim                *Trim      `json:"Trim__r,omitempty"`
	VehicleName         string     `json:"Name,omitempty"`
	VehicleStatus       string     `json:"Vehicle_Status__c,omitempty"`
	VignetteValidUntil  *time.Time `json:"vignette_valid_until__c,omitempty"`
	VIN                 string     `json:"VIN__c,omitempty"`
	Year                string     `json:"year__c,omitempty"`
}

const MakeObjectName = "make__c"

type MakeItemSet struct {
	BaseRecordSet
	Records []Make `json:"records"`
}

type Make struct {
	ID   string `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

const ModelObjectName = "Model__c"

type ModelItemSet struct {
	BaseRecordSet
	Records []Model `json:"records"`
}

type Model struct {
	ID     string `json:"Id,omitempty"`
	Make   *Make  `json:"Make__r,omitempty"`
	MakeId string `json:"Make__c,omitempty"`
	Name   string `json:"Name,omitempty"`
}

type TrimItemSet struct {
	BaseRecordSet
	Records []Trim `json:"records"`
}

const TrimObjectName = "Trim__c"

type Trim struct {
	ID      string `json:"Id,omitempty"`
	Make    *Make  `json:"make__r,omitempty"`
	MakeId  string `json:"make__c,omitempty"`
	Model   *Model `json:"Model__r,omitempty"`
	ModelId string `json:"Model__c,omitempty"`
	Name    string `json:"Name,omitempty"`
}
