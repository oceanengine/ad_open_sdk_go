/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanUniPromotionListV10MarketingGoal
type QianchuanUniPromotionListV10MarketingGoal string

// List of qianchuan_uni_promotion_list_v1.0_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanUniPromotionListV10MarketingGoal QianchuanUniPromotionListV10MarketingGoal = "LIVE_PROM_GOODS"
)

// All allowed values of QianchuanUniPromotionListV10MarketingGoal enum
var AllowedQianchuanUniPromotionListV10MarketingGoalEnumValues = []QianchuanUniPromotionListV10MarketingGoal{
	"LIVE_PROM_GOODS",
}

// NewQianchuanUniPromotionListV10MarketingGoalFromValue returns a pointer to a valid QianchuanUniPromotionListV10MarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionListV10MarketingGoalFromValue(v string) (*QianchuanUniPromotionListV10MarketingGoal, error) {
	ev := QianchuanUniPromotionListV10MarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionListV10MarketingGoal: valid values are %v", v, AllowedQianchuanUniPromotionListV10MarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionListV10MarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionListV10MarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_marketing_goal value
func (v QianchuanUniPromotionListV10MarketingGoal) Ptr() *QianchuanUniPromotionListV10MarketingGoal {
	return &v
}
