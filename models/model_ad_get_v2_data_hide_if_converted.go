/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataHideIfConverted
type AdGetV2DataHideIfConverted string

// List of ad_get_v2_data_hide_if_converted
const (
	ORGANIZATION_AdGetV2DataHideIfConverted AdGetV2DataHideIfConverted = "ORGANIZATION"
	AD_AdGetV2DataHideIfConverted           AdGetV2DataHideIfConverted = "AD"
	NO_EXCLUDE_AdGetV2DataHideIfConverted   AdGetV2DataHideIfConverted = "NO_EXCLUDE"
	CAMPAIGN_AdGetV2DataHideIfConverted     AdGetV2DataHideIfConverted = "CAMPAIGN"
	CUSTOMER_AdGetV2DataHideIfConverted     AdGetV2DataHideIfConverted = "CUSTOMER"
	APP_AdGetV2DataHideIfConverted          AdGetV2DataHideIfConverted = "APP"
	ADVERTISER_AdGetV2DataHideIfConverted   AdGetV2DataHideIfConverted = "ADVERTISER"
)

// All allowed values of AdGetV2DataHideIfConverted enum
var AllowedAdGetV2DataHideIfConvertedEnumValues = []AdGetV2DataHideIfConverted{
	"ORGANIZATION",
	"AD",
	"NO_EXCLUDE",
	"CAMPAIGN",
	"CUSTOMER",
	"APP",
	"ADVERTISER",
}

// NewAdGetV2DataHideIfConvertedFromValue returns a pointer to a valid AdGetV2DataHideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataHideIfConvertedFromValue(v string) (*AdGetV2DataHideIfConverted, error) {
	ev := AdGetV2DataHideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataHideIfConverted: valid values are %v", v, AllowedAdGetV2DataHideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataHideIfConverted) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataHideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_hide_if_converted value
func (v AdGetV2DataHideIfConverted) Ptr() *AdGetV2DataHideIfConverted {
	return &v
}