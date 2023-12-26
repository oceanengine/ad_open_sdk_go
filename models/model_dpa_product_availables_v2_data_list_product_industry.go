/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
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
	ECOMMERCE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE"
	VIDEO_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "VIDEO"
	FINANCE_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "FINANCE"
	WEALTH_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "WEALTH"
	AUTO_NEW_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_NEW"
	NOVEL_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "NOVEL"
	LIVE_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "LIVE"
	ECOMMERCE_V2_DpaProductAvailablesV2DataListProductIndustry     DpaProductAvailablesV2DataListProductIndustry = "ECOMMERCE_V2"
	GAME_DpaProductAvailablesV2DataListProductIndustry             DpaProductAvailablesV2DataListProductIndustry = "GAME"
	TOUR_ROUTE_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_ROUTE"
	TRANSPORT_TICKET_DpaProductAvailablesV2DataListProductIndustry DpaProductAvailablesV2DataListProductIndustry = "TRANSPORT_TICKET"
	MEDICINE_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "MEDICINE"
	GENERAL_DpaProductAvailablesV2DataListProductIndustry          DpaProductAvailablesV2DataListProductIndustry = "GENERAL"
	ESTATE_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "ESTATE"
	OTHER_DpaProductAvailablesV2DataListProductIndustry            DpaProductAvailablesV2DataListProductIndustry = "OTHER"
	MEDICAL_SERVICE_DpaProductAvailablesV2DataListProductIndustry  DpaProductAvailablesV2DataListProductIndustry = "MEDICAL_SERVICE"
	NEW_HOUSE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "NEW_HOUSE"
	COMMUNICATION_DpaProductAvailablesV2DataListProductIndustry    DpaProductAvailablesV2DataListProductIndustry = "COMMUNICATION"
	AUTO_OLD_DpaProductAvailablesV2DataListProductIndustry         DpaProductAvailablesV2DataListProductIndustry = "AUTO_OLD"
	TOUR_TICKET_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "TOUR_TICKET"
	FURNITURE_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "FURNITURE"
	TOUR_HOTEL_DpaProductAvailablesV2DataListProductIndustry       DpaProductAvailablesV2DataListProductIndustry = "TOUR_HOTEL"
	RECRUITMENT_DpaProductAvailablesV2DataListProductIndustry      DpaProductAvailablesV2DataListProductIndustry = "RECRUITMENT"
	MERCHANTS_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "MERCHANTS"
	CREDIT_DpaProductAvailablesV2DataListProductIndustry           DpaProductAvailablesV2DataListProductIndustry = "CREDIT"
	EDUCATION_DpaProductAvailablesV2DataListProductIndustry        DpaProductAvailablesV2DataListProductIndustry = "EDUCATION"
)

// All allowed values of DpaProductAvailablesV2DataListProductIndustry enum
var AllowedDpaProductAvailablesV2DataListProductIndustryEnumValues = []DpaProductAvailablesV2DataListProductIndustry{
	"ECOMMERCE",
	"VIDEO",
	"FINANCE",
	"WEALTH",
	"AUTO_NEW",
	"NOVEL",
	"LIVE",
	"ECOMMERCE_V2",
	"GAME",
	"TOUR_ROUTE",
	"TRANSPORT_TICKET",
	"MEDICINE",
	"GENERAL",
	"ESTATE",
	"OTHER",
	"MEDICAL_SERVICE",
	"NEW_HOUSE",
	"COMMUNICATION",
	"AUTO_OLD",
	"TOUR_TICKET",
	"FURNITURE",
	"TOUR_HOTEL",
	"RECRUITMENT",
	"MERCHANTS",
	"CREDIT",
	"EDUCATION",
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
