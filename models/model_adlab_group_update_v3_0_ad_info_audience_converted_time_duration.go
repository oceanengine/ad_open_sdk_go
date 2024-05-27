/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration
type AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration string

// List of adlab_group_update_v3.0_ad_info_audience_converted_time_duration
const (
	NONE_AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration         AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration = "NONE"
	ONE_MONTH_AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration    AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration = "ONE_MONTH"
	SEVEN_DAY_AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration    AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration = "SEVEN_DAY"
	SIX_MONTH_AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration    AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration = "SIX_MONTH"
	THREE_MONTH_AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration  AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration = "THREE_MONTH"
	TODAY_AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration        AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration = "TODAY"
	TWELVE_MONTH_AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration = "TWELVE_MONTH"
)

// All allowed values of AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration enum
var AllowedAdlabGroupUpdateV30AdInfoAudienceConvertedTimeDurationEnumValues = []AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration{
	"NONE",
	"ONE_MONTH",
	"SEVEN_DAY",
	"SIX_MONTH",
	"THREE_MONTH",
	"TODAY",
	"TWELVE_MONTH",
}

// NewAdlabGroupUpdateV30AdInfoAudienceConvertedTimeDurationFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoAudienceConvertedTimeDurationFromValue(v string) (*AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration, error) {
	ev := AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoAudienceConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoAudienceConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_audience_converted_time_duration value
func (v AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration) Ptr() *AdlabGroupUpdateV30AdInfoAudienceConvertedTimeDuration {
	return &v
}
