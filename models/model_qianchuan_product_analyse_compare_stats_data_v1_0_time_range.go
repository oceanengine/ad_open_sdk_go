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

// QianchuanProductAnalyseCompareStatsDataV10TimeRange
type QianchuanProductAnalyseCompareStatsDataV10TimeRange int64

// List of qianchuan_product_analyse_compare_stats_data_v1.0_time_range
const (
	Enum_15_QianchuanProductAnalyseCompareStatsDataV10TimeRange QianchuanProductAnalyseCompareStatsDataV10TimeRange = 15
	Enum_30_QianchuanProductAnalyseCompareStatsDataV10TimeRange QianchuanProductAnalyseCompareStatsDataV10TimeRange = 30
	Enum_7_QianchuanProductAnalyseCompareStatsDataV10TimeRange  QianchuanProductAnalyseCompareStatsDataV10TimeRange = 7
)

// All allowed values of QianchuanProductAnalyseCompareStatsDataV10TimeRange enum
var AllowedQianchuanProductAnalyseCompareStatsDataV10TimeRangeEnumValues = []QianchuanProductAnalyseCompareStatsDataV10TimeRange{
	15,
	30,
	7,
}

// NewQianchuanProductAnalyseCompareStatsDataV10TimeRangeFromValue returns a pointer to a valid QianchuanProductAnalyseCompareStatsDataV10TimeRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanProductAnalyseCompareStatsDataV10TimeRangeFromValue(v int64) (*QianchuanProductAnalyseCompareStatsDataV10TimeRange, error) {
	ev := QianchuanProductAnalyseCompareStatsDataV10TimeRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanProductAnalyseCompareStatsDataV10TimeRange: valid values are %v", v, AllowedQianchuanProductAnalyseCompareStatsDataV10TimeRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanProductAnalyseCompareStatsDataV10TimeRange) IsValid() bool {
	for _, existing := range AllowedQianchuanProductAnalyseCompareStatsDataV10TimeRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_product_analyse_compare_stats_data_v1.0_time_range value
func (v QianchuanProductAnalyseCompareStatsDataV10TimeRange) Ptr() *QianchuanProductAnalyseCompareStatsDataV10TimeRange {
	return &v
}
