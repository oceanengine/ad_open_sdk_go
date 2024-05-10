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

// QianchuanAwemeOrderDetailGetV10DataMarketingGoal
type QianchuanAwemeOrderDetailGetV10DataMarketingGoal string

// List of qianchuan_aweme_order_detail_get_v1.0_data_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanAwemeOrderDetailGetV10DataMarketingGoal  QianchuanAwemeOrderDetailGetV10DataMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanAwemeOrderDetailGetV10DataMarketingGoal QianchuanAwemeOrderDetailGetV10DataMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataMarketingGoal enum
var AllowedQianchuanAwemeOrderDetailGetV10DataMarketingGoalEnumValues = []QianchuanAwemeOrderDetailGetV10DataMarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanAwemeOrderDetailGetV10DataMarketingGoalFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataMarketingGoalFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataMarketingGoal, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataMarketingGoal: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_marketing_goal value
func (v QianchuanAwemeOrderDetailGetV10DataMarketingGoal) Ptr() *QianchuanAwemeOrderDetailGetV10DataMarketingGoal {
	return &v
}
