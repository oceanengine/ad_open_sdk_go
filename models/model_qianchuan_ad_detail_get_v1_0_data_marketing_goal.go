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

// QianchuanAdDetailGetV10DataMarketingGoal
type QianchuanAdDetailGetV10DataMarketingGoal string

// List of qianchuan_ad_detail_get_v1.0_data_marketing_goal
const (
	ALL_QianchuanAdDetailGetV10DataMarketingGoal              QianchuanAdDetailGetV10DataMarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanAdDetailGetV10DataMarketingGoal  QianchuanAdDetailGetV10DataMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAdDetailGetV10DataMarketingGoal QianchuanAdDetailGetV10DataMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAdDetailGetV10DataMarketingGoal enum
var AllowedQianchuanAdDetailGetV10DataMarketingGoalEnumValues = []QianchuanAdDetailGetV10DataMarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAdDetailGetV10DataMarketingGoalFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataMarketingGoalFromValue(v string) (*QianchuanAdDetailGetV10DataMarketingGoal, error) {
	ev := QianchuanAdDetailGetV10DataMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataMarketingGoal: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_marketing_goal value
func (v QianchuanAdDetailGetV10DataMarketingGoal) Ptr() *QianchuanAdDetailGetV10DataMarketingGoal {
	return &v
}
