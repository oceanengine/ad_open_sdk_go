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

// QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal
type QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal string

// List of qianchuan_uni_promotion_list_v1.0_data_ad_list_ad_info_marketing_goal
const (
	ALL_QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal              QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal = "ALL"
	LIVE_PROM_GOODS_QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal  QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal enum
var AllowedQianchuanUniPromotionListV10DataAdListAdInfoMarketingGoalEnumValues = []QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal{
	"ALL",
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanUniPromotionListV10DataAdListAdInfoMarketingGoalFromValue returns a pointer to a valid QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionListV10DataAdListAdInfoMarketingGoalFromValue(v string) (*QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal, error) {
	ev := QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal: valid values are %v", v, AllowedQianchuanUniPromotionListV10DataAdListAdInfoMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionListV10DataAdListAdInfoMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_data_ad_list_ad_info_marketing_goal value
func (v QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal) Ptr() *QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal {
	return &v
}
