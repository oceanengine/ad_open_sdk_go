/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderGetV10FilteringMarketingGoal
type QianchuanAwemeOrderGetV10FilteringMarketingGoal string

// List of qianchuan_aweme_order_get_v1.0_filtering_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAwemeOrderGetV10FilteringMarketingGoal  QianchuanAwemeOrderGetV10FilteringMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAwemeOrderGetV10FilteringMarketingGoal QianchuanAwemeOrderGetV10FilteringMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAwemeOrderGetV10FilteringMarketingGoal enum
var AllowedQianchuanAwemeOrderGetV10FilteringMarketingGoalEnumValues = []QianchuanAwemeOrderGetV10FilteringMarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAwemeOrderGetV10FilteringMarketingGoalFromValue returns a pointer to a valid QianchuanAwemeOrderGetV10FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderGetV10FilteringMarketingGoalFromValue(v string) (*QianchuanAwemeOrderGetV10FilteringMarketingGoal, error) {
	ev := QianchuanAwemeOrderGetV10FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderGetV10FilteringMarketingGoal: valid values are %v", v, AllowedQianchuanAwemeOrderGetV10FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderGetV10FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderGetV10FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_filtering_marketing_goal value
func (v QianchuanAwemeOrderGetV10FilteringMarketingGoal) Ptr() *QianchuanAwemeOrderGetV10FilteringMarketingGoal {
	return &v
}
