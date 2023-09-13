/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10DataListDeliverySettingDeepBidType
type QianchuanAdGetV10DataListDeliverySettingDeepBidType string

// List of qianchuan_ad_get_v1.0_data_list_delivery_setting_deep_bid_type
const (
	COMM_ROI_QianchuanAdGetV10DataListDeliverySettingDeepBidType         QianchuanAdGetV10DataListDeliverySettingDeepBidType = "COMM_ROI"
	MIN_QianchuanAdGetV10DataListDeliverySettingDeepBidType              QianchuanAdGetV10DataListDeliverySettingDeepBidType = "MIN"
	MIN_SECOND_STAGE_QianchuanAdGetV10DataListDeliverySettingDeepBidType QianchuanAdGetV10DataListDeliverySettingDeepBidType = "MIN_SECOND_STAGE"
	PACING_QianchuanAdGetV10DataListDeliverySettingDeepBidType           QianchuanAdGetV10DataListDeliverySettingDeepBidType = "PACING"
	PAY_ROI_QianchuanAdGetV10DataListDeliverySettingDeepBidType          QianchuanAdGetV10DataListDeliverySettingDeepBidType = "PAY_ROI"
)

// All allowed values of QianchuanAdGetV10DataListDeliverySettingDeepBidType enum
var AllowedQianchuanAdGetV10DataListDeliverySettingDeepBidTypeEnumValues = []QianchuanAdGetV10DataListDeliverySettingDeepBidType{
	"COMM_ROI",
	"MIN",
	"MIN_SECOND_STAGE",
	"PACING",
	"PAY_ROI",
}

// NewQianchuanAdGetV10DataListDeliverySettingDeepBidTypeFromValue returns a pointer to a valid QianchuanAdGetV10DataListDeliverySettingDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListDeliverySettingDeepBidTypeFromValue(v string) (*QianchuanAdGetV10DataListDeliverySettingDeepBidType, error) {
	ev := QianchuanAdGetV10DataListDeliverySettingDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListDeliverySettingDeepBidType: valid values are %v", v, AllowedQianchuanAdGetV10DataListDeliverySettingDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListDeliverySettingDeepBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListDeliverySettingDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_delivery_setting_deep_bid_type value
func (v QianchuanAdGetV10DataListDeliverySettingDeepBidType) Ptr() *QianchuanAdGetV10DataListDeliverySettingDeepBidType {
	return &v
}
