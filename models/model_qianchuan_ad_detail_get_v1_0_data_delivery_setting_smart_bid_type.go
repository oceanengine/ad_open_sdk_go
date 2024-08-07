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

// QianchuanAdDetailGetV10DataDeliverySettingSmartBidType
type QianchuanAdDetailGetV10DataDeliverySettingSmartBidType string

// List of qianchuan_ad_detail_get_v1.0_data_delivery_setting_smart_bid_type
const (
	SMART_BID_CONSERVATIVE_QianchuanAdDetailGetV10DataDeliverySettingSmartBidType QianchuanAdDetailGetV10DataDeliverySettingSmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_CUSTOM_QianchuanAdDetailGetV10DataDeliverySettingSmartBidType       QianchuanAdDetailGetV10DataDeliverySettingSmartBidType = "SMART_BID_CUSTOM"
)

// All allowed values of QianchuanAdDetailGetV10DataDeliverySettingSmartBidType enum
var AllowedQianchuanAdDetailGetV10DataDeliverySettingSmartBidTypeEnumValues = []QianchuanAdDetailGetV10DataDeliverySettingSmartBidType{
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_CUSTOM",
}

// NewQianchuanAdDetailGetV10DataDeliverySettingSmartBidTypeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataDeliverySettingSmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataDeliverySettingSmartBidTypeFromValue(v string) (*QianchuanAdDetailGetV10DataDeliverySettingSmartBidType, error) {
	ev := QianchuanAdDetailGetV10DataDeliverySettingSmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataDeliverySettingSmartBidType: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataDeliverySettingSmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataDeliverySettingSmartBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataDeliverySettingSmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_delivery_setting_smart_bid_type value
func (v QianchuanAdDetailGetV10DataDeliverySettingSmartBidType) Ptr() *QianchuanAdDetailGetV10DataDeliverySettingSmartBidType {
	return &v
}
