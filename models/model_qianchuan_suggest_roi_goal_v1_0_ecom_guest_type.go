/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanSuggestRoiGoalV10EcomGuestType
type QianchuanSuggestRoiGoalV10EcomGuestType string

// List of qianchuan_suggest_roi_goal_v1.0_ecom_guest_type
const (
	NO_BUY_QianchuanSuggestRoiGoalV10EcomGuestType        QianchuanSuggestRoiGoalV10EcomGuestType = "NO_BUY"
	NO_BUY_BRAND_QianchuanSuggestRoiGoalV10EcomGuestType  QianchuanSuggestRoiGoalV10EcomGuestType = "NO_BUY_BRAND"
	NO_BUY_DOUYIN_QianchuanSuggestRoiGoalV10EcomGuestType QianchuanSuggestRoiGoalV10EcomGuestType = "NO_BUY_DOUYIN"
)

// All allowed values of QianchuanSuggestRoiGoalV10EcomGuestType enum
var AllowedQianchuanSuggestRoiGoalV10EcomGuestTypeEnumValues = []QianchuanSuggestRoiGoalV10EcomGuestType{
	"NO_BUY",
	"NO_BUY_BRAND",
	"NO_BUY_DOUYIN",
}

// NewQianchuanSuggestRoiGoalV10EcomGuestTypeFromValue returns a pointer to a valid QianchuanSuggestRoiGoalV10EcomGuestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanSuggestRoiGoalV10EcomGuestTypeFromValue(v string) (*QianchuanSuggestRoiGoalV10EcomGuestType, error) {
	ev := QianchuanSuggestRoiGoalV10EcomGuestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanSuggestRoiGoalV10EcomGuestType: valid values are %v", v, AllowedQianchuanSuggestRoiGoalV10EcomGuestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanSuggestRoiGoalV10EcomGuestType) IsValid() bool {
	for _, existing := range AllowedQianchuanSuggestRoiGoalV10EcomGuestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_suggest_roi_goal_v1.0_ecom_guest_type value
func (v QianchuanSuggestRoiGoalV10EcomGuestType) Ptr() *QianchuanSuggestRoiGoalV10EcomGuestType {
	return &v
}
