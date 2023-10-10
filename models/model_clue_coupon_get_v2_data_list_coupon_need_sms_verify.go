/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponGetV2DataListCouponNeedSmsVerify
type ClueCouponGetV2DataListCouponNeedSmsVerify string

// List of clue_coupon_get_v2_data_list_coupon_need_sms_verify
const (
	Enum_0_ClueCouponGetV2DataListCouponNeedSmsVerify ClueCouponGetV2DataListCouponNeedSmsVerify = "0"
	Enum_1_ClueCouponGetV2DataListCouponNeedSmsVerify ClueCouponGetV2DataListCouponNeedSmsVerify = "1"
)

// All allowed values of ClueCouponGetV2DataListCouponNeedSmsVerify enum
var AllowedClueCouponGetV2DataListCouponNeedSmsVerifyEnumValues = []ClueCouponGetV2DataListCouponNeedSmsVerify{
	"0",
	"1",
}

// NewClueCouponGetV2DataListCouponNeedSmsVerifyFromValue returns a pointer to a valid ClueCouponGetV2DataListCouponNeedSmsVerify
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponGetV2DataListCouponNeedSmsVerifyFromValue(v string) (*ClueCouponGetV2DataListCouponNeedSmsVerify, error) {
	ev := ClueCouponGetV2DataListCouponNeedSmsVerify(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponGetV2DataListCouponNeedSmsVerify: valid values are %v", v, AllowedClueCouponGetV2DataListCouponNeedSmsVerifyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponGetV2DataListCouponNeedSmsVerify) IsValid() bool {
	for _, existing := range AllowedClueCouponGetV2DataListCouponNeedSmsVerifyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_get_v2_data_list_coupon_need_sms_verify value
func (v ClueCouponGetV2DataListCouponNeedSmsVerify) Ptr() *ClueCouponGetV2DataListCouponNeedSmsVerify {
	return &v
}
