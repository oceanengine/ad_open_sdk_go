/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10MarketingGoal
type QianchuanAdCreateV10MarketingGoal string

// List of qianchuan_ad_create_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAdCreateV10MarketingGoal  QianchuanAdCreateV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAdCreateV10MarketingGoal QianchuanAdCreateV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAdCreateV10MarketingGoal enum
var AllowedQianchuanAdCreateV10MarketingGoalEnumValues = []QianchuanAdCreateV10MarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAdCreateV10MarketingGoalFromValue returns a pointer to a valid QianchuanAdCreateV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10MarketingGoalFromValue(v string) (*QianchuanAdCreateV10MarketingGoal, error) {
	ev := QianchuanAdCreateV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10MarketingGoal: valid values are %v", v, AllowedQianchuanAdCreateV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_marketing_goal value
func (v QianchuanAdCreateV10MarketingGoal) Ptr() *QianchuanAdCreateV10MarketingGoal {
	return &v
}
