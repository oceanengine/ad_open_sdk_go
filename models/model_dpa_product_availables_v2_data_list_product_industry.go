/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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
	COMMUNICATION_DpaProductAvailablesV2DataListProductIndustry    DpaProductAvailablesV2DataListProductIndustry = "COMMUNICATION"
	RECRUITMENT_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "RECRUITMENT"
	GENERAL_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "GENERAL"
	OTHER_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "OTHER"
	CREDIT_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "CREDIT"
	VIDEO_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "VIDEO"
	AUTO_NEW_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_NEW"
	MEDICAL_SERVICE_DpaProductAvailablesV2DataListProductIndustry  DpaProductAvailablesV2DataListProductIndustry = "MEDICAL_SERVICE"
	ESTATE_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "ESTATE"
	ECOMMERCE_V2_DpaProductAvailablesV2DataListProductIndustry     DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE_V2"
	NEW_HOUSE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "NEW_HOUSE"
	GAME_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "GAME"
	ECOMMERCE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE"
	FINANCE_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "FINANCE"
	MEDICINE_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "MEDICINE"
	LIVE_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "LIVE"
	TOUR_HOTEL_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_HOTEL"
	NOVEL_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "NOVEL"
	WEALTH_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "WEALTH"
	FURNITURE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "FURNITURE"
	TRANSPORT_TICKET_DpaProductAvailablesV2DataListProductIndustry DpaProductAvailablesV2DataListProductIndustry = "TRANSPORT_TICKET"
	MERCHANTS_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "MERCHANTS"
	AUTO_OLD_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_OLD"
	EDUCATION_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "EDUCATION"
	TOUR_TICKET_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "TOUR_TICKET"
	TOUR_ROUTE_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_ROUTE"
)

// All allowed values of DpaProductAvailablesV2DataListProductIndustry enum
var AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues = []DpaProductAvailablesV2DataListProductIndustry{
	"COMMUNICATION",
	"RECRUITMENT",
	"GENERAL",
	"OTHER",
	"CREDIT",
	"VIDEO",
	"AUTO_NEW",
	"MEDICAL_SERVICE",
	"ESTATE",
	"ECOMMERCE_V2",
	"NEW_HOUSE",
	"GAME",
	"ECOMMERCE",
	"FINANCE",
	"MEDICINE",
	"LIVE",
	"TOUR_HOTEL",
	"NOVEL",
	"WEALTH",
	"FURNITURE",
	"TRANSPORT_TICKET",
	"MERCHANTS",
	"AUTO_OLD",
	"EDUCATION",
	"TOUR_TICKET",
	"TOUR_ROUTE",
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
