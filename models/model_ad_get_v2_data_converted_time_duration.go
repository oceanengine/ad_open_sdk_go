/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataConvertedTimeDuration
type AdGetV2DataConvertedTimeDuration string

// List of ad_get_v2_data_converted_time_duration
const (
	THREE_MONTH_AdGetV2DataConvertedTimeDuration  AdGetV2DataConvertedTimeDuration = "THREE_MONTH"
	ONE_MONTH_AdGetV2DataConvertedTimeDuration    AdGetV2DataConvertedTimeDuration = "ONE_MONTH"
	SEVEN_DAY_AdGetV2DataConvertedTimeDuration    AdGetV2DataConvertedTimeDuration = "SEVEN_DAY"
	SIX_MONTH_AdGetV2DataConvertedTimeDuration    AdGetV2DataConvertedTimeDuration = "SIX_MONTH"
	TWELVE_MONTH_AdGetV2DataConvertedTimeDuration AdGetV2DataConvertedTimeDuration = "TWELVE_MONTH"
	NONE_AdGetV2DataConvertedTimeDuration         AdGetV2DataConvertedTimeDuration = "NONE"
	TODAY_AdGetV2DataConvertedTimeDuration        AdGetV2DataConvertedTimeDuration = "TODAY"
)

// Ptr returns reference to ad_get_v2_data_converted_time_duration value
func (v AdGetV2DataConvertedTimeDuration) Ptr() *AdGetV2DataConvertedTimeDuration {
	return &v
}
