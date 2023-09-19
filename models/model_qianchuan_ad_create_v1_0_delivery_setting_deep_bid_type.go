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

// QianchuanAdCreateV10DeliverySettingDeepBidType
type QianchuanAdCreateV10DeliverySettingDeepBidType string

// List of qianchuan_ad_create_v1.0_delivery_setting_deep_bid_type
const (
	COMM_ROI_QianchuanAdCreateV10DeliverySettingDeepBidType         QianchuanAdCreateV10DeliverySettingDeepBidType = "COMM_ROI"
	MIN_QianchuanAdCreateV10DeliverySettingDeepBidType              QianchuanAdCreateV10DeliverySettingDeepBidType = "MIN"
	MIN_SECOND_STAGE_QianchuanAdCreateV10DeliverySettingDeepBidType QianchuanAdCreateV10DeliverySettingDeepBidType = "MIN_SECOND_STAGE"
	PACING_QianchuanAdCreateV10DeliverySettingDeepBidType           QianchuanAdCreateV10DeliverySettingDeepBidType = "PACING"
	PAY_ROI_QianchuanAdCreateV10DeliverySettingDeepBidType          QianchuanAdCreateV10DeliverySettingDeepBidType = "PAY_ROI"
)

// All allowed values of QianchuanAdCreateV10DeliverySettingDeepBidType enum
var AllowedQianchuanAdCreateV10DeliverySettingDeepBidTypeEnumValues = []QianchuanAdCreateV10DeliverySettingDeepBidType{
	"COMM_ROI",
	"MIN",
	"MIN_SECOND_STAGE",
	"PACING",
	"PAY_ROI",
}

// NewQianchuanAdCreateV10DeliverySettingDeepBidTypeFromValue returns a pointer to a valid QianchuanAdCreateV10DeliverySettingDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10DeliverySettingDeepBidTypeFromValue(v string) (*QianchuanAdCreateV10DeliverySettingDeepBidType, error) {
	ev := QianchuanAdCreateV10DeliverySettingDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10DeliverySettingDeepBidType: valid values are %v", v, AllowedQianchuanAdCreateV10DeliverySettingDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10DeliverySettingDeepBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10DeliverySettingDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_deep_bid_type value
func (v QianchuanAdCreateV10DeliverySettingDeepBidType) Ptr() *QianchuanAdCreateV10DeliverySettingDeepBidType {
	return &v
}
