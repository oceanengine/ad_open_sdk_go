/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2ConvertedTimeDuration
type AudiencePackageCreateV2ConvertedTimeDuration string

// List of audience_package_create_v2_converted_time_duration
const (
	SEVEN_DAY_AudiencePackageCreateV2ConvertedTimeDuration    AudiencePackageCreateV2ConvertedTimeDuration = "SEVEN_DAY"
	SIX_MONTH_AudiencePackageCreateV2ConvertedTimeDuration    AudiencePackageCreateV2ConvertedTimeDuration = "SIX_MONTH"
	THREE_MONTH_AudiencePackageCreateV2ConvertedTimeDuration  AudiencePackageCreateV2ConvertedTimeDuration = "THREE_MONTH"
	ONE_MONTH_AudiencePackageCreateV2ConvertedTimeDuration    AudiencePackageCreateV2ConvertedTimeDuration = "ONE_MONTH"
	TODAY_AudiencePackageCreateV2ConvertedTimeDuration        AudiencePackageCreateV2ConvertedTimeDuration = "TODAY"
	NONE_AudiencePackageCreateV2ConvertedTimeDuration         AudiencePackageCreateV2ConvertedTimeDuration = "NONE"
	TWELVE_MONTH_AudiencePackageCreateV2ConvertedTimeDuration AudiencePackageCreateV2ConvertedTimeDuration = "TWELVE_MONTH"
)

// All allowed values of AudiencePackageCreateV2ConvertedTimeDuration enum
var AllowedAudiencePackageCreateV2ConvertedTimeDurationEnumValues = []AudiencePackageCreateV2ConvertedTimeDuration{
	"SEVEN_DAY",
	"SIX_MONTH",
	"THREE_MONTH",
	"ONE_MONTH",
	"TODAY",
	"NONE",
	"TWELVE_MONTH",
}

// NewAudiencePackageCreateV2ConvertedTimeDurationFromValue returns a pointer to a valid AudiencePackageCreateV2ConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2ConvertedTimeDurationFromValue(v string) (*AudiencePackageCreateV2ConvertedTimeDuration, error) {
	ev := AudiencePackageCreateV2ConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2ConvertedTimeDuration: valid values are %v", v, AllowedAudiencePackageCreateV2ConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2ConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2ConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_converted_time_duration value
func (v AudiencePackageCreateV2ConvertedTimeDuration) Ptr() *AudiencePackageCreateV2ConvertedTimeDuration {
	return &v
}
