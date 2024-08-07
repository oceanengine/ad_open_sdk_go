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

// QianchuanAdCreateV10DeliverySettingSmartBidType
type QianchuanAdCreateV10DeliverySettingSmartBidType string

// List of qianchuan_ad_create_v1.0_delivery_setting_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanAdCreateV10DeliverySettingSmartBidType QianchuanAdCreateV10DeliverySettingSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanAdCreateV10DeliverySettingSmartBidType       QianchuanAdCreateV10DeliverySettingSmartBidType = "SMART_BID_CUSTOM"
)

// All allowed values of QianchuanAdCreateV10DeliverySettingSmartBidType enum
var AllowedQianchuanAdCreateV10DeliverySettingSmartBidTypeEnumValues = []QianchuanAdCreateV10DeliverySettingSmartBidType{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewQianchuanAdCreateV10DeliverySettingSmartBidTypeFromValue returns a pointer to a valid QianchuanAdCreateV10DeliverySettingSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10DeliverySettingSmartBidTypeFromValue(v string) (*QianchuanAdCreateV10DeliverySettingSmartBidType, error) {
	ev := QianchuanAdCreateV10DeliverySettingSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10DeliverySettingSmartBidType: valid values are %v", v, AllowedQianchuanAdCreateV10DeliverySettingSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10DeliverySettingSmartBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10DeliverySettingSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_smart_bid_type value
func (v QianchuanAdCreateV10DeliverySettingSmartBidType) Ptr() *QianchuanAdCreateV10DeliverySettingSmartBidType {
	return &v
}
