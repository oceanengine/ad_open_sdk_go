/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeReportOrderGetV10FilteringMarketingGoal
type QianchuanAwemeReportOrderGetV10FilteringMarketingGoal string

// List of qianchuan_aweme_report_order_get_v1.0_filtering_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAwemeReportOrderGetV10FilteringMarketingGoal  QianchuanAwemeReportOrderGetV10FilteringMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAwemeReportOrderGetV10FilteringMarketingGoal QianchuanAwemeReportOrderGetV10FilteringMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAwemeReportOrderGetV10FilteringMarketingGoal enum
var AllowedQianchuanAwemeReportOrderGetV10FilteringMarketingGoalEnumValues = []QianchuanAwemeReportOrderGetV10FilteringMarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAwemeReportOrderGetV10FilteringMarketingGoalFromValue returns a pointer to a valid QianchuanAwemeReportOrderGetV10FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeReportOrderGetV10FilteringMarketingGoalFromValue(v string) (*QianchuanAwemeReportOrderGetV10FilteringMarketingGoal, error) {
	ev := QianchuanAwemeReportOrderGetV10FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeReportOrderGetV10FilteringMarketingGoal: valid values are %v", v, AllowedQianchuanAwemeReportOrderGetV10FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeReportOrderGetV10FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeReportOrderGetV10FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_report_order_get_v1.0_filtering_marketing_goal value
func (v QianchuanAwemeReportOrderGetV10FilteringMarketingGoal) Ptr() *QianchuanAwemeReportOrderGetV10FilteringMarketingGoal {
	return &v
}
