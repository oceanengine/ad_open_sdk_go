/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportCreativeGetV10TimeGranularity
type QianchuanReportCreativeGetV10TimeGranularity string

// List of qianchuan_report_creative_get_v1.0_time_granularity
const (
	TIME_GRANULARITY_DAILY_QianchuanReportCreativeGetV10TimeGranularity  QianchuanReportCreativeGetV10TimeGranularity = "TIME_GRANULARITY_DAILY"
	TIME_GRANULARITY_HOURLY_QianchuanReportCreativeGetV10TimeGranularity QianchuanReportCreativeGetV10TimeGranularity = "TIME_GRANULARITY_HOURLY"
)

// All allowed values of QianchuanReportCreativeGetV10TimeGranularity enum
var AllowedQianchuanReportCreativeGetV10TimeGranularityEnumValues = []QianchuanReportCreativeGetV10TimeGranularity{
	"TIME_GRANULARITY_DAILY",
	"TIME_GRANULARITY_HOURLY",
}

// NewQianchuanReportCreativeGetV10TimeGranularityFromValue returns a pointer to a valid QianchuanReportCreativeGetV10TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportCreativeGetV10TimeGranularityFromValue(v string) (*QianchuanReportCreativeGetV10TimeGranularity, error) {
	ev := QianchuanReportCreativeGetV10TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportCreativeGetV10TimeGranularity: valid values are %v", v, AllowedQianchuanReportCreativeGetV10TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportCreativeGetV10TimeGranularity) IsValid() bool {
	for _, existing := range AllowedQianchuanReportCreativeGetV10TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_creative_get_v1.0_time_granularity value
func (v QianchuanReportCreativeGetV10TimeGranularity) Ptr() *QianchuanReportCreativeGetV10TimeGranularity {
	return &v
}
