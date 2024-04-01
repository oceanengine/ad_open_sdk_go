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

// QianchuanAwemeOrderCreateV10DeliverySettingBidType
type QianchuanAwemeOrderCreateV10DeliverySettingBidType string

// List of qianchuan_aweme_order_create_v1.0_delivery_setting_bid_type
const (
	AUTO_BID_QianchuanAwemeOrderCreateV10DeliverySettingBidType   QianchuanAwemeOrderCreateV10DeliverySettingBidType = "AUTO_BID"
	MANUAL_BID_QianchuanAwemeOrderCreateV10DeliverySettingBidType QianchuanAwemeOrderCreateV10DeliverySettingBidType = "MANUAL_BID"
)

// All allowed values of QianchuanAwemeOrderCreateV10DeliverySettingBidType enum
var AllowedQianchuanAwemeOrderCreateV10DeliverySettingBidTypeEnumValues = []QianchuanAwemeOrderCreateV10DeliverySettingBidType{
	"AUTO_BID",
	"MANUAL_BID",
}

// NewQianchuanAwemeOrderCreateV10DeliverySettingBidTypeFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10DeliverySettingBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10DeliverySettingBidTypeFromValue(v string) (*QianchuanAwemeOrderCreateV10DeliverySettingBidType, error) {
	ev := QianchuanAwemeOrderCreateV10DeliverySettingBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10DeliverySettingBidType: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10DeliverySettingBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10DeliverySettingBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10DeliverySettingBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_delivery_setting_bid_type value
func (v QianchuanAwemeOrderCreateV10DeliverySettingBidType) Ptr() *QianchuanAwemeOrderCreateV10DeliverySettingBidType {
	return &v
}
