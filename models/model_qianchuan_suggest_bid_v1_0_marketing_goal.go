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

// QianchuanSuggestBidV10MarketingGoal
type QianchuanSuggestBidV10MarketingGoal string

// List of qianchuan_suggest_bid_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanSuggestBidV10MarketingGoal  QianchuanSuggestBidV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanSuggestBidV10MarketingGoal QianchuanSuggestBidV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanSuggestBidV10MarketingGoal enum
var AllowedQianchuanSuggestBidV10MarketingGoalEnumValues = []QianchuanSuggestBidV10MarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanSuggestBidV10MarketingGoalFromValue returns a pointer to a valid QianchuanSuggestBidV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestBidV10MarketingGoalFromValue(v string) (*QianchuanSuggestBidV10MarketingGoal, error) {
	ev := QianchuanSuggestBidV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestBidV10MarketingGoal: valid values are %v", v, AllowedQianchuanSuggestBidV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestBidV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestBidV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_bid_v1.0_marketing_goal value
func (v QianchuanSuggestBidV10MarketingGoal) Ptr() *QianchuanSuggestBidV10MarketingGoal {
	return &v
}
