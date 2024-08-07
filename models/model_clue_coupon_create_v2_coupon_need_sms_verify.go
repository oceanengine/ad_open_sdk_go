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

// ClueCouponCreateV2CouponNeedSmsVerify
type ClueCouponCreateV2CouponNeedSmsVerify string

// List of clue_coupon_create_v2_coupon_need_sms_verify
const (
	Enum_0_ClueCouponCreateV2CouponNeedSmsVerify ClueCouponCreateV2CouponNeedSmsVerify = "0"
	Enum_1_ClueCouponCreateV2CouponNeedSmsVerify ClueCouponCreateV2CouponNeedSmsVerify = "1"
)

// All allowed values of ClueCouponCreateV2CouponNeedSmsVerify enum
var AllowedClueCouponCreateV2CouponNeedSmsVerifyEnumValues = []ClueCouponCreateV2CouponNeedSmsVerify{
	"0",
	"1",
}

// NewClueCouponCreateV2CouponNeedSmsVerifyFromValue returns a pointer to a valid ClueCouponCreateV2CouponNeedSmsVerify
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponCreateV2CouponNeedSmsVerifyFromValue(v string) (*ClueCouponCreateV2CouponNeedSmsVerify, error) {
	ev := ClueCouponCreateV2CouponNeedSmsVerify(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponCreateV2CouponNeedSmsVerify: valid values are %v", v, AllowedClueCouponCreateV2CouponNeedSmsVerifyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponCreateV2CouponNeedSmsVerify) IsValid() bool {
	for _, existing := range AllowedClueCouponCreateV2CouponNeedSmsVerifyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_create_v2_coupon_need_sms_verify value
func (v ClueCouponCreateV2CouponNeedSmsVerify) Ptr() *ClueCouponCreateV2CouponNeedSmsVerify {
	return &v
}
