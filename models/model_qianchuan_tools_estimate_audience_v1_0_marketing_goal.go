/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10MarketingGoal
type QianchuanToolsEstimateAudienceV10MarketingGoal string

// List of qianchuan_tools_estimate_audience_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanToolsEstimateAudienceV10MarketingGoal  QianchuanToolsEstimateAudienceV10MarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanToolsEstimateAudienceV10MarketingGoal QianchuanToolsEstimateAudienceV10MarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanToolsEstimateAudienceV10MarketingGoal enum
var AllowedQianchuanToolsEstimateAudienceV10MarketingGoalEnumValues = []QianchuanToolsEstimateAudienceV10MarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanToolsEstimateAudienceV10MarketingGoalFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10MarketingGoalFromValue(v string) (*QianchuanToolsEstimateAudienceV10MarketingGoal, error) {
	ev := QianchuanToolsEstimateAudienceV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10MarketingGoal: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_marketing_goal value
func (v QianchuanToolsEstimateAudienceV10MarketingGoal) Ptr() *QianchuanToolsEstimateAudienceV10MarketingGoal {
	return &v
}
