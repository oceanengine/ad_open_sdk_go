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

// QianchuanAdGetV10FilteringMarketingGoal
type QianchuanAdGetV10FilteringMarketingGoal string

// List of qianchuan_ad_get_v1.0_filtering_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAdGetV10FilteringMarketingGoal  QianchuanAdGetV10FilteringMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAdGetV10FilteringMarketingGoal QianchuanAdGetV10FilteringMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAdGetV10FilteringMarketingGoal enum
var AllowedQianchuanAdGetV10FilteringMarketingGoalEnumValues = []QianchuanAdGetV10FilteringMarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAdGetV10FilteringMarketingGoalFromValue returns a pointer to a valid QianchuanAdGetV10FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10FilteringMarketingGoalFromValue(v string) (*QianchuanAdGetV10FilteringMarketingGoal, error) {
	ev := QianchuanAdGetV10FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10FilteringMarketingGoal: valid values are %v", v, AllowedQianchuanAdGetV10FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_filtering_marketing_goal value
func (v QianchuanAdGetV10FilteringMarketingGoal) Ptr() *QianchuanAdGetV10FilteringMarketingGoal {
	return &v
}
