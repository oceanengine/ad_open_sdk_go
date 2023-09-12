/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaProductAvailablesV2DataListProductIndustry
type DpaProductAvailablesV2DataListProductIndustry string

// List of dpa_product_availables_v2_data_list_product_industry
const (
	ESTATE_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "ESTATE"
	CREDIT_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "CREDIT"
	VIDEO_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "VIDEO"
	OTHER_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "OTHER"
	GENERAL_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "GENERAL"
	ECOMMERCE_V2_DpaProductAvailablesV2DataListProductIndustry     DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE_V2"
	TOUR_ROUTE_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_ROUTE"
	COMMUNICATION_DpaProductAvailablesV2DataListProductIndustry    DpaProductAvailablesV2DataListProductIndustry = "COMMUNICATION"
	TOUR_HOTEL_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_HOTEL"
	GAME_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "GAME"
	TOUR_TICKET_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "TOUR_TICKET"
	MERCHANTS_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "MERCHANTS"
	NEW_HOUSE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "NEW_HOUSE"
	MEDICINE_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "MEDICINE"
	TRANSPORT_TICKET_DpaProductAvailablesV2DataListProductIndustry DpaProductAvailablesV2DataListProductIndustry = "TRANSPORT_TICKET"
	MEDICAL_SERVICE_DpaProductAvailablesV2DataListProductIndustry  DpaProductAvailablesV2DataListProductIndustry = "MEDICAL_SERVICE"
	RECRUITMENT_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "RECRUITMENT"
	EDUCATION_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "EDUCATION"
	AUTO_OLD_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_OLD"
	FINANCE_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "FINANCE"
	NOVEL_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "NOVEL"
	AUTO_NEW_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_NEW"
	ECOMMERCE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE"
	WEALTH_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "WEALTH"
	FURNITURE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "FURNITURE"
	LIVE_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "LIVE"
)

// All allowed values of DpaProductAvailablesV2DataListProductIndustry enum
var AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues = []DpaProductAvailablesV2DataListProductIndustry{
	"ESTATE",
	"CREDIT",
	"VIDEO",
	"OTHER",
	"GENERAL",
	"ECOMMERCE_V2",
	"TOUR_ROUTE",
	"COMMUNICATION",
	"TOUR_HOTEL",
	"GAME",
	"TOUR_TICKET",
	"MERCHANTS",
	"NEW_HOUSE",
	"MEDICINE",
	"TRANSPORT_TICKET",
	"MEDICAL_SERVICE",
	"RECRUITMENT",
	"EDUCATION",
	"AUTO_OLD",
	"FINANCE",
	"NOVEL",
	"AUTO_NEW",
	"ECOMMERCE",
	"WEALTH",
	"FURNITURE",
	"LIVE",
}

// NewDpaProductAvailablesV2DataListProductIndustryFromValue returns a pointer to a valid DpaProductAvailablesV2DataListProductIndustry
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaProductAvailablesV2DataListProductIndustryFromValue(v string) (*DpaProductAvailablesV2DataListProductIndustry, error) {
	ev := DpaProductAvailablesV2DataListProductIndustry(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaProductAvailablesV2DataListProductIndustry: valid values are %v", v, AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaProductAvailablesV2DataListProductIndustry) IsValid() bool {
	for _, existing := range AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_product_availables_v2_data_list_product_industry value
func (v DpaProductAvailablesV2DataListProductIndustry) Ptr() *DpaProductAvailablesV2DataListProductIndustry {
	return &v
}
