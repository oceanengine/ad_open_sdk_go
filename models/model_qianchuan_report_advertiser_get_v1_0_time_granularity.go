/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportAdvertiserGetV10TimeGranularity
type QianchuanReportAdvertiserGetV10TimeGranularity string

// List of qianchuan_report_advertiser_get_v1.0_time_granularity
const (
	TIME_GRANULARITY_DAILY_QianchuanReportAdvertiserGetV10TimeGranularity  QianchuanReportAdvertiserGetV10TimeGranularity = "TIME_GRANULARITY_DAILY"
	TIME_GRANULARITY_HOURLY_QianchuanReportAdvertiserGetV10TimeGranularity QianchuanReportAdvertiserGetV10TimeGranularity = "TIME_GRANULARITY_HOURLY"
)

// All allowed values of QianchuanReportAdvertiserGetV10TimeGranularity enum
var AllowedQianchuanReportAdvertiserGetV10TimeGranularityEnumValues = []QianchuanReportAdvertiserGetV10TimeGranularity{
	"TIME_GRANULARITY_DAILY",
	"TIME_GRANULARITY_HOURLY",
}

// NewQianchuanReportAdvertiserGetV10TimeGranularityFromValue returns a pointer to a valid QianchuanReportAdvertiserGetV10TimeGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportAdvertiserGetV10TimeGranularityFromValue(v string) (*QianchuanReportAdvertiserGetV10TimeGranularity, error) {
	ev := QianchuanReportAdvertiserGetV10TimeGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportAdvertiserGetV10TimeGranularity: valid values are %v", v, AllowedQianchuanReportAdvertiserGetV10TimeGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportAdvertiserGetV10TimeGranularity) IsValid() bool {
	for _, existing := range AllowedQianchuanReportAdvertiserGetV10TimeGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_advertiser_get_v1.0_time_granularity value
func (v QianchuanReportAdvertiserGetV10TimeGranularity) Ptr() *QianchuanReportAdvertiserGetV10TimeGranularity {
	return &v
}
