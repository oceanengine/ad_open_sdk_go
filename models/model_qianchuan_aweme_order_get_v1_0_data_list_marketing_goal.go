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

// QianchuanAwemeOrderGetV10DataListMarketingGoal
type QianchuanAwemeOrderGetV10DataListMarketingGoal string

// List of qianchuan_aweme_order_get_v1.0_data_list_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAwemeOrderGetV10DataListMarketingGoal  QianchuanAwemeOrderGetV10DataListMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAwemeOrderGetV10DataListMarketingGoal QianchuanAwemeOrderGetV10DataListMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAwemeOrderGetV10DataListMarketingGoal enum
var AllowedQianchuanAwemeOrderGetV10DataListMarketingGoalEnumValues = []QianchuanAwemeOrderGetV10DataListMarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAwemeOrderGetV10DataListMarketingGoalFromValue returns a pointer to a valid QianchuanAwemeOrderGetV10DataListMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderGetV10DataListMarketingGoalFromValue(v string) (*QianchuanAwemeOrderGetV10DataListMarketingGoal, error) {
	ev := QianchuanAwemeOrderGetV10DataListMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderGetV10DataListMarketingGoal: valid values are %v", v, AllowedQianchuanAwemeOrderGetV10DataListMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderGetV10DataListMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderGetV10DataListMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_get_v1.0_data_list_marketing_goal value
func (v QianchuanAwemeOrderGetV10DataListMarketingGoal) Ptr() *QianchuanAwemeOrderGetV10DataListMarketingGoal {
	return &v
}
