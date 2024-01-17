/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCreativeGetV10FilteringMarketingGoal
type QianchuanCreativeGetV10FilteringMarketingGoal string

// List of qianchuan_creative_get_v1.0_filtering_marketing_goal
const (
	ALL_QianchuanCreativeGetV10FilteringMarketingGoal              QianchuanCreativeGetV10FilteringMarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanCreativeGetV10FilteringMarketingGoal  QianchuanCreativeGetV10FilteringMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanCreativeGetV10FilteringMarketingGoal QianchuanCreativeGetV10FilteringMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanCreativeGetV10FilteringMarketingGoal enum
var AllowedQianchuanCreativeGetV10FilteringMarketingGoalEnumValues = []QianchuanCreativeGetV10FilteringMarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanCreativeGetV10FilteringMarketingGoalFromValue returns a pointer to a valid QianchuanCreativeGetV10FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCreativeGetV10FilteringMarketingGoalFromValue(v string) (*QianchuanCreativeGetV10FilteringMarketingGoal, error) {
	ev := QianchuanCreativeGetV10FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCreativeGetV10FilteringMarketingGoal: valid values are %v", v, AllowedQianchuanCreativeGetV10FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCreativeGetV10FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanCreativeGetV10FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_creative_get_v1.0_filtering_marketing_goal value
func (v QianchuanCreativeGetV10FilteringMarketingGoal) Ptr() *QianchuanCreativeGetV10FilteringMarketingGoal {
	return &v
}
