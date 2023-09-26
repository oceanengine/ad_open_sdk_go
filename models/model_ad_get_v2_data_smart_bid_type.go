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

// AdGetV2DataSmartBidType
type AdGetV2DataSmartBidType string

// List of ad_get_v2_data_smart_bid_type
const (
	SMART_BID_CUSTOM_AdGetV2DataSmartBidType       AdGetV2DataSmartBidType = "SMART_BID_CUSTOM"
	SMART_BID_NO_BID_AdGetV2DataSmartBidType       AdGetV2DataSmartBidType = "SMART_BID_NO_BID"
	SMART_BID_CONSERVATIVE_AdGetV2DataSmartBidType AdGetV2DataSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_RADICAL_AdGetV2DataSmartBidType      AdGetV2DataSmartBidType = "SMART_BID_RADICAL"
)

// All allowed values of AdGetV2DataSmartBidType enum
var AllowedAdGetV2DataSmartBidTypeEnumValues = []AdGetV2DataSmartBidType{
	"SMART_BID_CUSTOM",
	"SMART_BID_NO_BID",
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_RADICAL",
}

// NewAdGetV2DataSmartBidTypeFromValue returns a pointer to a valid AdGetV2DataSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataSmartBidTypeFromValue(v string) (*AdGetV2DataSmartBidType, error) {
	ev := AdGetV2DataSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataSmartBidType: valid values are %v", v, AllowedAdGetV2DataSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataSmartBidType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_smart_bid_type value
func (v AdGetV2DataSmartBidType) Ptr() *AdGetV2DataSmartBidType {
	return &v
}