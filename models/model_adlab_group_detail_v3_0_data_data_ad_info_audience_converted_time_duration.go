/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration
type AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration string

// List of adlab_group_detail_v3.0_data_data_ad_info_audience_converted_time_duration
const (
	NONE_AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration         AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration = "NONE"
	ONE_MONTH_AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration    AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration = "ONE_MONTH"
	SEVEN_DAY_AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration    AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration = "SEVEN_DAY"
	SIX_MONTH_AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration    AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration = "SIX_MONTH"
	THREE_MONTH_AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration  AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration = "THREE_MONTH"
	TODAY_AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration        AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration = "TODAY"
	TWELVE_MONTH_AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration = "TWELVE_MONTH"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration enum
var AllowedAdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDurationEnumValues = []AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration{
	"NONE",
	"ONE_MONTH",
	"SEVEN_DAY",
	"SIX_MONTH",
	"THREE_MONTH",
	"TODAY",
	"TWELVE_MONTH",
}

// NewAdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDurationFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDurationFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_audience_converted_time_duration value
func (v AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration) Ptr() *AdlabGroupDetailV30DataDataAdInfoAudienceConvertedTimeDuration {
	return &v
}
