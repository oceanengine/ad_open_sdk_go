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

// QianchuanAwemeEstimateProfitV10MarketingGoal
type QianchuanAwemeEstimateProfitV10MarketingGoal string

// List of qianchuan_aweme_estimate_profit_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAwemeEstimateProfitV10MarketingGoal  QianchuanAwemeEstimateProfitV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAwemeEstimateProfitV10MarketingGoal QianchuanAwemeEstimateProfitV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAwemeEstimateProfitV10MarketingGoal enum
var AllowedQianchuanAwemeEstimateProfitV10MarketingGoalEnumValues = []QianchuanAwemeEstimateProfitV10MarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAwemeEstimateProfitV10MarketingGoalFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10MarketingGoalFromValue(v string) (*QianchuanAwemeEstimateProfitV10MarketingGoal, error) {
	ev := QianchuanAwemeEstimateProfitV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10MarketingGoal: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_marketing_goal value
func (v QianchuanAwemeEstimateProfitV10MarketingGoal) Ptr() *QianchuanAwemeEstimateProfitV10MarketingGoal {
	return &v
}
