/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponGetV2DataListCouponNeedPhone
type ClueCouponGetV2DataListCouponNeedPhone string

// List of clue_coupon_get_v2_data_list_coupon_need_phone
const (
	Enum_0_ClueCouponGetV2DataListCouponNeedPhone ClueCouponGetV2DataListCouponNeedPhone = "0"
	Enum_1_ClueCouponGetV2DataListCouponNeedPhone ClueCouponGetV2DataListCouponNeedPhone = "1"
)

// All allowed values of ClueCouponGetV2DataListCouponNeedPhone enum
var AllowedClueCouponGetV2DataListCouponNeedPhoneEnumValues = []ClueCouponGetV2DataListCouponNeedPhone{
	"0",
	"1",
}

// NewClueCouponGetV2DataListCouponNeedPhoneFromValue returns a pointer to a valid ClueCouponGetV2DataListCouponNeedPhone
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponGetV2DataListCouponNeedPhoneFromValue(v string) (*ClueCouponGetV2DataListCouponNeedPhone, error) {
	ev := ClueCouponGetV2DataListCouponNeedPhone(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponGetV2DataListCouponNeedPhone: valid values are %v", v, AllowedClueCouponGetV2DataListCouponNeedPhoneEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponGetV2DataListCouponNeedPhone) IsValid() bool {
	for _, existing := range AllowedClueCouponGetV2DataListCouponNeedPhoneEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_need_phone value
func (v ClueCouponGetV2DataListCouponNeedPhone) Ptr() *ClueCouponGetV2DataListCouponNeedPhone {
	return &v
}
