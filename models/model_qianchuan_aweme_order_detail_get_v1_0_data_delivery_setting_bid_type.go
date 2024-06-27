/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType
type QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType string

// List of qianchuan_aweme_order_detail_get_v1.0_data_delivery_setting_bid_type
const (
	AUTO_BID_QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType   QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType = "AUTO_BID"
	MANUAL_BID_QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType = "MANUAL_BID"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType enum
var AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidTypeEnumValues = []QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType{
	"AUTO_BID",
	"MANUAL_BID",
}

// NewQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidTypeFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidTypeFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_delivery_setting_bid_type value
func (v QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType) Ptr() *QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType {
	return &v
}
