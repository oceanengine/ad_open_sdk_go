/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportSearchWordGetV10FilteringMarketingGoal
type QianchuanReportSearchWordGetV10FilteringMarketingGoal string

// List of qianchuan_report_search_word_get_v1.0_filtering_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanReportSearchWordGetV10FilteringMarketingGoal  QianchuanReportSearchWordGetV10FilteringMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanReportSearchWordGetV10FilteringMarketingGoal QianchuanReportSearchWordGetV10FilteringMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanReportSearchWordGetV10FilteringMarketingGoal enum
var AllowedQianchuanReportSearchWordGetV10FilteringMarketingGoalEnumValues = []QianchuanReportSearchWordGetV10FilteringMarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanReportSearchWordGetV10FilteringMarketingGoalFromValue returns a pointer to a valid QianchuanReportSearchWordGetV10FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportSearchWordGetV10FilteringMarketingGoalFromValue(v string) (*QianchuanReportSearchWordGetV10FilteringMarketingGoal, error) {
	ev := QianchuanReportSearchWordGetV10FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportSearchWordGetV10FilteringMarketingGoal: valid values are %v", v, AllowedQianchuanReportSearchWordGetV10FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportSearchWordGetV10FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanReportSearchWordGetV10FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_search_word_get_v1.0_filtering_marketing_goal value
func (v QianchuanReportSearchWordGetV10FilteringMarketingGoal) Ptr() *QianchuanReportSearchWordGetV10FilteringMarketingGoal {
	return &v
}
