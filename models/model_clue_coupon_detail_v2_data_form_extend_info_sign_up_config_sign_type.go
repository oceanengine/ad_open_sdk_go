/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType
type ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType string

// List of clue_coupon_detail_v2_data_form_extend_info_sign_up_config_sign_type
const (
	SIGN_TYPE_SCROLL_WALL_ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_WALL"
	SIGN_TYPE_SCROLL_BAR_ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType  ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_BAR"
)

// All allowed values of ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType enum
var AllowedClueCouponDetailV2DataFormExtendInfoSignUpConfigSignTypeEnumValues = []ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType{
	"SIGN_TYPE_SCROLL_WALL",
	"SIGN_TYPE_SCROLL_BAR",
}

// NewClueCouponDetailV2DataFormExtendInfoSignUpConfigSignTypeFromValue returns a pointer to a valid ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataFormExtendInfoSignUpConfigSignTypeFromValue(v string) (*ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType, error) {
	ev := ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType: valid values are %v", v, AllowedClueCouponDetailV2DataFormExtendInfoSignUpConfigSignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataFormExtendInfoSignUpConfigSignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_form_extend_info_sign_up_config_sign_type value
func (v ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType) Ptr() *ClueCouponDetailV2DataFormExtendInfoSignUpConfigSignType {
	return &v
}
